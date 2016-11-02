package main

import (
	"fmt"
)


type Animal struct {
}

func (p *Animal) Eat() {
	fmt.Printf("Is eating... \n")
}

func (p *Animal) Talk() {
	fmt.Printf("talking... \n")
}

type Dog struct {
	Animal
	Name string
	Age  int
}

func (p *Dog) Eat() {
	fmt.Printf("%s is eating... \n", p.Name)
}

func (p *Dog) Talk() {
	fmt.Printf("%s is barking... \n", p.Name)
}

type Lion struct {
	Animal
	Name string
	Age  int
}

func (p *Lion) Eat() {
	fmt.Printf("%s is eating... \n", p.Name)
}
func (p *Lion) Talk() {
	fmt.Printf("%s is roaring... \n", p.Name)
}

func main() {
	dog := &Dog{Name: "Pedro", Age: 34}
	dog.Eat()
	dog.Talk()

	lion := &Lion{Name: "Tom", Age: 34}
	lion.Eat()
	lion.Talk()
}
