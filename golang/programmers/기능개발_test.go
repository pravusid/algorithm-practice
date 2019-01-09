package programmers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.EqualValues(
		t,
		[]int{2, 1},
		solution([]int{93, 30, 55}, []int{1, 30, 5}),
	)
}
