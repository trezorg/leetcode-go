package roman

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToRoman(t *testing.T) {

	data := []struct {
		data   int
		result string
	}{
		{3, "III"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
		{3998, "MMMCMXCVIII"},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%d", v.data), func(t *testing.T) {
			assert.Equal(t, v.result, intToRoman(v.data))
		})
	}

}
