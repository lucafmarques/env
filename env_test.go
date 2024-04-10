package env_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/carlmjohnson/be"
	"github.com/lucafmarques/env"
)

var Error = errors.New("mock test error")

func TestGet(t *testing.T) {
	testurl, _ := url.Parse("https://github.com/lucafmarques")
	tt := []struct {
		key      string
		value    string
		expected any
		err      string
		fallback any
	}{
		{
			key:      "TESTE_STRING",
			value:    "string value",
			expected: "string value",
		},
		{
			key:      "TESTE_STRING_FALLBACK",
			expected: "fallback",
			fallback: "fallback",
			err:      env.ErrUnset.Error(),
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
			key:      "TESTE_CUSTOM_TEXT_UNMARSHALER",
			value:    `{"name":"Luca Marques","github":"https://github.com/lucafmarques"}`,
			expected: user{Name: "Luca Marques", Github: *testurl},
		},
		{
			key:      "TESTE_UNIMPLEMENTED_TEXT_UNMARSHALER",
			value:    `{"name":"Luca Marques","github":"https://github.com/lucafmarques"}`,
			expected: fail{},
			err:      env.ErrUnmarshaler.Error(),
		},
		{
			key:      "TEST_UNSET",
			value:    "",
			expected: "",
			err:      env.ErrUnset.Error(),
		},
	}

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
			defer os.Unsetenv(tc.key)

			switch tc.expected.(type) {
			case string:
				if s, ok := tc.fallback.(string); !ok {
					res, err = env.Get[string](tc.key)
				} else {
					res, err = env.Get[string](tc.key, s)
				}
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
				if err := recover(); err == env.ErrUnset {
					be.Nonzero(t, err)
				}
			}()

			be.Equal(t, tc.expected, env.MustGet[string](tc.key))
		})
	}
}

func TestSet(t *testing.T) {
	now := time.Now()

	tt := []struct {
		key      string
		value    any
		expected string
		err      error
	}{
		{
			key:      "TESTE_BOOL",
			value:    true,
			expected: "true",
		},
		{
			key:      "TESTE_STRING",
			value:    "string",
			expected: "string",
		},
		{
			key:      "TESTE_INT",
			value:    10,
			expected: "10",
		},
		{
			key:      "TESTE_FLOAT",
			value:    4.20,
			expected: "4.2",
		},
		{
			key:      "TESTE_TIME.TIME",
			value:    now,
			expected: now.Format(time.RFC3339Nano),
		},
		{
			key:      "TESTE_CUSTOM_TEXT_MARSHALER",
			value:    user{Name: "Luca Marques", Github: url.URL{Host: "github.com", Path: "/lucafmarques", Scheme: "https"}},
			expected: "Name: Luca Marques - GitHub: https://github.com/lucafmarques",
		},
		{
			key:      "TESTE_MARSHAL_TEXT_ERROR",
			value:    user{error: Error},
			expected: "",
			err:      Error,
		},
		{
			key:      "TESTE_UNIMPLEMENTED_TEXT_MARSHALER_ERROR",
			value:    fail{},
			expected: "",
			err:      env.ErrMarshaler,
		},
	}

	for _, tc := range tt {
		t.Run(tc.key, func(t *testing.T) {
			err := env.Set(tc.key, tc.value)
			val := os.Getenv(tc.key)

			be.Equal(t, tc.err, err)
			be.Equal(t, tc.expected, val)
		})
	}
}
func TestMustSet(t *testing.T) {
	tt := []struct {
		key      string
		value    any
		expected string
		err      error
	}{
		{
			key:      "TESTE_CUSTOM_TEXT_MARSHALER",
			value:    user{Name: "Luca Marques", Github: url.URL{Host: "github.com", Path: "/lucafmarques", Scheme: "https"}},
			expected: "Name: Luca Marques - GitHub: https://github.com/lucafmarques",
		},
		{
			key:      "TESTE_MARSHAL_TEXT_ERROR",
			value:    user{error: Error},
			expected: "",
			err:      Error,
		},
	}

	for _, tc := range tt {
		t.Run(tc.key, func(t *testing.T) {
			defer func() {
				err, ok := recover().(error)
				if ok {
					be.Equal(t, tc.err, err)
				} else {
					val := os.Getenv(tc.key)
					be.Equal(t, tc.expected, val)
				}
			}()

			env.MustSet(tc.key, tc.value)
		})
	}
}

type fail struct{}

type user struct {
	Name   string  `json:"name"`
	Github url.URL `json:"github"`
	error  error
}

func (u user) MarshalText() ([]byte, error) {
	if u.error != nil {
		return []byte{}, u.error
	}
	return []byte(fmt.Sprintf("Name: %s - GitHub: %s", u.Name, u.Github.String())), nil
}

func (u *user) UnmarshalText(data []byte) error {
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
