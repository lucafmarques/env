// Package env provides a simple and flexible way of fetching ENVs and converting them to Go types.
//
// The flexibility of the package stems from the Builder interface, which allows ENVs to be converted to user-defined types.
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
	// ErrUnmarshallText represents a failure converting an existing ENV to the specified special type.
	ErrUnmarshalText error = errors.New("type doesn't implement encoding.TextUnmarshaler")
)

// Get attempts to retrieve an ENV and convert it to the specified type.
// Get can natively convert envs to the types string, bool, int, float64 and any type that implements Builder.
func Get[T any](key string, fallback ...T) (value T, err error) {
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
		err = ErrUnmarshalText
	}

	if err != nil {
		value = fb
	}

	return value, err
}

// MustGet panics if Get fails.
func MustGet[T any](key string) T {
	env, err := Get[T](key)
	if err != nil {
		panic(fmt.Errorf("env $%s [%T]: %w", key, env, err))
	}

	return env
}

// Set is a convenience function for os.Setenv.
func Set(key, value string) error {
	return os.Setenv(key, value)
}

// MustSet panics if Set fails.
func MustSet(key, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		panic(err)
	}
}
