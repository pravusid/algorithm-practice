package kata

import (
	"fmt"
)

// h = 3, bounce = 0.66, window = 1.5, result is 3
// h = 3, bounce = 1, window = 1.5, result is -1 (Condition 2) not fulfilled).
func Example_BouncingBall() {
	fmt.Println(BouncingBall(3, 0.66, 1.5))
	fmt.Println(BouncingBall(3, 1, 1.5))
	// Output:
	// 3
	// -1
}
