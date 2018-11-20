// Take 2 strings s1 and s2 including only letters from ato z. Return a new sorted string, the longest possible, containing distinct letters,

// each taken only once - coming from s1 or s2.

// #Examples:
// a = "xyaabbbccccdefww" b = "xxxxyyyyabklmopq" longest(a, b) -> "abcdefklmopqwxy"
// a = "abcdefghijklmnopqrstuvwxyz" longest(a, a) -> "abcdefghijklmnopqrstuvwxyz"

package kata

import (
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	s := s1 + s2
	strs := strings.Split(s, "")
	result := sortAndDeduplicate(strs)
	return strings.Join(result, "")
}

func sortAndDeduplicate(strs []string) []string {
	sort.Strings(strs)
	j := 0
	for i := 1; i < len(strs); i++ {
		if strs[j] == strs[i] {
			continue
		}
		j++
		strs[i], strs[j] = strs[j], strs[i]
	}
	result := strs[:j+1]
	return result
}
