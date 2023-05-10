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

	/*	toRemove := make([]string, 0)
		for i := 0; i < len(args); i++ {
			if args[i] == "-u" {
				if len(args) < i+2 {
					return fmt.Errorf("%w: -u requires an argument", ErrInvalidArgCount)
				}
				toRemove = append(toRemove, args[i+1])
				i++
			}
		}

		toShow := make([]string, 0)
		for _, env := range os.Environ() {
			show := true
			for _, v := range toRemove {
				if strings.HasPrefix(env, v+"=") {
					show = false
					break
				}
			}
			if show {
				toShow = append(toShow, env)
			}
		}

		_, err := fmt.Fprintln(w, strings.Join(toShow, "\n"))

		return err*/
	return nil
}
