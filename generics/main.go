package main

import "fmt"

// Define new type argument called "age"
// Specify the list of types that age can be
func add[num int64 | float32](n num) {
	// In order to convert a type, the function must define type constraints like above and not use "any"
	val := float32(n) + 1
	fmt.Println(val)
}

// Ability to accept any type with "any" keyword
func display[age any](myAge age) {
	// Note that Println takes a variadic interface and can handle any type
	fmt.Println(myAge)
}

// Better way is to define a "Number" interface and pass in all the allowed types
type Number interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

// Can then use this type constraint in the function
func add2[num Number](n num) {
	val := float32(n) + 1
	fmt.Println(val)
}

// A BubbleSort computation that accepts any type of any slice
func bubbleSort[N Number](input []N) []N {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func main() {
	var num1 int64 = 28
	var num2 float32 = 28.6
	var str string = "Generics!"

	add(num1)
	add(num2)

	display(num1)
	display(num2)
	display(str)

	add2(num1)
	add2(num2)

	listInt := []int32{3, 8, 2, 1, 5, 6, 4, 7}
	listFloat := []float32{3.2, 8.7, 2.2, 1.1, 5.1, 6.1, 4.1, 7.1}

	sortedInt := bubbleSort(listInt)
	fmt.Println(sortedInt)

	sortedFloat := bubbleSort(listFloat)
	fmt.Println(sortedFloat)
}
