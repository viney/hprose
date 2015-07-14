package main

import (
	"fmt"

	hprose "github.com/hprose/hprose-go"
)

type Client struct {
	Say   func(string)
	Hello func() string
}

func main() {
	client := hprose.NewClient("http://127.0.0.1:7080")

	c := &Client{}

	client.UseService(c)

	c.Say("world")
	fmt.Println(c.Hello())
}
