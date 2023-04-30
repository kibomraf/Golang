package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Tipe data Number")
	fmt.Println("tanpa decimal")
	fmt.Println("int dan uint memiliki 4 versi dengan jumlah yang berbeda")
	fmt.Println("int8 :", math.MinInt8, "-", math.MaxInt8)
	fmt.Println("int16 :", math.MinInt16, "-", math.MaxInt16)
	fmt.Println("int32 :", math.MinInt32, "-", math.MaxInt32)
	fmt.Println("int64 :", math.MinInt64, "-", math.MaxInt64)
	fmt.Println("Alias : int", math.MinInt, "-", math.MaxInt, "alias dari int ini bergantung pada bit os kita kalau bit os kita 32\nmaka alias int sama dengan jumlah int32\njika bit os kita 64 maka sama dengan int64")
	fmt.Println("uint8 : 0 - ", math.MaxUint8)
	fmt.Println("uint16 : 0 - ", math.MaxUint16)
	fmt.Println("uint32 : 0 - ", math.MaxUint32)
	fmt.Println("uint64 : 0 - ", uint(math.MaxUint64))
	fmt.Println("Alias uint : 0 -", uint(math.MaxUint), "sama seperti int jumlahnya tergantung bit os kita")
	// "\" ini disebut escape squence untuk menandai bahwa kutip dua bukan sebagai pembatas string
	fmt.Println("Alias : \"byte\" sama seperti uint8 :")
	fmt.Println("Alias : \"rune\" sama seperti int32 :")
	fmt.Println("dengan decimal")
	fmt.Println("float , hanya memiliki 2 versi float32 dan float64")
	fmt.Println("float32 : ", "standa CPU minimal (", math.SmallestNonzeroFloat32, ")representasi standar dari IEEE 754(", -math.MaxFloat32, ") maksimal : ", math.MaxFloat32)
	fmt.Println("float64 : ", "standa CPU minimal (", math.SmallestNonzeroFloat64, ")representasi standar dari IEEE 754(", -math.MaxFloat64, ") maksimal : ", math.MaxFloat64)
	fmt.Println("tipe data string \nitu yang biasanya ditandai oleh kutip 2 \"\"\nbaik termasuk angka,huruf dan simbol seperti contoh sebelumya dan sekarang yang menjadi output diterminal")
	fmt.Println("menghitung total jumlah string \"kibomraf\"", len("kibomraf"))
	fmt.Println("menampilkan jumah bit yang digunakan untuk mencetak satu huruf", "kibomraf"[1])
	fmt.Println("tipe data boolean hanya memiliki 2 hasil yaitu ", true, "dan", false)

}
