package library

import (
	"errors"
	"fmt"
	"path"
	"strings"
)

//Errors to return
var ErrNoArgument = errors.New("No user argument found!")
var ErrWrongCommand = errors.New("Unknown command type was used!")
var ErrWrongSearchCommand = errors.New("One argument was given for search command!")

//Check args to check if there is a mistake with the given arguments
func CheckUserArguments(args []string) (string, error) {
	projectName := path.Base(args[0])
	if err := checkSize(args); err != nil {
		printNoArgument()
		printOptions(projectName)
		return "", err
	} else {
		command, err := checkCommand(args)
		if err != nil {
			switch err {
			case ErrWrongCommand:
				printWrongCommandUsed()
			case ErrWrongSearchCommand:
				printWrongSearchCommand()
			}
			printOptions(projectName)
		}
		return command, err
	}
}

//Check argument size
//Should be at least 2.
func checkSize(args []string) error {
	if len(args) == 1 {
		return ErrNoArgument
	}
	return nil
}

//Check argument size if search command is needed.
//Should be at least 3
func checkSearchCommandArgumentSize(args []string) (string, error) {
	if len(args) <= 2 {
		return SearchCommand, ErrWrongSearchCommand
	}
	return SearchCommand, nil
}

//Check if second argument contains one of list or
//search commands
func checkCommand(args []string) (string, error) {
	command := args[1]
	if strings.EqualFold(command, "list") {
		return ListCommand, nil
	} else if strings.EqualFold(command, "search") {
		return checkSearchCommandArgumentSize(args)
	} else {
		return "", ErrWrongCommand
	}
}

func printNoArgument() {
	PrintLineMessage("No command found. Please check your command!")
}

func printWrongCommandUsed() {
	PrintLineMessage("This command is unknown to the program. Please check your command!")
}

func printWrongSearchCommand() {
	PrintLineMessage("Search command needs a following argument. Please check your command!")
}

func printOptions(projectName string) {
	fmt.Printf("%s Application Usage:\n", projectName)
	PrintLineMessage("\tgo run main.go <command> [arguments]")
	PrintLineMessage("The commands are:")
	PrintLineMessage("\tlist")
	PrintLineMessage("\tsearch [argument]")
}
