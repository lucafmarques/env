# `env`: read and write envs with generics
![Coverage](https://img.shields.io/badge/coverage-43.4%25-yellow)
[![Go Reference](https://pkg.go.dev/badge/github.com/lucafmarques/env.svg)](https://pkg.go.dev/github.com/lucafmarques/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucafmarques/env)](https://goreportcard.com/report/github.com/lucafmarques/env)

`env` allows parsing environment values directly from and into `string`, `bool`, `int`, `float64` and any type that implement the [`encoding.TextMarshaler`](https://pkg.go.dev/encoding#TextMarshaler) and/or [`encoding.TextUnmarshaler`](https://pkg.go.dev/encoding#TextUnmarshaler) interfaces.

## Example

```go
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
```

Running it:
```sh
$ INTEGER=10 STRING="this is an example" LOG_FORMAT="DEBUG,example" go run main.go
10 this is an example {DEBUG example}
0001-01-01 00:00:00 +0000 UTC unset env
{INFO rewritten}
$ env | grep LOG_FORMAT
LOG_FORMAT=INFO,rewritten
```
