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
	"strings"
)

var (
	// ErrUnset represents a missing ENV.
	ErrUnset error = errors.New("unset env")
	// ErrMarshaler represents a failure parsing any type to a string ENV.
	ErrMarshaler error = errors.New("type doesn't implement encoding.TextMarshaler")
	// ErrUnmarshaler represents a failure parsing an existing ENV to the any type.
	ErrUnmarshaler error = errors.New("type doesn't implement encoding.TextUnmarshaler")

	sep = ","
)

// Get attempts to retrieve an ENV and parse it to the given type T.
func Get[T any](key string, fallback ...T) (value T, err error) {
	var fb T
	if len(fallback) > 0 {
		fb = fallback[0]
	}

	env, ok := os.LookupEnv(key)
	if !ok {
		return fb, fmt.Errorf("%w: %v", ErrUnset, key)
	}

	value, err = parse(env, value)
	if err != nil {
		return fb, fmt.Errorf("%w: %v", err, key)
	}

	return value, nil
}

// MustGet panics if Get errors.
func MustGet[T any](key string) T {
	env, err := Get[T](key)
	if err != nil {
		panic(err)
	}

	return env
}

// Set attempts to set an ENV from any type T.
func Set[T any](key string, value T) error {
	switch v := any(value).(type) {
	case string, bool, []string, []bool,
		float32, float64, []float32, []float64,
		complex64, complex128, []complex64, []complex128,
		int, int8, int16, int32, int64, []int, []int8, []int16, []int32, []int64,
		uint, uint8, uint16, uint32, uint64, []uint, []uint8, []uint16, []uint32, []uint64:
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
func MustSet[T any](key string, value T) {
	err := Set(key, value)
	if err != nil {
		panic(err)
	}
}

func parse[T any](env string, value T) (T, error) {
	var err error

	switch v := any(&value).(type) {
	case *string:
		*v = env
	case *bool:
		*v, err = strconv.ParseBool(env)
	case *int:
		*v, err = strconv.Atoi(env)
	case *int8:
		d, er := strconv.ParseInt(env, 10, 8)
		*v, err = int8(d), er
	case *int16:
		d, er := strconv.ParseInt(env, 10, 16)
		*v, err = int16(d), er
	case *int32:
		d, er := strconv.ParseInt(env, 10, 32)
		*v, err = int32(d), er
	case *int64:
		*v, err = strconv.ParseInt(env, 10, 64)
	case *uint:
		d, er := strconv.ParseUint(env, 10, 64)
		*v, err = uint(d), er
	case *uint8:
		d, er := strconv.ParseUint(env, 10, 8)
		*v, err = uint8(d), er
	case *uint16:
		d, er := strconv.ParseUint(env, 10, 16)
		*v, err = uint16(d), er
	case *uint32:
		d, er := strconv.ParseUint(env, 10, 32)
		*v, err = uint32(d), er
	case *uint64:
		*v, err = strconv.ParseUint(env, 10, 64)
	case *float32:
		d, er := strconv.ParseFloat(env, 32)
		*v, err = float32(d), er
	case *float64:
		*v, err = strconv.ParseFloat(env, 64)
	case *complex64:
		d, er := strconv.ParseComplex(env, 64)
		*v, err = complex64(d), er
	case *complex128:
		*v, err = strconv.ParseComplex(env, 128)
	case *[]string:
		*v, err = parseSlice(env, "")
	case *[]bool:
		*v, err = parseSlice(env, false)
	case *[]int:
		*v, err = parseSlice(env, 0)
	case *[]int8:
		*v, err = parseSlice(env, int8(0))
	case *[]int16:
		*v, err = parseSlice(env, int16(0))
	case *[]int32:
		*v, err = parseSlice(env, int32(0))
	case *[]int64:
		*v, err = parseSlice(env, int64(0))
	case *[]uint:
		*v, err = parseSlice(env, uint(0))
	case *[]uint8:
		*v, err = parseSlice(env, uint8(0))
	case *[]uint16:
		*v, err = parseSlice(env, uint16(0))
	case *[]uint32:
		*v, err = parseSlice(env, uint32(0))
	case *[]uint64:
		*v, err = parseSlice(env, uint64(0))
	case *[]float32:
		*v, err = parseSlice(env, float32(0))
	case *[]float64:
		*v, err = parseSlice(env, float64(0))
	case *[]complex64:
		*v, err = parseSlice(env, complex64(0))
	case *[]complex128:
		*v, err = parseSlice(env, complex128(0))
	case encoding.TextUnmarshaler:
		err = v.UnmarshalText([]byte(env))
	default:
		err = ErrUnmarshaler
	}

	return value, err
}

func parseSlice[T any](env string, value T) ([]T, error) {
	items := strings.Split(env, sep)

	result := make([]T, len(items))
	errs := make([]error, len(items))

	for i, item := range items {
		result[i], errs[i] = parse(item, value)
	}

	return result, errors.Join(errs...)
}
