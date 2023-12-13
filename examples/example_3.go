package main

import (
	"fmt"

	"github.com/lucafmarques/env"
)

type config struct {
	intValue    int
	boolValue   bool
	stringValue string
	floatValue  float64
}

func (c *config) Build() (err error) {
	// if every value has a fallback, you can omit the type entirely
	c.intValue, err = env.Get("INT", 10)
	c.boolValue, err = env.Get("BOOL", false)
	c.floatValue, err = env.Get("FLOAT", 10.)
	c.stringValue, err = env.Get("STRING", "teste")

	return
}

func main() {
	var c config
	err := c.Build()
	fmt.Println(c.boolValue, c.stringValue, c.intValue, c.floatValue, err)
}
