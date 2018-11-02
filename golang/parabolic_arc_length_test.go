package kata

import "fmt"

func Example_LenCurve() {
	fmt.Println(LenCurve(1))
	fmt.Println(LenCurve(10))
	fmt.Println(LenCurve(40))
	fmt.Println(LenCurve(200))
	// Output:
	// 1.414213562
	// 1.478197397
	// 1.478896272
	// 1.478940994
}
