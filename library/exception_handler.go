package library

import (
	"errors"
	"strings"
)

var ErrNoArgument = errors.New("No user argument found!")
var ErrWrongCommand = errors.New("Unknown command type was used!")
var ErrWrongSearchCommand = errors.New("One argument was given for search command!")

func CheckUserArguments(args []string) (string, error) {
	if err := checkSize(args); err != nil {
		printNoArgument()
		printOptions()
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
			printOptions()
		}
		return command, err
	}
}

func checkSize(args []string) error {
	if cap(args) <= 1 {
		return ErrNoArgument
	}
	return nil
}

func checkSearchCommandArgumentSize(args []string) (string, error) {
	if cap(args) <= 2 {
		return SearchCommand, ErrWrongSearchCommand
	}
	return SearchCommand, nil
}

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
	PrintLineMessage("No command found. Please check your request. Please check your command!")
}

func printWrongCommandUsed() {
	PrintLineMessage("This command is unknown to the program. Please check your command!")
}

func printWrongSearchCommand() {
	PrintLineMessage("Search command needs a following argument. Please check your command!")
}

func printOptions() {
	PrintLineMessage("Usage:")
	PrintLineMessage("\tgo run main.go <command> [arguments]")
	PrintLineMessage("The commands are:")
	PrintLineMessage("\tlist")
	PrintLineMessage("\tsearch [argument]")
}
