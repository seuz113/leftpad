package main

import (
	"fmt"

	"github.com/seuz113/leftpad"
)

func main() {
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.Format("goodbye", 15))
	fmt.Println(leftpad.Format("Como esta?", 15))
	fmt.Println(leftpad.Format("Internacionalization", 15))
	fmt.Println(leftpad.FormatRune("hello", 15, '😃'))
	fmt.Println(leftpad.FormatRune("goodbye", 15, '😃'))
	fmt.Println(leftpad.FormatRune("Como esta?", 15, '😃'))
}
