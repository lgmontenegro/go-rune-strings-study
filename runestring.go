package main

import (
	"fmt"
)

func main() {
	l := "😂"

	fmt.Println(l, len(l), []byte(l), []rune(l))

	fmt.Printf("%U\n", []rune("ç"))
}
