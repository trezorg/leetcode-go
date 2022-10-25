package trapwater

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func getFunctionName(f interface{}) string {
	s := strings.Split((runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()), ".")
	return s[len(s)-1]
}

func TestTrap(t *testing.T) {

	type value struct {
		data   []int
		result int
		trap   func([]int) int
	}

	data := []value{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6, trap},
		{[]int{4, 2, 0, 3, 2, 5}, 9, trap},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6, trap2},
		{[]int{4, 2, 0, 3, 2, 5}, 9, trap2},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%v-%v", v.data, getFunctionName(v.trap)), func(t *testing.T) {
			assert.Equal(t, v.result, v.trap(v.data))
		})
	}

}
