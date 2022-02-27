package library

import (
	"fmt"
)

const SearchCommand = "search"
const ListCommand = "list"

//Line message is used in many place
//Prints a line message.
func PrintLineMessage(message string) {
	fmt.Println(message)
}
