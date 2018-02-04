package main

import (
	"fmt"

	// Imports everything from `/lib`.
	"github.com/asday/understanding-go-imports/lib"
)

func main() {
	fmt.Println("ugi-import-from-subfolder")

	// Defined in `/lib/c.go`.
	d.D()
}
