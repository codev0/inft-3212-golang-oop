package main

import "fmt"

type Course struct {
	author string
	price  int
}

// Course or *Course
func (c *Course) SetAuthor(name string) {
	c.author = name
}

// Course or *Course
func (c *Course) Author() string {
	return c.author
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func method() {
	// p := Person{
	// 	FirstName: "John",
	// 	LastName:  "Doe",
	// 	Age:       25,
	// }
	// p.String()
	// fmt.Println(p)

	c := new(Course)
	c.SetAuthor("Airat Zh")
	fmt.Println(c.Author())

	var course Course
	course.SetAuthor("John Doe")
	fmt.Println(course.Author())
}
