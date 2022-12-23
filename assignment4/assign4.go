package main

import "fmt"

func main() {

	//1
	fmt.Println("1st question")
	var a [5]int

	fmt.Println("enter 5 values into array")

	for i := 0; i < 5; i++ {
		fmt.Scanln(&a[i])
	}
	fmt.Println("5 values of array are :")

	fmt.Println(a)

	fmt.Printf("type and space of array is %T", a)
	fmt.Println("")

	//2
	fmt.Println("2nd question")
	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("type and space of slice is %T", b)
	fmt.Println("")
	//3
	fmt.Println("3rd question")
	fmt.Println(b[0:5])
	fmt.Println(b[5:])
	fmt.Println(b[2:7])
	fmt.Println(b[1:6])

	//4
	fmt.Println("4th question")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

}
