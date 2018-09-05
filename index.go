package main

import (
	"fmt"
	"reflect"
)

type Cat interface {
	Meow()
}

type Lion struct{}

func (l Lion) Meow() {
	fmt.Println("Roar")
}

type CatFactory func() Cat

var cf CatFactory = func() Cat {
	return CreateLion()
}

func CreateLion() Cat {
	return Lion{}
}

func main() {
	lion := CreateLion()
	fmt.Println(reflect.TypeOf(lion))
	lion.Meow()

	catlion := cf()
	catlion.Meow()
	fmt.Println(reflect.TypeOf(catlion))
	catlion.Meow()
}
