package roman

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToRoman(t *testing.T) {

	data := []struct {
		data   string
		result int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MMMCMXCVIII", 3998},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%s", v.data), func(t *testing.T) {
			assert.Equal(t, v.result, romanToInt(v.data))
		})
	}

}
