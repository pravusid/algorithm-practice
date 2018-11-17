package kata

import (
	"fmt"
)

func Example_BlackOrWhiteKey() {
	fmt.Println(BlackOrWhiteKey(1))
	fmt.Println(BlackOrWhiteKey(5))
	fmt.Println(BlackOrWhiteKey(12))
	fmt.Println(BlackOrWhiteKey(42))
	fmt.Println(BlackOrWhiteKey(88))
	fmt.Println(BlackOrWhiteKey(89))
	fmt.Println(BlackOrWhiteKey(92))
	fmt.Println(BlackOrWhiteKey(100))
	fmt.Println(BlackOrWhiteKey(111))
	fmt.Println(BlackOrWhiteKey(200))
	fmt.Println(BlackOrWhiteKey(2017))
	// Output:
	// white
	// black
	// black
	// white
	// white
	// white
	// white
	// black
	// white
	// black
	// white
}
