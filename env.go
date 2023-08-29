package env

import (
	"fmt"
	"os"
	"strconv"
)

type constraint interface {
	string | int | float64 | bool
}

func Get[T constraint](key string, fallback ...T) (value T, ok bool) {
	var (
		v   any
		err error
	)

	env := os.Getenv(key)
	if env == "" && len(fallback) > 0 {
		return fallback[0], false
	}

	switch any(value).(type) {
	case string:
		v, ok = any(env).(T)
	case int:
		v, err = strconv.Atoi(env)
		if err == nil {
			ok = true
		}
	case float64:
		v, err = strconv.ParseFloat(env, 64)
		if err == nil {
			ok = true
		}
	case bool:
		v, err = strconv.ParseBool(env)
		if err == nil {
			ok = true
		}
	}

	return v.(T), ok
}

func MustGet[T constraint](key string) T {
	env, ok := Get[T](key)
	if !ok {
		panic(fmt.Sprintf("env cannot be cast to type %T", env))
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
