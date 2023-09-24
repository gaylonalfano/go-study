package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	A int
	B string
}

func main() {
	s := `
{
"a": 10,
"b": "Hello"
}`

	var f Foo
	err := json.Unmarshal([]byte(s), &f)
	fmt.Println(f, err) // {10 Hello} <nil>
	s2, err := json.Marshal(f)
	fmt.Println(string(s2), err) // {"A":10,"B":"Hello"} <nil>
}
