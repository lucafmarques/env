package main

import (
	"fmt"

	"github.com/lucafmarques/env"
)

func main() {
	fmt.Println(env.MustGet[string]("STRING"))
	fmt.Println(env.MustGet[int]("INTEGER"))
}
