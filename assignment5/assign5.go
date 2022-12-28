package main

// "strings"
import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//1
	fmt.Println("First problem:")
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	//2
	fmt.Println("Second problem:")
	i := 0
	for i < 50 {
		if i%2 != 0 {
			fmt.Println(i)
		}
		i++
	}

	//3
	fmt.Println("Third problem:")
	i = 0
	for {
		if i < 50 {
			if i%2 == 0 {
				fmt.Println(i)
			}
			i++
		}
		if i == 50 {
			break
		}

	}

	//4
	fmt.Println("Fourth problem:")
	for i := 50; i < 106; i++ {
		fmt.Println(i, i/6)
	}

	//5
	fmt.Println("Fifth problem:")
	y := "Golang tutorial"
	fmt.Println("enter string")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	x := scanner.Text()

	if x == y {
		fmt.Println("Welcome")
	} else {
		fmt.Println("end")
	}

	//6
	fmt.Println("sixth problem:")
	for i := 1; i <= 80; i++ {
		if i%2 == 0 && i%4 == 0 {
			fmt.Println("Golang tutorial")
		} else if i%2 == 0 {
			fmt.Println("Golang")
		} else if i%4 == 0 {
			fmt.Println("tutorial")
		} else {
			fmt.Println(i)
		}
	}

}
