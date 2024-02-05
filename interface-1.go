package main

import "fmt"

// Quacker is an interface that describes anything that can quack.
type Quacker interface {
	quack()
}

// Duck is a type that satisfies the Quacker interface.
type Duck struct{}

func (d Duck) quack() {
	fmt.Println("Quack, quack!")
}

// Frog is another type that satisfies the Quacker interface.
type Frog struct{}

func (p Frog) quack() {
	fmt.Println("The frog imitates a duck's quack.")
}

// makeItQuack accepts anything that satisfies the Quacker interface.
func makeItQuack(q Quacker) {
	q.quack()
}

func interface1() {
	duck := Duck{}
	frog := Frog{}

	makeItQuack(duck) // Output: Quack, quack!
	makeItQuack(frog) // Output: The person imitates a duck's quack.
}
