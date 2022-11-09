package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var data = []struct {
	data   []int
	result [][]int
}{
	{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	{[]int{1, 2, 3, 4}, [][]int{
		{1, 2, 3, 4},
		{1, 2, 4, 3},
		{1, 3, 2, 4},
		{1, 3, 4, 2},
		{1, 4, 2, 3},
		{1, 4, 3, 2},
		{2, 1, 3, 4},
		{2, 1, 4, 3},
		{2, 3, 1, 4},
		{2, 3, 4, 1},
		{2, 4, 1, 3},
		{2, 4, 3, 1},
		{3, 1, 2, 4},
		{3, 1, 4, 2},
		{3, 2, 1, 4},
		{3, 2, 4, 1},
		{3, 4, 1, 2},
		{3, 4, 2, 1},
		{4, 1, 2, 3},
		{4, 1, 3, 2},
		{4, 2, 1, 3},
		{4, 2, 3, 1},
		{4, 3, 1, 2},
		{4, 3, 2, 1},
	}},
}

func sortSlice(a [][]int) {
	sort.Slice(a, func(i, j int) bool {
		a1 := a[i]
		a2 := a[j]
		for t := 0; t < len(a1); t++ {
			if a1[t] < a2[t] {
				return true
			} else if a1[t] > a2[t] {
				return false
			}
		}
		return true
	})
}

func TestPermute(t *testing.T) {

	funcs := []func([]int) [][]int{permute, permuteIterable, permuteNext}

	for _, f := range funcs {
		for _, v := range data {
			l := make([]int, len(v.data))
			copy(l, v.data)
			s := strings.Split((runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()), ".")[1]
			t.Run(fmt.Sprintf("%s-%v", s, l), func(t *testing.T) {
				result := f(l)
				sortSlice(result)
				sortSlice(v.result)
				require.Equalf(t, v.result, result, "len expected: %d. len result: %d", len(v.result), len(result))
			})
		}
	}
}

func BenchmarkPermute(b *testing.B) {

	funcs := []func([]int) [][]int{permute, permuteIterable, permuteNext}

	var data = []struct {
		data []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, f := range funcs {
		for _, v := range data {
			s := strings.Split((runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()), ".")[1]
			b.Run(fmt.Sprintf("%s-%v", s, v.data), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f(v.data)
				}
			})
		}
	}
}
