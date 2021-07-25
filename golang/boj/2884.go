package main

import "fmt"

func BojNo2884() {
	var h, m int
	fmt.Scanf("%d %d", &h, &m)

	m -= 45

	if m < 0 {
		m += 60
		h--
	}
	if h < 0 {
		h += 24
	}

	fmt.Printf("%d %d", h, m)
}
