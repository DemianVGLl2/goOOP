package main

import (
	"fmt"
	"github/DemianVGLl2/goOOP/book"
)

func main() {
	// Golang doesn't implement OOP
	// Classes -> Structs
	// Methods -> Methods over structures
	// Encapsulation -> There is no modifiers such as private or public,
	// but if we have a function with a capital letter it will be public, otherwise it will be private
	// Inheritance -> Composition: Structures within structures
	// Polymorphism -> Interfaces

	/*var myBook = book.Book{
		title:  "Los cinco lenguajes del amor",
		author: "Gary Chapman",
		pages:  200,
	}*/

	// myBook.PrintInfo()

	var myBook2 = book.NewBook("Ensayo sobre la ceguera", "Jose Saramago", 500)

	myBook2.PrintInfo()

	var myTextBook = book.NewTextBook("I don't know", "Someone", 1000, "Somewhat", "dunno")

	myTextBook.PrintInfo()

	fmt.Println(myTextBook.GetTitle())
}
