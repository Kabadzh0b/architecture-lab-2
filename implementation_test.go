package lab2

import (
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPrefixCalculate(c *C) {
	// Test a case where the prefix notation is valid.
	res, err := PrefixCalculate("+ 5 * - 4 2 3")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "11") // The expected result of "+ 5 * - 4 2 3" is "-1".

	// Test a case where the prefix notation is invalid.
	res, err = PrefixCalculate("+ 5 * - 4 3")
	c.Assert(err, NotNil) // An error should be returned because the prefix notation is invalid.

	// Test a case where the prefix notation results in division by zero.
	res, err = PrefixCalculate("/ 5 0")
	c.Assert(err, NotNil) // An error should be returned because division by zero is not allowed.
}
