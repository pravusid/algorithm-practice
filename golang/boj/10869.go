package main

import "fmt"

func BojNo10869() {
	var i, j, k int
	fmt.Scanf("%d %d %d", &i, &j, &k)

	fmt.Printf("%d\n%d\n%d\n%d", (i+j)%k, ((i%k)+(j%k))%k, (i*j)%k, ((i%k)*(j%k))%k)
}
