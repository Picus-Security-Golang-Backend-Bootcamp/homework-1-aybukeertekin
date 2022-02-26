package library

import (
	"strings"
)

var books []string = make([]string, 10)
var size int = 0

func AddBook(newBook string) {
	books[size] = newBook
	size++
}

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

func ListBooks() {
	for index := 0; index < size; index++ {
		PrintLineMessage("\t" + books[index])
	}
}
