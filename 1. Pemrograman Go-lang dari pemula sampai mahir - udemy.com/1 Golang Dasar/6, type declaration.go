package main

import "fmt"

type huruf string
type angka int

func main() {
	var nama huruf = "Muhamad Rizki Arif Fadillah"
	var alamat huruf = "Purwakarta"
	var usia angka = 23
	fmt.Println("Hallo, Perkenalkan nama saya", nama, "saya berasal dari", alamat, "usia saya saat ini", usia)

}
