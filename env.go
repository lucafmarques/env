// Package env provides a simple and flexible way of interacting with ENVs directly from and to Go types.
//
// The flexibility of the package stems from the [encoding.TextMarshaler] and [encoding.TextUnmarshaler] interfaces, which allows ENVs to be parsed to non-native and user-defined types.
//
// [encoding.TextMarshaler]: https://pkg.go.dev/encoding#TextMarshaler
// [encoding.TextUnmarshaler]: https://pkg.go.dev/encoding#TextUnmarshaler
package env

import (
	"encoding"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	// ErrUnset represents a missing ENV.
	ErrUnset error = errors.New("unset env")
	// ErrMarshaler represents a failure parsing any type to a string ENV.
	ErrMarshaler error = errors.New("type doesn't implement encoding.TextMarshaler")
	// ErrUnmarshaler represents a failure parsing an existing ENV to the any type.
	ErrUnmarshaler error = errors.New("type doesn't implement encoding.TextUnmarshaler")
)

// Get attempts to retrieve an ENV and parse it to the given type T.
func Get[T bool | string | int | float64 | any](key string, fallback ...T) (value T, err error) {
	var fb T
	if len(fallback) > 0 {
		fb = fallback[0]
	}

	env, ok := os.LookupEnv(key)
	if !ok {
		err = ErrUnset
		if len(fallback) > 0 {
			value = fb
		}

		return value, err
	}

	switch v := any(&value).(type) {
	case *string:
		value = any(env).(T)
	case *bool:
		*v, err = strconv.ParseBool(env)
		value = any(*v).(T)
	case *int:
		*v, err = strconv.Atoi(env)
		value = any(*v).(T)
	case *float64:
		*v, err = strconv.ParseFloat(env, 64)
		value = any(*v).(T)
	case encoding.TextUnmarshaler:
		err = v.UnmarshalText([]byte(env))
	default:
		err = ErrUnmarshaler
	}

	if err != nil {
		value = fb
	}

	return value, err
}

// MustGet panics if Get errors.
func MustGet[T bool | string | int | float64 | any](key string) T {
	env, err := Get[T](key)
	if err != nil {
		panic(fmt.Errorf("env $%s [%T]: %w", key, env, err))
	}

	return env
}

// Set attempts to set an ENV from any type T.
func Set[T bool | string | int | float64 | any](key string, value T) error {
	switch v := any(value).(type) {
	case bool, string, int, float64:
		return os.Setenv(key, fmt.Sprint(v))
	case encoding.TextMarshaler:
		val, err := v.MarshalText()
		if err != nil {
			return err
		}
		return os.Setenv(key, string(val))
	default:
		return ErrMarshaler
	}
}

// MustSet panics if Set errors.
func MustSet[T bool | string | int | float64 | any](key string, value T) {
	err := Set(key, value)
	if err != nil {
		panic(err)
	}
}
