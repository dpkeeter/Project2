package builtins

import "fmt"

func Shift() {
	var i0 int
	var i1 int
	fmt.Printf("For x << n Enter x: ")
	fmt.Scanf("%d", &i0)
	fmt.Printf("For x << n Enter n: ")
	fmt.Scanf("%d", &i1)
	fmt.Printf("\n%d << %d = %d", i0, i1, i0<<i1)
	fmt.Printf("\n%d >> %d = %d \n", i0, i1, i0>>i1)

	return
}
