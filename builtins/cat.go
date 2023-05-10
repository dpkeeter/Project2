//currently expects either syntax as "cat [filename]" or "cat"

package builtins

import (
	"bufio"
	"fmt"
	"os"
)

func Cat(userArgs ...string) {
	args := userArgs

	//file input case.
	if len(args) > 0 {
		//data, err := os.Open(file)
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Print("$ Unable to find existing file with filename: ")
			fmt.Println(args[0])
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		var linenum int = 1

		for scanner.Scan() {
			fmt.Print(linenum)
			fmt.Print(". ")
			fmt.Println(scanner.Text())
			linenum++
		}
		if err := scanner.Err(); err != nil {
			fmt.Print("$ File with filename '")
			fmt.Print(args[0])
			fmt.Println("' found, but unable to read file contents.")
			return
		}

	} else {
		//otherwise, echoes back user input (only does this once).
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		fmt.Println(input)

		return
	}
	return
}
