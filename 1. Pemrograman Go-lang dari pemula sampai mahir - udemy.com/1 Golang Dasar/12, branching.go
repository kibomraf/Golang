package main

import "fmt"

func main() {
	var value uint
	fmt.Println("Please enter a value from one to ten")
	fmt.Printf("value : ")
	fmt.Scan(&value)
	if value == 10 {
		fmt.Println("You graduated with perfect grade.")
	} else if value > 7 && value <= 9 {
		fmt.Println("You graduated quite well")

	} else if value >= 5 && value <= 7 {
		fmt.Println("You graduate with a mediocre grade")
	} else {
		fmt.Println("You are not trying enough. You can do it, try again and improve your score in the next exam.")
	}
	switch value {
	case 1:
		fmt.Printf("your score is %v , you do not pass this test\n", value)
	case 2:
		fmt.Printf("your score is %v , you do not pass this test\n", value)
	case 3:
		fmt.Printf("your score is %v , you do not pass this test\n", value)
	case 4:
		fmt.Printf("your score is %v , you do not pass this test\n", value)
	case 5:
		fmt.Printf("your score is %v , you pass this test, great! keep trying and practicing your skills so that your test score grow.\n", value)
	case 6:
		fmt.Printf("your score is %v , you pass this test, great! please upgrade your skills and get better grade\n", value)
	case 7:
		fmt.Printf("your score is %v , you pass this test, great! please upgrade your skills and get better grade\n", value)
	case 8:
		fmt.Printf("your score is %v , you pass this test, a little more for your score to be perfect, keep practicing your skills.\n", value)
	case 9:
		fmt.Printf("your score is %v , you pass this test. a little more for your score to be perfect, keep practicing your skills.\n", value)
	case 10:
		fmt.Printf("your score is %v , you pass this test. your score is perfect, keep up your skill by practicing it every day.\n", value)
	}
}
