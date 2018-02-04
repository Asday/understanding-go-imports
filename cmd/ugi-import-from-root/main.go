package main

import (
	"fmt"

	// Imports everything from the root.
	"github.com/asday/understanding-go-imports"
)

func main() {
	fmt.Println("ugi-import-from-root")

	// Defined in `a.go`.
	b.B()
}
