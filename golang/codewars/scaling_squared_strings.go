package kata

import "strings"

func Scale(s string, k, n int) string {
	if s == "" {
		return ""
	}
	split := strings.Split(s, "\n")
	var result []string
	for _, line := range split {
		var newLine string
		for _, str := range line {
			newLine += strings.Repeat(string(str), k)
		}
		for i := 0; i < n; i++ {
			result = append(result, newLine)
		}
	}
	return strings.Join(result, "\n")
}
