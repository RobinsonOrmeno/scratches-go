package main

import (
	"books/animal"
)

func main() {
	//	myBook := book.NewBook("Moby Dick", "Herman Melville", 300)
	//	myBook.PrintInfo()
	//
	//	myBook.SetTitle("Moby Dick (Edición especial)")
	//	fmt.Println(myBook.GetTitle())
	//	myBook.PrintInfo()
	//	fmt.Println("-------------")
	//	//	COMPOSISICÓN, una nueva estructura textbook que hereda de book parametros.
	//	myTextBook := book.NewTextBook("Comunicación", "Jaime Gamarra", 261, "Santillana SAC", "Secundaria")
	//	myTextBook.PrintInfo()
	//	fmt.Println("-------------")
	//
	//	book.Print(myBook)
	//	book.Print(myTextBook)

	//fmt.Println("-------------")
	//
	//miPerro := animal.Perro{Nombre: "Albert"}
	//miGato := animal.Gato{Nombre: "Joaquin"}
	//animal.HacerSonido(&miPerro)
	//animal.HacerSonido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Alberto"},
		&animal.Gato{Nombre: "Ramon"},
		&animal.Perro{Nombre: "Jose"},
		&animal.Gato{Nombre: "Joaquin"},
	}
	for _, animal := range animales {
		animal.Sonido()
	}
}
