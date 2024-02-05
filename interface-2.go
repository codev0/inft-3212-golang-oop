package main

import "fmt"

// Define an interface that abstracts the behavior.
type Greeter interface {
	Greet() string
}

// Define a concrete struct that implements the Greeter interface.
type EnglishGreeter struct {
	name    string
	country string
}

// Implement the Greet method for EnglishGreeter.
func (EnglishGreeter) Greet() string {
	return "Hello!"
}

// Define another concrete struct.
type SpanishGreeter struct{}

// Implement the Greet method for SpanishGreeter.
func (SpanishGreeter) Greet() string {
	return "¡Hola!"
}

// greet function accepts a Greeter interface, demonstrating "Accept interfaces".
func greet(g Greeter) {
	fmt.Println(g.Greet())
}

// newEnglishGreeter function returns a concrete struct, demonstrating "Return structs".
func newEnglishGreeter(name string, country string) EnglishGreeter {
	return EnglishGreeter{
		name:    name,
		country: country,
	}
}

// newSpanishGreeter function returns a concrete struct, demonstrating "Return structs".
func newSpanishGreeter() SpanishGreeter {
	return SpanishGreeter{}
}

func interface2() {
	englishGreeter := newEnglishGreeter("John", "USA")
	spanishGreeter := newSpanishGreeter()

	fmt.Println(englishGreeter.country, englishGreeter.name)

	greet(englishGreeter)
	greet(spanishGreeter)
}
