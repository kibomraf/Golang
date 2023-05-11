package main

import "fmt"

func main() {
	students_name := [4]string{"Muhamad", "Rizki", "Arif", "Fadillah"}
	slice := students_name[0:2]
	fmt.Println("Array =", students_name)
	fmt.Println("Slice = ", slice)
	fmt.Println("Slice termasuk tipe data by refence jad ketika value slice diubah maka sumbernya pun valuenya ikut berubah")
	slice[1] = "Kibo"
	fmt.Println("Array =", students_name)
	fmt.Println("Slice = ", slice)
	//append
	slice = append(slice, "Ega")
	fmt.Println("Slice = ", slice)
	//len
	fmt.Println("Panjang Slice =", len(slice))
	//make
	make_slice := make([]string, 2, 5)
	make_slice[0] = "Rizki"
	make_slice[1] = "Bokir"
	//make_slice[2] = "Kibo" Error
	fmt.Println("Ketika kita menentukan panjang slice, kita tidak bisa mendeklarasikan lebih dari panjang indeks yang ditentukan tetapi kita dapat menambahkan data terhadap slice kita")
	make_slice = append(make_slice, "Kibo")
	fmt.Println("Panjang make slice:", len(make_slice))
	fmt.Println("Kapasitas make slice:", cap(make_slice))
	fmt.Println("make slice", make_slice)
	//copy
	fromSlice := make_slice[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println("To Slice :", toSlice)
	fmt.Println("Perbedaan antara slice dan array secara pemahaman saya array harus memiliki batas maksimal data tapi kalau slice tidak menentukan berapa banyak data,\njadi kita bisa memasukan berapapun data yang kita inginkan")
}
