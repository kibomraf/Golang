package main

import "fmt"

func main() {
	var angka1 uint = 10
	var angka2 uint = 100
	fmt.Printf("%v + %v = %v\n", angka2, angka1, angka2+angka1)
	fmt.Printf("%v + %v = %v\n", angka2, angka1, angka2-angka1)
	fmt.Printf("%v + %v = %v\n", angka2, angka1, angka2*angka1)
	fmt.Printf("%v + %v = %v\n", angka2, angka1, angka2/angka1)
	angka3 := 99
	angka4 := 30
	fmt.Println(angka3, "+", angka4, "=", angka3+angka4)
	fmt.Println(angka3, "-", angka4, "=", angka3-angka4)
	fmt.Println(angka3, "*", angka4, "=", angka3*angka4)
	fmt.Println(angka3, "/", angka4, "=", float64(angka3)/float64(angka4))
	fmt.Println(angka3, "%", angka4, "=", angka3%angka4)
	var koma1 = 10.6
	var koma2 = 30.2
	fmt.Print(koma1, "+", koma2, "=", koma1+koma2)

}
