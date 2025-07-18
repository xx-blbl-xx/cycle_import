package main

import (
	p1 "cycle_import/package_1"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	p1.F()
	fmt.Println(p1.P1)
}
