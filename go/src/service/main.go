package main

import (
	"fmt"
	"net/http"

	hprose "github.com/hprose/hprose-go"
)

type Person struct {
	Content string
}

func (p *Person) Say(content string) {
	p.Content = content
}

func (p *Person) Hello() string {
	return "Hello " + p.Content + "!"
}

func main() {
	service := hprose.NewHttpService()

	service.AddMethods(&Person{})

	if err := http.ListenAndServe(":7080", service); err != nil {
		fmt.Println(err)
		return
	}
}
