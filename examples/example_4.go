package main

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/lucafmarques/env"
)

type log struct {
	Format string
	Prefix string
}

func (e log) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%s,%s", e.Format, e.Prefix)), nil
}

func (e *log) UnmarshalText(data []byte) error {
	v := bytes.Split(data, []byte(","))
	if len(v) < 2 {
		return errors.New("missing values in env")
	}
	e.Format = string(v[0])
	e.Prefix = string(v[1])

	return nil
}

func main() {
	fmt.Println(env.MustGet[int]("INTEGER"), env.MustGet[string]("STRING"), env.MustGet[log]("LOG_FORMAT"))
	fmt.Println(env.Get[time.Time]("TIME"))
	env.MustSet("LOG_FORMAT", log{Format: "INFO", Prefix: "rewritten"})
	fmt.Println(env.MustGet[log]("LOG_FORMAT"))
}
