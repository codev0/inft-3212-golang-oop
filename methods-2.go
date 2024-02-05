package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func methods2() {
	c := Counter{}
	fmt.Println(c)
	// When you use a pointer receiver with a local variable that’s a value type, Go automatically converts it to a pointer type.
	// In this case, c.Increment() is converted to (&c).Increment().
	c.Increment()
	(&c).Increment()
	fmt.Println(c)
}
