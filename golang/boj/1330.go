package main

import "fmt"

func BojNo1330() {
	var i, j int
	fmt.Scanf("%d %d", &i, &j)
	switch {
	case i > j:
		fmt.Printf(">")
	case i < j:
		fmt.Printf("<")
	default:
		fmt.Printf("==")
	}
}
