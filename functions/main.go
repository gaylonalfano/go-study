package main

import (
	"fmt"
)

// Functions - Closures
func calculator(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func closures() {
	base := 10

	result := calculator(20, 3, func(a int, b int) int {
		return a * b * base
	})

	fmt.Println(result) // 600
}

// Higher Order Functions
func adder(start int) func(int) int {
	return func(a int) int {
		return a + start
	}
}

func higherOrderFunctions() {
	twoAdder := adder(2) // func(int) int
	threeAdder := adder(3)
	fmt.Println(twoAdder(10), threeAdder(10))
}

// Call-by-value language (a COPY is sent to the function)
// REF: https://youtu.be/3zcKjKySMKE?t=1211
func double(a int) {
	a = a * 2
}

func callByValue() {
	a := 10
	// Copy of 'a' is sent to double()
	double(a)
	fmt.Println(a) // 10
}

// NOTE: Use pointers when you don't want a copy and you want to modify parameters
// & - Use & to get the reference to a pointer to a variable
// * - Use * to indicate a type is a pointer type: a *int (pointer to int)
// * - Use * with a variable to dereference (i.e, you want to refer to the variable that's pointed to by the pointer)
func doubleWithPointers(a *int) {
	*a = *a * 2
}

func pointersToModifyValue() {
	a := 10
	// NOTE: A pointer parameter (a *int). A pointer argument (&a).
	// Pointer parameter is still a Call-by-value. The difference is the value that
	// you are copying is the memory location of the data, NOT the data itself!
	doubleWithPointers(&a) // &a to get a pointer that references 'a'
	fmt.Println(&a)        // 0x1400009e028
	fmt.Println(a)         // 20. The original value of 'a' has been modified
}

func main() {
	closures()
	higherOrderFunctions()
	callByValue()           // 10
	pointersToModifyValue() // 20
}
