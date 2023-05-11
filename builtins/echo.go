package builtins

import (
	"fmt"
	"strings"
)

func Echo(userArgs ...string) {
	if len(userArgs) > 0 {
		//echoes back user input args seperated by whitespace and ending in newline.
		fmt.Println(strings.Join(userArgs, " "))
	} else {
		fmt.Println("$ Unable to echo given input, expected: echo [echo input] ")
	}
	return
}
