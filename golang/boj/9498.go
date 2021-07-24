package main

import "fmt"

func BojNo9498() {
	var score int
	fmt.Scanf("%d", &score)
	switch {
	case score >= 90:
		fmt.Printf("A")
	case score >= 80:
		fmt.Printf("B")
	case score >= 70:
		fmt.Printf("C")
	case score >= 60:
		fmt.Printf("D")
	default:
		fmt.Printf("F")
	}
}
