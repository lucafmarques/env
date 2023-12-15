package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/lucafmarques/env"
)

type example struct {
	A string
	B string
}

func (e example) Build(env string) (any, error) {
	v := strings.Split(env, ",")
	return example{
		A: v[0],
		B: v[1],
	}, nil
}

func main() {
	v, _ := env.Get[example]("LOG_FORMAT", example{})
	fmt.Println(reflect.TypeOf(v) == reflect.TypeOf(example{}))
	fmt.Println(env.Get[example]("LOG_FORMAT", example{}))
	fmt.Println(env.MustGet[string]("STRING"))
	fmt.Println(env.MustGet[int]("INTEGER"))
}
