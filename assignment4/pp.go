package main

import (
	"fmt"
)

func main() {

	test := [4]string{"a", "b", "c", "d"}
	fmt.Println(test)
	updateThirdElement(&test)
	fmt.Println(test)

}

func updateThirdElement(x *[4]string) {

	(*x)[2] = "Texas"

}
