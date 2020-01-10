package fmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive1(t *testing.T) {
	nums := []int{7, 8, 9, 11, 12}
	assert.Equal(t, 1, firstMissingPositive(nums))
}

func TestFirstMissingPositive2(t *testing.T) {
	nums := []int{}
	assert.Equal(t, 1, firstMissingPositive(nums))
}

func TestFirstMissingPositive3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 6, firstMissingPositive(nums))
}

func TestFirstMissingPositive4(t *testing.T) {
	nums := []int{1, 2, -12, 6, 5}
	assert.Equal(t, 3, firstMissingPositive(nums))
}

func TestFirstMissingPositive5(t *testing.T) {
	nums := []int{5, 4, 3, 2, 0}
	assert.Equal(t, 1, firstMissingPositive(nums))
}

func TestFirstMissingPositive6(t *testing.T) {
	nums := []int{1}
	assert.Equal(t, 2, firstMissingPositive(nums))
}

func TestFirstMissingPositive7(t *testing.T) {
	nums := []int{1, 1}
	assert.Equal(t, 2, firstMissingPositive(nums))
}
