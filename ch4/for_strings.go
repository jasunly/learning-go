package main

import "fmt"

func main() {
	words := []string{"hi", "salutations", "hello"}

	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "this is a short word!")
		case wordLen > 10:
			fmt.Println(word, "this is a long word!")
		default:
			fmt.Println(word, "this is a right size!")
		}
	}
}
