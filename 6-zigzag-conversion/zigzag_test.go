package zigzag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZigZag(t *testing.T) {

	data := []struct {
		data   string
		result string
		rows   int
	}{
		{"PAYPALISHIRING", "PAHNAPLSIIGYIR", 3},
		{"PAYPALISHIRING", "PINALSIGYAHRPI", 4},
		{"A", "A", 1},
	}

	for _, v := range data {
		t.Run(fmt.Sprintf("%s-%d", v.data, v.rows), func(t *testing.T) {
			assert.Equal(t, v.result, convert(v.data, v.rows))
		})
	}

}
