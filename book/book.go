package book

import "fmt"

type Book struct {
	title  string
	author string
	pages  int
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n\n", b.title, b.author, b.pages)
}

func NewBook(title, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

// Compositions: Structures within structures, combines its attributes and methods from another structure
type Textbook struct {
	Book      // with this line, we are telling Golan to contain Book structure
	editorial string
	level     string
}

func (tb *Textbook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nLevel: %s\n\n", tb.title, tb.author, tb.pages, tb.editorial, tb.level)
}

func NewTextBook(title, author string, pages int, editorial, level string) *Textbook {
	return &Textbook{
		Book: Book{
			title:  title,
			author: author,
			pages:  pages,
		},
		editorial: editorial,
		level:     level,
	}
}
