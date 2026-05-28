package main

import "fmt"

type Book struct {
	Title       string
	Author      string
	pages       int
	isAvailable bool
}

func newBook(title string, author string, pages int) Book {
	if pages <= 0 {
		pages = 1
	}

	return Book{
		Title:       title,
		Author:      author,
		pages:       pages,
		isAvailable: true,
	}
}

func (b Book) Info() {
	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
	fmt.Println("Pages:", b.pages)
	fmt.Println("Available:", b.isAvailable)
}

func (b *Book) Borrow() {
	if b.isAvailable {
		b.isAvailable = false
		fmt.Println("Book borrowed")
	} else {
		fmt.Println("Book is not available")
	}
}

func (b *Book) ReturnBook() {
	b.isAvailable = true
	fmt.Println("Book returned")
}

func (b Book) GetPages() int {
	return b.pages
}

func (b *Book) SetPages(p int) {
	if p <= 0 {
		fmt.Println("Error: pages must be positive")
		return
	}

	b.pages = p
}

func main() {
	book := newBook("Go Basics", "Ali", 200)

	book.Info()

	fmt.Println()

	book.Borrow()
	book.Borrow()
	book.ReturnBook()

	fmt.Println()

	book.SetPages(250)
	fmt.Println("Pages:", book.GetPages())

	fmt.Println()

	book.Info()
}
