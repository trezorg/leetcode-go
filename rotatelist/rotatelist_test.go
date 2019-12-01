package rotatelist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateList(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}
	head := newList(s)
	res := head.toSlice()

	assert.Equal(t, s, res)

	res = head.toSlice()

	assert.Equal(t, s, res)
}

func TestPrintList(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}
	head := newList(s)

	assert.Equal(t, "1->2->3->4->5", head.toString())
}

func TestGetLastListElement(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}
	head := newList(s)
	last := head.tail()
	assert.Equal(t, 5, last.Val)

	s = []int{1}
	head = newList(s)
	last = head.tail()
	assert.Equal(t, 1, last.Val)

	s = []int{}
	head = newList(s)
	last = head.tail()
	assert.Nil(t, nil)
	assert.Nil(t, nil)

}

func TestLength(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}
	head := newList(s)
	assert.Equal(t, 5, head.len())

	s = []int{1}
	head = newList(s)
	assert.Equal(t, 1, head.len())

	s = []int{}
	head = newList(s)
	assert.Equal(t, 0, head.len())

}

func TestRotateList(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}
	head := newList(s)
	headSlice := head.toSlice()
	newHead := head.rotate().rotate()
	assert.Equal(t, headSlice, newHead.toSlice())
}

func TestRotateListRight(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}
	head := newList(s)
	newHead := rotateRight(head, 2)
	assert.Equal(t, []int{4, 5, 1, 2, 3}, newHead.toSlice())

	s = []int{1, 2, 3, 4, 5}
	head = newList(s)
	newHead = rotateRight(head, 1)
	assert.Equal(t, []int{5, 1, 2, 3, 4}, newHead.toSlice())

	s = []int{1, 2, 3, 4, 5}
	head = newList(s)
	newHead = rotateRight(head, 101)
	assert.Equal(t, []int{5, 1, 2, 3, 4}, newHead.toSlice())

	s = []int{1, 2, 3, 4, 5}
	head = newList(s)
	newHead = rotateRight(head, 10)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, newHead.toSlice())

}
