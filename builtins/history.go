package builtins

import (
	"fmt"
)

func History(currenthistory []string, args ...string) (error, []string) {

	if len(args) > 1 {
		return fmt.Errorf("%w: INCORRECT NUMBER OF ARGS", ErrInvalidArgCount), currenthistory
	}
	if len(args) == 0 {
		for i := 0; i < len(currenthistory); i++ {
			fmt.Println("Event [", i, "]", currenthistory[i])
		}
		return nil, currenthistory
	}
	var arg = args[0]

	switch arg {
	default:
		fmt.Println("Not a valid command only (-h) (-c) (-r) and (n) are allowed")
	case "-h":
		for i := 0; i < len(currenthistory); i++ {
			fmt.Println(currenthistory[i])
		}
	case "-c":
		currenthistory = nil
		return nil, currenthistory
	case "-r":

		for i := len(currenthistory) - 1; i >= 0; i-- {
			fmt.Println("Event [", i, "]", currenthistory[i])
		}
		return nil, currenthistory
	}

	return nil, currenthistory
}
