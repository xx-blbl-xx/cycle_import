package package2

import (
	"fmt"

	p1 "github.com/xx-blbl-xx/cycle_import/package_1"
)

var P2 = "Package 2 Variable"

func G() {
	fmt.Println("Hello from package2!")
	fmt.Println(p1.P1)
}
