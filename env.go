// Package env provides a simple and flexible way of fetching ENVs and converting them to Go types.
//
// The flexibility of the package stems from the Builder interface, which allows ENVs to be converted to user-defined types.
package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type constraint interface {
	string | bool | int | float64
}

// Builder is an interface for a type that can be build itself from a string ENV.
type Builder interface {
	Build(env string) (any, error)
}

var (
	// ErrUnset represents a missing ENV.
	ErrUnset error = errors.New("unset env")
	// ErrBuilder represents a failure converting an existing ENV to the specified special type.
	ErrBuilder error = errors.New("type doesn't implement Builder")
)

// Get attempts to retrieve an ENV and convert it to the specified type.
// Get can natively convert envs to the types string, bool, int, float64 and any type that implements Builder.
func Get[T constraint | any](key string, fallback ...T) (value T, err error) {
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

	switch v := any(value).(type) {
	case string:
		value = any(env).(T)
	case bool:
		v, err = strconv.ParseBool(env)
		value = any(v).(T)
	case int:
		v, err = strconv.Atoi(env)
		value = any(v).(T)
	case float64:
		v, err = strconv.ParseFloat(env, 64)
		value = any(v).(T)
	case Builder:
		res, e := v.Build(env)
		value = res.(T)
		err = e
	default:
		err = ErrBuilder
	}

	if err != nil {
		value = fb
	}

	return value, err
}

// MustGet panics if Get fails.
func MustGet[T constraint | any](key string) T {
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
