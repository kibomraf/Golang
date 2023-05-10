package main

import "fmt"

func main() {
	var angka1 uint = 10
	angka2 := 20
	var comparison bool = angka1 > uint(angka2)
	fmt.Println("Apakah angka 1 > angka 2 =", comparison)
	comparison = angka1 < uint(angka2)
	fmt.Println("Apakah Angka1 < angka2 =", comparison)
	angka1 = 20
	angka2 = 20
	comparison = angka1 == uint(angka2)
	fmt.Println("apakah angka1 sama dengan angka2 =", comparison)
	fmt.Println("Ada juga >= <= !=")
	comparison = angka1 != uint(angka2)
	fmt.Println("apakah angka1 sama dengan angka2 =", comparison)

}
