package main

import "fmt"

func main() {
	for a := 0; a < 10; a++ {
		fmt.Println("increment repetition:", a)
	}
	for b := 10; b > 0; b-- {
		fmt.Println("decrement repetition:", b)
	}
	fmt.Println("Create a box using a looping program")
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			fmt.Print(" * ")
		}
		fmt.Println()
	}
	fmt.Println("Create an isoceles triangle from the left using a looping program")
	for z := 1; z < 7; z++ {
		for x := 0; x < z; x++ {
			fmt.Print("#")
		}
		fmt.Println("")
	}
	fmt.Println("Create an isoceles triangle from the right using a looping program")
	for c := 1; c <= 10; c++ {
		for d := 10; d > c; d-- {
			fmt.Print("_")
		}
		for e := 0; e < c; e++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
