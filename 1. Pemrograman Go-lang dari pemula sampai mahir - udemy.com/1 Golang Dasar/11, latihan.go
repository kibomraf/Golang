package main

import (
	"fmt"
	"math"
)

func main() {
	//buatlah program dimana meminta pengguna memasukan 2 angka yang hasilnya termasuk penjumlaha, pengurangan,perkalian,pembagian, modulus,perpangkatan.
	var a int
	var b int
	fmt.Print("Masukan angka pertama : ")
	fmt.Scan(&a)
	fmt.Print("Masukan angka kedua : ")
	fmt.Scan(&b)
	fmt.Printf("Hasil dari %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Hasil dari %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Hasil dari %d x %d = %d\n", a, b, a*b)
	fmt.Printf("Hasil dari %d / %d = %d\n", a, b, a/b)
	fmt.Printf("Hasil dari %d %% %d = %d\n", a, b, a%b)
	fmt.Printf("Hasil dari %d ^ %d = %d\n", a, b, int(math.Pow(float64(a), float64(b))))
	//buatlah sebuah program comparison operation yang hasilnya menyatakan bahwa apakah angka pertama lebih besar sama dengan angka kedua atau tidak
	comparison := a >= b
	fmt.Println("Apakah angka", a, "lebih besar sama dengan angka", b, comparison)
}
