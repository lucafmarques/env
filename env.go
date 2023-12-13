package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type constraint interface {
	string | int | float64 | bool
}

func Get[T constraint](key string, fallback ...T) (value T, err error) {
	env := os.Getenv(key)
	if env == "" {
		if len(fallback) > 0 {
			return fallback[0], nil
		}

		return value, errors.New("env not set")
	}
	switch any(value).(type) {
	case string:
		value = any(env).(T)
	case int:
		var v any
		v, err = strconv.Atoi(env)
		value = v.(T)
	case float64:
		var v any
		v, err = strconv.ParseFloat(env, 64)
		value = v.(T)
	case bool:
		var v any
		v, err = strconv.ParseBool(env)
		value = v.(T)
	}

	return
}

func MustGet[T constraint](key string) (env T) {
	env, err := Get[T](key)
	if err != nil {
		panic(fmt.Errorf("env $%s error: %w", key, err))
	}

	return env
}

func Set(key, value string) error {
	return os.Setenv(key, value)
}

func MustSet(key, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		panic(err)
	}
}
