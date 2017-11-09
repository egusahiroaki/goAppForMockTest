package main

import (
	"fmt"

	"github.com/goSamples/ducktype/service"
	"github.com/goSamples/ducktype/service/pet"
)

func main() {
	human := service.NewHuman("alice")
	fmt.Println(human.GetName())

	dog := pet.NewDog("pochi")
	human.HavePet(dog)

	cat := pet.NewCat("tama")
	human.HavePet(cat)
}
