package library

import (
	"strings"
)

var books []string = make([]string, 10) //book list
var size int = 0                        //book list size

//This function adds a new book name to the book list.
func AddBook(newBook string) {
	books[size] = newBook
	size++
}

//This function searchs book list and look if a book exists
//with the given word or characters and prints their name.
//Prints error message if no book exists.
func SearchBooks(searchWord string) {
	bookFound := false
	searchWord = strings.ToLower(searchWord)
	for index := 0; index < size; index++ {
		bookName := books[index]
		if strings.Contains(strings.ToLower(bookName), searchWord) {
			PrintLineMessage("\t" + bookName)
			bookFound = true
		}
	}
	if !bookFound {
		PrintLineMessage("\tNo books are found with given argument.")
	}
}

//This function prints book list.
func ListBooks() {
	for index := 0; index < size; index++ {
		PrintLineMessage("\t" + books[index])
	}
}
