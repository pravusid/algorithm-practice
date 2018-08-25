package kata

import (
	"strings"
)

func Arrange(s string) string {
	split := strings.Split(s, " ")
	for i := 0; i < len(split); i++ {
		if i < len(split)-1 && !isCorrectOrder(split, i) {
			split[i], split[i+1] = split[i+1], split[i]
		}
		if i%2 != 0 {
			split[i] = strings.ToUpper(split[i])
		} else {
			split[i] = strings.ToLower(split[i])
		}
	}
	return strings.Join(split, " ")
}

func isCorrectOrder(split []string, i int) bool {
	var x, y int
	if i%2 == 0 {
		x, y = i, i+1
	} else {
		x, y = i+1, i
	}
	return len(split[x]) <= len(split[y])
}
