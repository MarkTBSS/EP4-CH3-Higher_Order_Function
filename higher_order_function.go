package main

import (
	"fmt"
)

type Decorator func(s string) error

func outer(next Decorator) Decorator {
	return func(c string) error {
		fmt.Println("outer")
		r := c + " added"
		return next(r)
	}
}

func inner(s string) error {
	fmt.Println("inner", s)
	return nil
}

func main() {
	decorator := outer(inner)
	fmt.Println(decorator("decorator"))
	//w := wrapped("forth")
	//fmt.Println("end result:", w)
}
