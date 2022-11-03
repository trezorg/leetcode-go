package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

var data = []struct {
	s      string
	locked string
	result bool
}{
	{"()()()()()()()()()()()()()()()()()()()()()()()()()()()()()()()()()()", "11111111111111111111111111111111111111111111111111111111111111111111", true},
	{"))()))", "010100", true},
	{"()()", "0000", true},
	{")", "0", false},
	{"))", "00", true},
	{"))", "10", false},
	{"))))", "0101", true},
}

func TestIsValid(t *testing.T) {

	for _, v := range data {
		t.Run(fmt.Sprintf("%s-%v", v.s, v.locked), func(t *testing.T) {
			assert.Equal(t, v.result, canBeValid(v.s, v.locked))
		})
	}

}

func BenchmarkIsValid(b *testing.B) {

	funcs := []func(string, string) bool{canBeValid}

	for _, f := range funcs {
		s := strings.Split((runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()), ".")[1]
		d := data[0]
		b.Run(s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(d.s, d.locked)
			}
		})
	}

}
