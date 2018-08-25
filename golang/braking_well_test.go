package kata

import (
	"fmt"
)

func Example_Dist() {
	fmt.Println(Dist(144, 0.3))
	fmt.Println(Dist(92, 0.5))
	// Output:
	// 311.83146449201496
	// 92.12909477605366
}

func Example_Speed() {
	fmt.Println(Speed(159, 0.8))
	fmt.Println(Speed(164, 0.7))
	// Output:
	// 153.79671564846308
	// 147.91115701756493
}
