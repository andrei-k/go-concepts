package main

import "fmt"

func Fib2(n int) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fib2(n-1) + Fib2(n-2)
	}
}

func main() {
	fmt.Println(Fib2(30)) // Answer is: 832040
}
