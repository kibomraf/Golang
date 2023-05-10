package main

import "fmt"

func main() {
	const introduce string = "Const adalah sebuah deklarasi yang digunakan, saat ketika didekalarsikan\nmaka value dari variable tersebut tidak bisa diganti, apabila mau diganti maka akan menyebabkan error"
	fmt.Println("Penjelasan : ", introduce)
	fmt.Println("Deklarasi mutiplec constant")
	const (
		firstname = "Muhamad Rizki"
		lastname  = "Arif Fadillah"
	)
	fmt.Println(firstname, lastname)
	const alamat, negara string = "Purwakarta", "Indonesia"
	fmt.Println(alamat, negara)
}
