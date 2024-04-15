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

var (
	Error      = errors.New("mock test error")
	TestURL, _ = url.Parse("https://github.com/lucafmarques")
)

type test[T any] struct {
	key      string
	env      string
	err      error
	want     T
	fallback T
}

func (t test[T]) Error() string {
	if t.err != nil {
		return fmt.Errorf("%w: %s", t.err, t.key).Error()
	}

	return ""
}

var TestGetTable = []func(t *testing.T){
	func(t *testing.T) {
		tc := test[bool]{
			key: "CONVERSION_ERROR",
			env: "truth",
			err: &strconv.NumError{Func: "ParseBool", Num: "truth", Err: strconv.ErrSyntax},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[bool](tc.key)
			be.Equal(t, tc.want, val)
			be.Equal(t, tc.Error(), err.Error())
		})
	},
	func(t *testing.T) {
		tc := test[bool]{
			key:  "UNSET",
			env:  "true",
			want: true,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[bool](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[string]{
			key:  "STRING",
			env:  "text",
			want: "text",
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[string](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[int]{
			key:  "INT",
			env:  "10",
			want: 10,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[int](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[int8]{
			key:  "INT8",
			env:  "10",
			want: 10,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[int8](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[int16]{
			key:  "INT16",
			env:  "10",
			want: 10,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[int16](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[int32]{
			key:  "INT32",
			env:  "10",
			want: 10,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[int32](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[int64]{
			key:  "INT64",
			env:  "10",
			want: 10,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[int64](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[uint]{
			key:  "UINT",
			env:  "10",
			want: 10,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[uint](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[uint8]{
			key:  "UINT8",
			env:  "10",
			want: 10,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[uint8](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[uint16]{
			key:  "UINT16",
			env:  "10",
			want: 10,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[uint16](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[uint32]{
			key:  "UINT32",
			env:  "10",
			want: 10,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[uint32](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[uint64]{
			key:  "UINT64",
			env:  "10",
			want: 10,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[uint64](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[float32]{
			key:  "FLOAT32",
			env:  "10.0",
			want: 10.0,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[float32](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[float64]{
			key:  "FLOAT64",
			env:  "10.0",
			want: 10.0,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[float64](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[complex64]{
			key:  "COMPLEX64",
			env:  "420+69i",
			want: complex64(complex(420, 69)),
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[complex64](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[complex128]{
			key:  "COMPLEX128",
			env:  "420+69i",
			want: complex(420, 69),
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[complex128](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]bool]{
			key:  "BOOL_SLICE",
			env:  "true,false",
			want: []bool{true, false},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]bool](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]string]{
			key:  "STRING_SLICE",
			env:  "split,text",
			want: []string{"split", "text"},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]string](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]int]{
			key:  "INT_SLICE",
			env:  "420,69",
			want: []int{420, 69},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]int](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]int8]{
			key:  "INT8_SLICE",
			env:  "127,69",
			want: []int8{127, 69},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]int8](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]int16]{
			key:  "INT16_SLICE",
			env:  "420,69",
			want: []int16{420, 69},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]int16](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]int32]{
			key:  "INT32_SLICE",
			env:  "420,69",
			want: []int32{420, 69},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]int32](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]int64]{
			key:  "INT64_SLICE",
			env:  "420,69",
			want: []int64{420, 69},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]int64](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]uint]{
			key:  "UINT_SLICE",
			env:  "420,69",
			want: []uint{420, 69},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]uint](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]uint8]{
			key:  "UINT8_SLICE",
			env:  "255,69",
			want: []uint8{255, 69},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]uint8](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]uint16]{
			key:  "UINT16_SLICE",
			env:  "420,69",
			want: []uint16{420, 69},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]uint16](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]uint32]{
			key:  "UINT32_SLICE",
			env:  "420,69",
			want: []uint32{420, 69},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]uint32](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]uint64]{
			key:  "UINT64_SLICE",
			env:  "420,69",
			want: []uint64{420, 69},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]uint64](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]float32]{
			key:  "FLOAT32_SLICE",
			env:  "420.0,69.",
			want: []float32{420.0, 69.},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]float32](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]float64]{
			key:  "FLOAT64_SLICE",
			env:  "420.0,69.",
			want: []float64{420.0, 69.},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]float64](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]complex64]{
			key:  "COMPLEX64_SLICE",
			env:  "420+69i,69+420i",
			want: []complex64{complex64(complex(420, 69)), complex64(complex(69, 420))},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]complex64](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[[]complex128]{
			key:  "COMPLEX128_SLICE",
			env:  "420+69i,69+420i",
			want: []complex128{complex(420, 69), complex(69, 420)},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[[]complex128](tc.key)
			be.AllEqual(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[string]{
			key:      "FALLBACK",
			fallback: "fallback",
			err:      env.ErrUnset,
		}

		val, err := env.Get(tc.key, tc.fallback)
		be.Equal(t, tc.fallback, val)
		be.Equal(t, tc.Error(), err.Error())
	},
	func(t *testing.T) {
		tc := test[user]{
			key:  "CUSTOM_TEXT_UNMARSHALER",
			env:  `{"name":"Luca Marques","github":"https://github.com/lucafmarques"}`,
			want: user{Name: "Luca Marques", Github: *TestURL},
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[user](tc.key)
			be.Equal(t, tc.want, val)
			be.NilErr(t, err)
		})
	},
	func(t *testing.T) {
		tc := test[fail]{
			key:  "UNIMPLEMENTED_TEXT_UNMARSHALER",
			env:  `{"name":"Luca Marques","github":"https://github.com/lucafmarques"}`,
			want: fail{},
			err:  env.ErrUnmarshaler,
		}

		t.Run(tc.key, func(t *testing.T) {
			t.Setenv(tc.key, tc.env)

			val, err := env.Get[fail](tc.key)
			be.Equal(t, tc.want, val)
			be.Equal(t, tc.Error(), err.Error())
		})
	},
}

func TestGet(t *testing.T) {
	for _, tc := range TestGetTable {
		tc(t)
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
