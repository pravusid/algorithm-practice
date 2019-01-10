package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScale(t *testing.T) {
	input := "abcd\nefgh\nijkl\nmnop"
	expected := "aabbccdd\naabbccdd\naabbccdd\neeffgghh\neeffgghh\neeffgghh\niijjkkll\niijjkkll\niijjkkll\nmmnnoopp\nmmnnoopp\nmmnnoopp"
	assert.Equal(t, expected, Scale(input, 2, 3))
}
