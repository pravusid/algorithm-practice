// Return the number (count) of vowels in the given string.
// We will consider a, e, i, o, and u as vowels for this Kata.
// The input string will only consist of lower case letters and/or spaces.

package kata

func GetCount(str string) (count int) {
	vowels := map[rune]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
	for _, s := range str {
		count += vowels[rune(s)]
	}
	return count
}
