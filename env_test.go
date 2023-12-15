package env_test

import (
	"encoding/json"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/lucafmarques/env"
)

func TestGet(t *testing.T) {
	var (
		testurl, _ = url.Parse("https://github.com/lucafmarques")
		tt         = []struct {
			key      string
			value    string
			expected any
			err      string
		}{
			{
				key:      "TESTE_STRING",
				value:    "string value",
				expected: "string value",
			},
			{
				key:      "TESTE_BOOL",
				value:    "TRUE",
				expected: true,
			},
			{
				key:      "TESTE_BOOL_PARSE_ERROR",
				value:    "truthy",
				expected: false,
				err:      (&strconv.NumError{Func: "ParseBool", Num: "truthy", Err: strconv.ErrSyntax}).Error(),
			},
			{
				key:      "TESTE_INT",
				value:    "10",
				expected: 10,
			},
			{
				key:      "TESTE_INT_PARSE_ERROR",
				value:    "10.0",
				expected: 0,
				err:      (&strconv.NumError{Func: "Atoi", Num: "10.0", Err: strconv.ErrSyntax}).Error(),
			},
			{
				key:      "TESTE_FLOAT64",
				value:    "10",
				expected: 10.,
			},
			{
				key:      "TESTE_FLOAT64_PARSE_ERROR",
				value:    "R$ 4.20",
				expected: 0.,
				err:      (&strconv.NumError{Func: "ParseFloat", Num: "R$ 4.20", Err: strconv.ErrSyntax}).Error(),
			},
			{
				key:      "TESTE_BUILDER",
				value:    `{"name":"Luca Marques","github":"https://github.com/lucafmarques"}`,
				expected: user{Name: "Luca Marques", Github: *testurl},
			},
			{
				key:      "TESTE_MISSING_BUILDER",
				value:    `{"name":"Luca Marques","github":"https://github.com/lucafmarques"}`,
				expected: fail{},
				err:      env.ErrBuilder.Error(),
			},
			{
				key:      "TEST_UNSET",
				value:    "",
				expected: "",
				err:      env.ErrUnset.Error(),
			},
		}
	)

	for _, tc := range tt {
		t.Run(tc.key, func(t *testing.T) {
			var (
				res any
				err error
			)

			if tc.value != "" {
				if err := os.Setenv(tc.key, tc.value); err != nil {
					t.Fatal(err)
				}
			}
			defer func() {
				if err := os.Unsetenv(tc.key); err != nil {
					t.Fatal(err)
				}
			}()

			switch tc.expected.(type) {
			case string:
				res, err = env.Get[string](tc.key)
			case bool:
				res, err = env.Get[bool](tc.key)
			case int:
				res, err = env.Get[int](tc.key)
			case float64:
				res, err = env.Get[float64](tc.key)
			case user:
				res, err = env.Get[user](tc.key)
			case fail:
				res, err = env.Get[fail](tc.key)
			}

			if err != nil {
				be.Equal(t, tc.err, err.Error())
			}
			be.Equal(t, tc.expected, res)
		})
	}
}

func TestMustGet(t *testing.T) {
	tt := []struct {
		key      string
		value    string
		expected string
		set      bool
	}{
		{
			key:      "TEST_SET",
			value:    "test",
			expected: "test",
			set:      true,
		},
		{
			key:      "TEST_UNSET",
			value:    "test",
			expected: "",
			set:      false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.key, func(t *testing.T) {
			if tc.set {
				if err := os.Setenv(tc.key, tc.value); err != nil {
					t.Fatal(err)
				}
				defer func() {
					if err := os.Unsetenv(tc.key); err != nil {
						t.Fatal(err)
					}
				}()
			}

			defer func() {
				if err := recover(); err != nil {
					be.Nonzero(t, err)
				}
			}()

			be.Equal(t, tc.expected, env.MustGet[string](tc.key))
		})
	}
}

type fail struct{}

type user struct {
	Name   string  `json:"name"`
	Github url.URL `json:"github"`
}

func (u user) Build(env string) (any, error) {
	var us user
	if err := json.Unmarshal([]byte(env), &us); err != nil {
		return user{}, err
	}

	return us, nil
}

func (u *user) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Github string `json:"github"`
		Name   string `json:"name"`
	}{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	url, err := url.Parse(tmp.Github)
	if err != nil {
		return err
	}

	u.Github = *url
	u.Name = tmp.Name

	return nil
}
