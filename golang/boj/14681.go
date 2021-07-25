package main

import "fmt"

func BojNo14681() {
	var a, b, res int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	if a > 0 {
		if b > 0 {
			res = 1
		} else {
			res = 4
		}
	} else {
		if b > 0 {
			res = 2
		} else {
			res = 3
		}
	}

	fmt.Print(res)
}
