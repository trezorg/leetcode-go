package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParentheses(t *testing.T) {

	type value struct {
		data   string
		result bool
	}

	data := []value{
		{"))", false},
		{"()", true},
		{"()()", true},
		{"()(*", true},
		{"()()*", true},
		{"()*)", true},
		{"*****", true},
		{"))**", false},
		{"))*", false},
		{")*", false},
		{"(*", true},
		{"(*()))*(", false},
	}

	funcs := []func(string) bool{checkValidStringBruteForceRecursive, checkValidStringBruteForce, checkValidString}

	for _, f := range funcs {
		s := strings.Split((runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()), ".")[1]
		for _, v := range data {
			t.Run(fmt.Sprintf("%s-%s", s, v.data), func(t *testing.T) {
				assert.Equal(t, v.result, f(v.data))
			})
		}
	}

}

func BenchmarkParentheses(b *testing.B) {

	str := []string{
		"(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())",
		"**************************************************))))))))))))))))))))))))))))))))))))))))))))))))))",
	}
	funcs := []func(string) bool{checkValidString}

	for _, f := range funcs {
		s := strings.Split((runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()), ".")[1]
		for _, ss := range str {
			b.Run(fmt.Sprintf("%s-%s", s, ss), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f(ss)
				}
			})
		}
	}

}
