package book

import "fmt"

/*
Instructions
Write a Go program to manage a library of books.
Each book has the following attributes:
Title (string)
Author (string)
Pages (int)
CopiesAvailable (int)
Part 1: Struct and Pointers
Create a Book struct with the attributes mentioned above.
Write a function CreateBook that takes the details of a book and returns a pointer to a new Book instance.
Part 2: Methods with Value Receivers
Write a method Display for the Book struct that prints the book details in a formatted way.
(Use a value receiver for this method.)
Part 3: Methods with Pointer Receivers
Write a method Borrow that:

Decrements the CopiesAvailable field by 1 if there are copies available.
Returns a message indicating whether the borrow was successful or not. (Use a pointer receiver for this method.)
Write a method ReturnBook that:

Increments the CopiesAvailable field by 1. (Use a pointer receiver for this method.)
Part 4: Pointer Manipulations
Create a function SwapTitles that takes two *Book pointers and swaps their titles.

*/

type Book struct {
	Title           string
	Author          string
	Pages           int
	CopiesAvailable int
}

func CreateBook(b Book) *Book {

	return &b
}

func (b Book) Display() {
	fmt.Printf("Title of the Book:\t%s\n", b.Title)
	fmt.Printf("Author of the Book:\t%s\n", b.Author)
	fmt.Printf("Pages in the Book:\t%d\n", b.Pages)
	fmt.Printf("CopiesAvailable of the Book:\t%d\n", b.CopiesAvailable)
}

func (b *Book) Borrow() string {
	if b.CopiesAvailable > 0 {
		b.CopiesAvailable = b.CopiesAvailable - 1
		return "Borrow was successful"
	} else {
		return "Borrow was not Successful"
	}
}

func (b *Book) ReturnBook() {
	b.CopiesAvailable = b.CopiesAvailable + 1
}

func SwapTitles(b *Book, c *Book) {
	str := b.Title
	b.Title = c.Title
	c.Title = str
}
