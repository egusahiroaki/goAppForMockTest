package pet

import (
	"fmt"
)

// Dog is
type Dog struct {
	name string
}

// NewDog is
func NewDog(name string) *Dog {
	return &Dog{name: name}
}

// Cry is
func (dog *Dog) Cry() {
	// fmt.Printf("%v is vowvow\n", dog.name)
	fmt.Printf("%v: woof!!\n", dog.name)
}
