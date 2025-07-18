package package1

import (
	p2 "cycle_import/package_2"
	"fmt"
)

var P1 = "Package 1 Variable"

func F() {
	fmt.Println("Hello from package1!")
	fmt.Println(p2.P2)
}
