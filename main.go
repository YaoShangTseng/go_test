package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
}

func foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}

func bar() {
	fmt.Println("BAR")
}
