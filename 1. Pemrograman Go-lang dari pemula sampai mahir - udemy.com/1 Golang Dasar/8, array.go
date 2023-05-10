package main

import "fmt"

func main() {
	var list_of_students [4]string
	list_of_students[0] = "Muhamad"
	list_of_students[1] = "Rizki"
	list_of_students[2] = "Arif"
	list_of_students[3] = "Fadillah"
	fmt.Println("List of student = ", list_of_students)
	fmt.Println("1. ", list_of_students[0])
	fmt.Println("2. ", list_of_students[1])
	fmt.Println("3. ", list_of_students[2])
	fmt.Println("4. ", list_of_students[3])
	for _, value := range list_of_students {
		fmt.Println("Nama : ", value)
	}
}
