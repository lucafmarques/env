package main

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/lucafmarques/env"
)

type example struct {
	A string
	B string
}

func (e *example) UnmarshalText(data []byte) error {
	v := bytes.Split(data, []byte(","))
	if len(v) < 2 {
		return errors.New("missing values in env")
	}
	e.A = string(v[0])
	e.B = string(v[1])
	return nil
}

func main() {
	fmt.Println(env.MustGet[int]("INTEGER"))
	fmt.Println(env.MustGet[string]("STRING"))
	fmt.Println(env.MustGet[example]("LOG_FORMAT"))
}
