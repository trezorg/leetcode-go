package parentheness

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParentheses(t *testing.T) {

	data := []struct {
		data   int
		result []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{1, []string{"()"}},
		{2, []string{"(())", "()()"}},
		{0, []string{}},
	}

	funcs := []func(int) []string{generateParenthesis, generateParenthesisUpdated}

	for _, f := range funcs {
		for _, v := range data {
			s := strings.Split((runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()), ".")[1]
			t.Run(fmt.Sprintf("%s-%v", s, v.data), func(t *testing.T) {
				assert.Equal(t, v.result, f(v.data))
			})
		}
	}

}

func BenchmarkTestGenerateParentheses(b *testing.B) {

	n := 5
	funcs := []func(int) []string{generateParenthesis, generateParenthesisUpdated}

	for _, f := range funcs {
		s := strings.Split((runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()), ".")[1]
		b.Run(s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(n)
			}
		})
	}

}
