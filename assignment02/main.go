package main

import (
	"Module/book"
	"fmt"
)

func main() {
	var b, k book.Book
	b = book.Book{"KingMaker", "Maheshwar", 1500, 0}
	k = book.Book{"Lion King", "Sanjay", 2000, 25}
	b.Display()
	fmt.Println(b.Borrow())
	b.Display()
	b.ReturnBook()
	b.Display()
	book.SwapTitles(book.CreateBook(b), book.CreateBook(k))
	b.Display()
	k.Display()
}
