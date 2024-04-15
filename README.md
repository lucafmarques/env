# `env`: read and write envs with generics
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
![Coverage](https://img.shields.io/badge/coverage-43.4%25-yellow)
[![Go Reference](https://pkg.go.dev/badge/github.com/lucafmarques/env.svg)](https://pkg.go.dev/github.com/lucafmarques/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucafmarques/env)](https://goreportcard.com/report/github.com/lucafmarques/env)

`env` allows parsing environment values directly to and from native Go types and any custom types that implement the `encoding.TextUnmarshaler` and/or `encoding.TextMarshaler` interfaces.

```
go get github.com/lucafmarques/env
```

`env` works with the following native types out of the box:
- `bool`
- `string` 
- `int`, variants and aliases
- `uint`, variants and aliases
- `float32` and `float64`
- `complex64` and `complex128`
- `[]T` where `T` is any of the above types
- `encoding.TextUnmarshaler`

Custom types must implement [`encoding.TextUnmarshaler`](https://pkg.go.dev/encoding#TextUnmarshaler) and/or [`encoding.TextMarshaler`](https://pkg.go.dev/encoding#TextMarshaler) to work with `env`.

## Examples
<table>
<tr>
<th><code>native.go</code></th>
<th><code>custom.go</code></th>
</tr>
<tr>
<td>
  
```go
package main

import (
    "time"

    "github.com/lucafmarques/env"
)

func main() {
    // MustGet panics if the env isn't set
    intV := env.MustGet[int]("INTEGER")
    strV := env.MustGet[string]("STRING")
    timeV, err := env.Get[time.Time]("TIME")
    // ...
    env.MustSet("COMPLEX", complex128(420))
    err = env.Set("UINT32", uint32(69))
}
```
</td>
<td>

```go
package main

import (
    "bytes"
    "fmt"

    "github.com/lucafmarques/env"
)

type log struct {
    Format string
    Level string
}

func (l log) MarshalText() ([]byte, error) {
    s := fmt.Sprintf("%s,%s", e.Format, e.Level) 
    return []byte(s), nil
}

func (l *log) UnmarshalText(d []byte) error {
    v := bytes.Split(d, []byte(","))
    if len(v) != 2 {
        return fmt.Errorf("can't unrmarshal")
    }
    l.Format = string(v[0])
    l.Level = string(v[1])
    return nil
}

func main() {
    logV := env.MustGet[log]("LOG_FORMAT"))
    logV = log{Format: "JSON", Level: "INFO"}
    err := env.Set("LOG_FORMAT", logV)	
}
```
</td>
</tr>
</table>
