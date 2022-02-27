package main

import (
	"os"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-1-aybukeertekin/library"
)

func addBooks() {
	library.AddBook("Don Quixote")
	library.AddBook("The Odyssey")
	library.AddBook("The Brothers Karamazov")
	library.AddBook("Anna Karenina")
	library.AddBook("The Iliad")
}

func main() {
	args := os.Args
	command, error := library.CheckUserArguments(args)
	//If user arguments are correct, add books to the list and
	//apply commands.
	if error == nil {
		addBooks()
		library.PrintLineMessage("Books:")
		switch command {
		case library.SearchCommand:
			library.SearchBooks(args[2])
		case library.ListCommand:
			library.ListBooks()
		}
	}
}
