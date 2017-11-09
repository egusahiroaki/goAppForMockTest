package service

import "github.com/goSamples/ducktype/service/pet"

// Human is
type Human struct {
	name string
}

// NewHuman is constructor
func NewHuman(name string) *Human {
	return &Human{name: name}
}

// GetName is return human name
func (human Human) GetName() string {
	return human.name
}

// SetName is to set name
func (human *Human) SetName(name string) {
	human.name = name
}

// HavePet is to create Pet
func (human *Human) HavePet(p pet.Pet) {
	p.Cry()
}

// HavePet is to create Pet
// func (human *Human) HavePet() {
// 	dog := new(pet.Dog)
// 	dog.Voice()
// }
