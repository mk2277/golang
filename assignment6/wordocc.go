package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("enter the statement:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input_str := scanner.Text()
	newmap := map[string]int{}
	x := strings.Split(input_str, " ")

	keymap := map[string]int{}
	//counting words by using map++
	for i := range x {

		keymap[x[i]] = keymap[x[i]] + 1

	}
	fmt.Println(keymap)
	//counting words using string count method
	for i := range x {
		count := strings.Count(input_str, x[i])
		newmap[x[i]] = count
	}

	fmt.Println(newmap)

}
