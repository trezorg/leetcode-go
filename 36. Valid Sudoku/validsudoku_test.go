package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var data = []struct {
	data   [][]byte
	result bool
}{
	{
		data: [][]byte{
			[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
		result: true,
	},
	{
		data: [][]byte{
			[]byte{'3', '3', '.', '.', '7', '.', '.', '.', '.'},
			[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
		result: false,
	},
}

func TestValidSudoku(t *testing.T) {
	for _, v := range data {
		t.Run(fmt.Sprintf("%v-%v", len(v.data), v.result), func(t *testing.T) {
			require.Equal(t, v.result, isValidSudoku(v.data))
		})
	}

}

func BenchmarkValidSudoku(b *testing.B) {
	for _, v := range data {
		b.Run(fmt.Sprintf("%v-%v", len(v.data), v.result), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				isValidSudoku(v.data)
			}
		})
	}

}
