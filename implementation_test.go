package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixCalculate(t *testing.T) {
	res, err := PrefixCalculate("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 2 - 3 * 5 +", res)
	}
}

func ExamplePrefixCalculate() {
	res, _ := PrefixCalculate("+ 2 2")
	fmt.Println(res)

	// Output:
	// 4
}
