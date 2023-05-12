package main

import "fmt"

func main() {
	//map adalah sebuah tipe data yang didalamnya harus menentukan key and valuenya harus dengan type apa contoh
	//map[type_of_key]type_of_value
	var ini_map = map[string]string{
		"Name":    "Muhamad Rizki Arif Fadillah",
		"Address": "Purwakarta, Jawabarat",
		"Hobby":   "Read a book about self improvment, and learn programing language",
	}
	fmt.Println("ini map :", ini_map)
	fmt.Println(ini_map["Name"])
	delete(ini_map, "Address")
	fmt.Println("new ", ini_map)
}
