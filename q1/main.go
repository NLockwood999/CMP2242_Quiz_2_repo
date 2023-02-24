package main

import "fmt"

// created a type Dog with three fields
type Dog struct {
	Name   string
	Specie string
	Color  string
}

// created a type Fish with five fields
type Fish struct {
	Name     string
	Specie   string
	Color    string
	Size     int
	Venomous bool
}

func (d Dog) GetName() string {
	return d.Name
}

func (f Fish) GetName() string {
	return f.Name
}

func (d Dog) Species() string {
	return "Great Dane"
}

func (f Fish) Species() string {
	return "Deep Sea fish"
}

func (d Dog) Speak() string {
	return "Ruff Ruff"
}

func (f Fish) Speak() string {
	return "Blub Blub Blub"
}

//create an interfaces called Animals and implemented
//and creatde three method signatures
type Animal interface {
	GetName() string
	Speak() string
	Species() string
}

func main() {
	//create obj of type Dog and assigned
	dog := Dog{
		Name:   "Krypto",
		Specie: "Great_Dane",
		Color:  "White",
	}
	//create obj of type Fish and assigned
	fish := Fish{
		Name:     "Nemo",
		Specie:   "Clown fish",
		Color:    "Orange",
		Size:     5,
		Venomous: false,
	}
	PrintInfo(dog)
	PrintInfo(fish)
	//fmt.Println(dog.Name)
	//fmt.Println(fish.Name)

}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Speak(),
		"is named", a.GetName(),
		"and is of breed:", a.Species())
}

/*interfaces can be described as a named collection of methods*/
