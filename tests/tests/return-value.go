package main

import "fmt"

func otherfunc() int {
	// in other func
	fmt.Println("in other func")
	return 100
}

func main() {
	// 100
	printf("%d\n", otherfunc())
}
