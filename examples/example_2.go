package main

import (
	"fmt"

	"github.com/lucafmarques/env"
)

func main() {
	intValue, _ := env.Get("INT", 10)
	boolValue, _ := env.Get("BOOL", false)
	floatValue, _ := env.Get("FLOAT", 10.)
	stringValue, _ := env.Get("STRING", "teste")

	fmt.Println(boolValue, stringValue, intValue, floatValue)
}
