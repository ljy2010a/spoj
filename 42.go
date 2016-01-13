package main

import "fmt"

func main() {
	for value := 0; ; {
		fmt.Scanln(&value)
		if value == 42 {
			break
		}
		fmt.Println(value)
	}
}
