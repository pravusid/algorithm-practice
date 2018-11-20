package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	str := new(strings.Builder)
	for i := 1; i <= n; i++ {
		str.WriteString(strconv.Itoa(i) + "\n")
	}
	fmt.Print(str)
}
