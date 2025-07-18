package package1

import (
	"fmt"

	p2 "cycle_import_0/package_2"
)

var P1 = "Package 1 Variable"

func F() {
	fmt.Println("Hello from package1!")
	fmt.Println(p2.P2)
}
