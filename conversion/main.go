package main

import "fmt"

func main() {
	z := 4

	var m float32 = 43.742

	z = int(m)

	fmt.Printf("%v of type %T\n", z, z)
}
