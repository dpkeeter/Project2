package builtins

import (
	"fmt"
	"io"
	"os"
)

func PrintWorkingDirectory(w io.Writer, args ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Printf(" \"%v\" is the working directory.\n", wd)
	return err
}