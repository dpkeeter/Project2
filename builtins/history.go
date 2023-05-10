package builtins

import (
	"fmt"
)

func History(currenthistory []string, name string, args ...string) error {

	if len(args) > 1 {
		return fmt.Errorf("%w: INCORRECT NUMBER OF ARGS", ErrInvalidArgCount)
	}
	if len(args) == 0 {
		fmt.Println("Print History Here")
		return nil
	}
	var arg = args[0]

	switch arg {
	default:
		fmt.Println("Not a valid command only (-h) (-c) (-r) and (n) are allowed")
	case "-h":
		fmt.Println("Print History without even numbers here")
	case "-c":
		fmt.Println("Clear History Here")
	case "-r":
		fmt.Println("Print History in reverse Here")
	}

	return nil
}
