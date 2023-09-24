package main

import "fmt"

// Compound Data Types (slices, structs)
// Slices are sequential list of vals that can grow
func slices() {
	var a []int
	for i := 0; i < 5; i++ {
		a = append(a, i*2)
	}

	fmt.Println("slice:", a)
	fmt.Println("slice length:", len(a))
	fmt.Println("value at position 2:", a[2])
	fmt.Println("slice from position 1 to 3", a[1:4])
	fmt.Println("all the indices and values:")
	for k, v := range a {
		fmt.Println(k, v)
	}
}

// Maps
func maps() {
	a := map[string]int{}
	a["bobby"] = 10
	a["pat"] = 20
	a["alex"] = 17
	fmt.Println("map:", a) // map: map[alex:17 bobby:10 pat:20]
	fmt.Println("map length:", len(a))
	fmt.Println("all the keys and values:")
	for k, v := range a {
		fmt.Println(k, v)
	}
	fmt.Println("value for pat:", a["pat"])
	delete(a, "pat")
	// NOTE: If a key doesn't exist, you get the default value of the type (int => 0)
	fmt.Println("value for pat after delete:", a["pat"]) // 0 (default of int type)
	// NOTE: Comma-Ok-Idiom: You put TWO vars on the LHS when reading from a map.
	// var1 => value of the key, var2 => boolean based on if key exists in map
	v, ok := a["pat"]
	fmt.Println("value and check for pat:", v, ok) // 0, false
}

// Structs (custom VALUE types, just like any core type)
// NOTE:Again, Go is pass-by-copy! If you pass a struct as an arg
// to a function, the function gets a copy of the struct.
// NOTE: If you want to modify a struct field, you pass a POINTER to the struct (&myStruct)
func structs() {
	type Foo struct {
		A int
		B string
	}

	func() {
		f := Foo{A: 10, B: "Hello"}
		fmt.Println(f.A, f.B)
		f.A = 20
		fmt.Println(f)
	}()
}

func structsAreValueTypes() {
	type Foo struct {
		A int
		B string
	}

	// Q: Can  you declare named functions inside other functions?
	// Can't seem to have another named function inside existing named func
	// A: Yes! It's a Closure! You have to assign to a variable.
	valueFoo := func(f Foo) {
		f.A = f.A * 2
	}

	// *Foo is a pointer TYPE
	pointerFoo := func(f *Foo) {
		f.A = f.A * 2
	}

	func() {
		f := Foo{A: 10, B: "pat"}
		valueFoo(f)
		fmt.Println("after valueFoo:", f) // after valueFoo: {10 pat}
		pointerFoo(&f)
		fmt.Println("after pointerFoo:", f) // after pointerFoo: {20 pat}
	}()
}

// Methods on structs
type Foo struct {
	A int
	B string
}

// Delegation (Go doesn't have inheritance)
// REF: https://youtu.be/3zcKjKySMKE?t=1492
type Bar struct {
	// Embedding Foo inside Bar
	// This lets you access all the fields of Foo, inside an
	// instance of Bar
	Foo
	C string
}

func (f Foo) LessThan(other int) bool {
	return f.A < other
}

func structMethods() {
	func() {
		f := Foo{A: 10, B: "Hello"}
		fmt.Println(f.LessThan(20))
	}()
}

func delegation() {
	ReadFoo := func(f Foo) {
		fmt.Println(f.B)
	}

	func() {
		f := Foo{A: 10, B: "Hello"}
		b := Bar{Foo: f, C: "Good Day"}
		// delegation!
		fmt.Println("A and B are in Foo:", b.A, b.B)
		fmt.Println("So is LessThan():", b.LessThan(20))
		fmt.Println("C is in Bar:", b.C)
		// ReadFoo(b)     // this is a compile-time error. Cannot use b as Foo value in arg to ReadFoo
		ReadFoo(b.Foo) // we pass in the Foo inside the Bar
	}()
}

// Interface - Defines one or more methods that a Type must implement
// in order to meet the interface. Reminds me of Rust Traits in a way.
// NOTE: In Go, interfaces are IMPLICIT; a type automatically meets
// an interfaces any time it implements ALL of an interfaces' methods.
// This really changes how you work with abstract types. It's the code
// that uses the implementation that defines the interface it needs.
// "Accept interfaces, return structs." Or, "Accept interfaces and
// return concrete types."
// REF: https://tutorialedge.net/golang/accept-interfaces-return-structs/
func interfaces() {
	type Lesser interface {
		LessThan(int) bool
	}

	printLesser := func(l Lesser, i int) {
		println(l.LessThan(i))
	}

	func() {
		f := Foo{A: 10, B: "Hello"}
		b := Bar{Foo: f, C: "Good Day"}
		printLesser(f, 20) // no interface declaration, but works!
		printLesser(b, 20) // this works too, bc Bar imbeds Foo (delegation), which meets interface!
	}()
}

func main() {
	slices()
	maps()
	structs()
	structsAreValueTypes()
	structMethods()
	delegation()
	interfaces()
}
