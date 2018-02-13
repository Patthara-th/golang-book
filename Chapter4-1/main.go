package main

import "fmt"

func main() {
	// Long Declaration
	var x string = "Hello, World"
	fmt.Println(x)

	var y string
	y = "Hello, World"
	fmt.Println(y)

	// Short Declaration
	// Type Inference
	z := "Hello, World"
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)

	fmt.Println("")
	fmt.Println("++++++++++++++++++++++")
	fmt.Println("")

	zz := 1
	fmt.Println(zz)
	fmt.Printf("Type: %T\n", zz)

	zzz := 2.0
	fmt.Println(zzz)
	fmt.Printf("Type: %T\n", zzz)

	zzzz := true
	fmt.Println(zzzz)
	fmt.Printf("Type: %T\n", zzzz)

	const xx = "Hello, World"

	fmt.Println("")
	fmt.Println("++++++++++++++++++++++")
	fmt.Println("")

	var (
		a = 5
		b = 10
		c = 15
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	v1, v2, v3, v4, v5 := "first", "sec", 0, 1.0, true

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)

	fmt.Println("")
	fmt.Println("++++++++++++++++++++++")
	fmt.Println("")

	foo, bar := 1, 2
	fmt.Println(foo)
	fmt.Println(bar)
	foo, bar = bar, foo
	fmt.Println(foo)
	fmt.Println(bar)

}
