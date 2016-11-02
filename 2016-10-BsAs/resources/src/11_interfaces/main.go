package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Talk()
}

type Dog struct {
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
	var dog Animal = &Dog{Name: "Pedro", Age: 34}

	dog.Eat()
	dog.Talk()

	fmt.Println("")

	var lion Animal = &Lion{Name: "Tom", Age: 34}

	lion.Eat()
	lion.Talk()

	/*
		animals := []Animal{&Dog{Name: "Pedro", Age: 34}, &Lion{Name: "Tom", Age: 34}}

		for _, animal := range animals {
			animal.Eat()
			animal.Talk()
		}
	*/
}
