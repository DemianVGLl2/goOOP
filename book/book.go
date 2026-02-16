package book

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.Title, b.Author, b.Pages)
}

func NewBook(title, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

// Compositions: Structures within structures, combines its attributes and methods from another structure
type Textbook struct {
	Book      // with this line, we are telling Golan to contain Book structure
	Editorial string
	Level     string
}

func (tb *Textbook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nLevel: %s\n", tb.Title, tb.Author, tb.Pages, tb.Editorial, tb.Level)
}

func NewTextBook(title, author string, pages int, editorial, level string) *Textbook {
	return &Textbook{
		Book: Book{
			Title:  title,
			Author: author,
			Pages:  pages,
		},
		Editorial: editorial,
		Level:     level,
	}
}
