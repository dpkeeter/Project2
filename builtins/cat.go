//currently expects syntax: cat [filename]

package builtins

import (
	"bufio"
	"fmt"
	"os"
)

func Cat(userArgs ...string) {

	args := userArgs

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
		fmt.Println("$ too few arguments, expected:  cat [filename]")
		return
	}
	return
}
