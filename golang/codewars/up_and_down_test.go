package kata

import "fmt"

func Example_Arrange() {
	fmt.Println(Arrange("who hit retaining The That a we taken")) // 3
	fmt.Println(Arrange("on I came up were so grandmothers"))     // 4
	fmt.Println(Arrange("way the my wall them him"))              // 1
	// Output:
	// who RETAINING hit THAT a THE we TAKEN
	// i CAME on WERE up GRANDMOTHERS so
	// way THE my WALL him THEM
}
