package kata

import (
	"fmt"
)

func Example_TwoToOne() {
	fmt.Println(TwoToOne("aretheyhere", "yestheyarehere"))
	fmt.Println(TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding"))
	// Output:
	// aehrsty
	// abcdefghilnoprstu
}
