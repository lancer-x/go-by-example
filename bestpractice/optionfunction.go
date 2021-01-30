package main

import (
	"fmt"
	"net/http"
)

var ServerURI = "aaaa"

type Modifier func(r *http.Request)

type Client struct {
	modifiers []Modifier
}

func NewAPIRequest() (*http.Request, error) {
	return http.NewRequest(http.MethodGet, ServerURI, nil)
}

func NewClient(opts ...Modifier) *Client {
	c := &Client{}
	c.modifiers = append(c.modifiers, opts...)
	return c
}

// 获取
func (c *Client) Request() (*http.Request, error) {
	req, err := NewAPIRequest()
	if err != nil {
		return nil, err
	}
	for _, mod := range c.modifiers {
		mod(req)
	}
	return req, nil
}

func withPrint(msg string) Modifier {
	return func(r *http.Request) {
		fmt.Println("print a new " + msg)
	}
}

func main() {
	c := NewClient(withPrint("request"))
	c.Request()

	p3 := NewPersonWithOption(WithName("hahah"), WithAge(30))
	fmt.Println(p3)
}

type Person struct {
	name   string
	age    int
	salary int
	height int
}

/**
 */
func NewPerson(name string, age, salary, height int) *Person {
	return &Person{
		name:   name,
		age:    age,
		salary: salary,
		height: height,
	}
}

func NewPersonByName(name string) *Person {
	return &Person{
		name: name,
	}
}

type option func(*Person)

/**
 */
func NewPersonWithOption(options ...option) *Person {
	person := &Person{}
	for _, opt := range options {
		opt(person)
	}
	return person
}

/**
 */
func WithName(name string) option {
	return func(p *Person) {
		p.name = name
	}
}

func WithAge(age int) option {
	return func(p *Person) {
		p.age = age
	}
}
