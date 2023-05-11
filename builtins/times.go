package builtins

import (
	"fmt"
	"syscall"
)

func Times() error {
	ut := new(syscall.Rusage)
	kt := new(syscall.Rusage)
	return fmt.Errorf("User Time in nano seconds %v  Kernel Time in nano seconds %v ", ut.UserTime.Nanoseconds(), kt.KernelTime.Nanoseconds())
}
