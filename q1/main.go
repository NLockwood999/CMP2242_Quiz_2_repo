package main

import "fmt"

// create a class called dog
type Dog struct {
	Name   string
	Specie string
	Color  string
}

type Fish struct {
	Name     string
	Specie   string
	Color    string
	Size     int
	Venomous bool
}

func (d Dog) Speak() string {
	return "Ruff Ruff"
}

func (d Dog) NumberOflegs() int {
	return 4
}

func (f Fish) Speak() string {
	return "Blub Blub Blub"
}

func (f Fish) NumberOfSpines() int {
	return 1
}

func (d Dog) Species() string {
	return "Great Dane"
}

func (d Dog) GetName() string {
	return d.Name
}

func (f Fish) GetName() string {
	return f.Name
}
func (f Fish) Species() string {
	return "Deep Sea fish"
}

//construct for interfaces
type Animal interface {
	GetName() string
	Speak() string
	Species() string
}

func main() {
	//create obj cat
	dog := Dog{
		Name:   "Krypto",
		Specie: "Great Dane",
		Color:  "White",
	}
	fish := Fish{
		Name:     "Nemo",
		Specie:   "Clown fish",
		Color:    "Orange",
		Size:     5,
		Venomous: false,
	}
	PrintInfo(dog)
	PrintInfo(fish)
	//fmt.Println(doggo.Breed)

}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Speak(),
		"is named", a.GetName(),
		"and is of breed:", a.Species())
}
