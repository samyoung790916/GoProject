package main

import "fmt"

// func Divide(a, b int) (result int, success bool) {
// 	if b == 0 {
// 		result = 0
// 		success = false
// 		return
// 	}
// 	result = a / b
// 	success = true
// 	return
// }

// func main() {
// 	c, success := Divide(9, 3)
// 	fmt.Println(c, success)
// 	d, success := Divide(9, 0)
// 	fmt.Println(d, success)
// }

// func printNo(n int) {
// 	if n == 0 {
// 		return
// 	}
// 	fmt.Println(n)
// 	printNo(n - 1)
// 	fmt.Println("After", n)
// }

// func main() {
// 	printNo(100)
// }

func getMyAge() int {
	return 22
}

func main() {

	switch age := getMyAge(); true {
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("TeenAger")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age)
	}
}
