package main

import (
	"fmt"
)

func main() {
	fmt.Println("Penulisan deklarasi variable digolang memiliki 3 cara, yang dikategorikan menjadi 2 tipe penulisan\n1. manifest typing")
	fmt.Println("Contoh: ")
	fmt.Println("var namavariable string = value\n\"var\" digunakan untuk mendeklarasikan bahwa dia adalah variable, dilanjut dengan nama variable,terakhir itu tipe data\ndituliskan secara ekspilisit")
	fmt.Println("var namavariable = value\nseperti sebelumnya\n\"var\"digunakan untuk mendeklarasikan variable hanya saja bedanya yang kedua adalah tidak menuliskan tipe datanya.\njadi golang langsung menentukan tipe data secara otomatis terhadapnya tergantung value dari variable tersebut")
	fmt.Println("2 type inherence\nnama := value")
	var fullname string
	fmt.Println("Default string jika belum didefinisikan maka dia berisi string kosong = \"\"", fullname)
	fullname = "susumurni"
	fmt.Println("Setelah dirubah value sebuah variable : ", fullname)
	var minuint = len(fullname)
	fmt.Println("tipe inherence lebih singkat lagi,\"namavariable := value", minuint, "\"")
	akuganteng := true
	fmt.Println("Apakah aku ganteng ?", akuganteng)
	const nama = "muhamad Rizki arif Fadillah"
	fmt.Println("Nama : ", nama, "penggunaan const adalah ketika constanta didekalarsikan isinya tidak bisa dirubah kembali")
}
