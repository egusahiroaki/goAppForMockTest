package pet

import (
	"fmt"
)

// Cat is
type Cat struct {
	name string
}

// NewCat is
func NewCat(name string) *Cat {
	return &Cat{name: name}
}

// Cry is
func (cat *Cat) Cry() {
	// fmt.Printf("%v is mewmew\n", cat.name)
	fmt.Printf("%v: mew\n", cat.name)
}
