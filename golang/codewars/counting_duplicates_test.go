package kata

import (
	"fmt"
)

func Example_duplicate_count() {
	fmt.Println(duplicate_count("abcde"))
	fmt.Println(duplicate_count("abcdea"))
	fmt.Println(duplicate_count("abcdeaB11"))
	fmt.Println(duplicate_count("indivisibility"))
	// Output:
	// 0
	// 1
	// 3
	// 1
}
