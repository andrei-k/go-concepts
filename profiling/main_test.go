package main

import "testing"

func TestGetVal(t *testing.T) {
	// Running it 1,000 times.
	for i := 0; i < 1000; i++ {
		if Fib2(30) != 832040 {
			t.Error("Incorrect!")
		}
	}
}
