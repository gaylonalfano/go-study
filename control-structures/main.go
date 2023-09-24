package main

import "fmt"

func basicIfStatements() {
	a := 5
	var b string
	if a < 10 {
		b = "less than 10"
	} else {
		b = "greater or equal to 10"
	}
	fmt.Println(b)

	// Neat declare a var before the if condition and then use it
	if c := 10; c == 10 {
		fmt.Println("foo", c)
	} else {
		fmt.Println("bar", c)
	}
}

func main() {
	basicIfStatements()

	s := "Hello, World!"
	for k, v := range s {
		fmt.Println(k, "\t", v, "\t", string(v))
		// NOTE:
		// k = The position in bytes in the string
		// v = The Unicode value at the position as a rune (char, 32bit number)
		// Runes are 32-bit numbers, so you turn into a Unicode character,
		// you cast it to a string(v) before printing it.
		// 0        72      H
		// 1        101     e
		// 2        108     l
		// 3        108     l
		// 4        111     o
		// 5        44      ,
		// 6        32
		// 7        87      W
		// 8        111     o
		// 9        114     r
		// 10       108     l
		// 11       100     d
		// 12       33      !
	}

	// Switch
	n := 10
	switch n {
	case 1, 2, 3:
		println("n <= 3")
	case 4, 5, 6:
		fmt.Println("4 <= n <= 6")
	default:
		fmt.Println("n is big")
	}

	// Switch + fallthrough
	i := 1
	switch i {
	case 1, 2, 3:
		println("i <= 3")
		fallthrough // both i<=3 and 4<=i<=6 will print
	case 4, 5, 6:
		fmt.Println("4 <= i <= 6")
	default:
		fmt.Println("i is big")
	}

	// x is only in scope for the switches' cases (reminds me of lifetimes)
	j := 1
	switch x := j / 2; x {
	case 1, 2, 3:
		println("x <= 3")
	case 4, 5, 6:
		fmt.Println("x <= 6")
	default:
		fmt.Println("x is big")
	}

	// Blank switch where each case is like a boolean condition
	k := 10
	switch {
	case k%2 == 1:
		println("odd")
	case k%5 == 0:
		fmt.Println("multiple of 5") // Gets printed
	case k%2 == 0:
		fmt.Println("even")
	default:
		fmt.Println("a weird number")
	}

}
