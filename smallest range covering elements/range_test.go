package ranges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange1(t *testing.T) {
	r := [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}
	assert.Equal(t, []int{20, 24}, smallestRange(r))
}

func TestRange2(t *testing.T) {
	r := [][]int{{0, 2, 3, 4, 6}, {1, 3, 5, 7, 9}, {2, 5, 9, 13, 17}}
	assert.Equal(t, []int{1, 2}, smallestRange(r))
}

func TestRange3(t *testing.T) {
	r := [][]int{{-5, -4, -3, -2, -1}, {1, 2, 3, 4, 5}}
	assert.Equal(t, []int{-1, 1}, smallestRange(r))
}

func TestRange4(t *testing.T) {
	r := [][]int{{-5, -4, -3, -2, -1}, {-10, -9, 12, 4, 5}}
	assert.Equal(t, []int{-9, -5}, smallestRange(r))
}

func BenchmarkRange(b *testing.B) {
	r := [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}
	for i := 0; i <= b.N; i++ {
		assert.Equal(b, []int{20, 24}, smallestRange(r))
	}
}
