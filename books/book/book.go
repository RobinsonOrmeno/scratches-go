package book

import "fmt"

// Book es una estructura que debe estar en un package aparte llamado book ya que tiene a su constructor y sus metodos setter y getter
type Book struct {
	title  string
	author string
	pages  int
}

func NewBook(title, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) PrintInfo() {
	fmt.Printf("Titulo: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}

// Composición de un textbook a partir de un book
type TextBook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book: Book{
			title:  title,
			author: author,
			pages:  pages,
		},
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Titulo: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nNivel: %s\n", b.title, b.author, b.pages, b.editorial, b.level)
}

// Polimorfismo, capacidad de objetos o clases para responder el mismo mensaje o metodo, en go se implementa a través de interface
// un interface es un conjunto de metodos que una estructura debe implementar para satisfacer la interface
type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}
