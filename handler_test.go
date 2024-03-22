package lab2

import (
	"os"
	"strings"

	. "gopkg.in/check.v1"
)

type ComputeSuite struct {
	h *ComputeHandler
}

var _ = Suite(&ComputeSuite{})

func (s *ComputeSuite) SetUpTest(c *C) {
	s.h = &ComputeHandler{
		W: os.Stdout,
	}
}

func (s *ComputeSuite) TestCompute(c *C) {
	testCases := []struct {
		name     string
		input    string
		expected string
		errMsg   string
	}{
		{
			name:     "The output should match the input expression",
			input:    "+ * - 4 2 3 5",
			expected: "11",
			errMsg:   "",
		},
		{
			name:     "Returns an error for the invalid expression",
			input:    "+ * - 4 2 2 3 5",
			expected: "",
			errMsg:   "invalid prefix notation: insufficient operands",
		},
	}

	for _, tc := range testCases {
		b := strings.Builder{}
		s.h.R = strings.NewReader(tc.input)
		s.h.W = &b
		err := s.h.Compute()
		result := b.String()
		if tc.errMsg != "" {
			c.Assert(err, NotNil)
			c.Assert(err.Error(), Equals, tc.errMsg)
			c.Assert(result, Equals, tc.expected)
		} else {
			c.Assert(err, IsNil)
			c.Assert(result, Equals, tc.expected)
		}
	}
}

func (s *ComputeSuite) TestExampleComputeHandler_Compute(c *C) {
	s.h.R = strings.NewReader("4 2 - 3 * 5 +")
	err := s.h.Compute()
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "invalid prefix notation: not enough operands")
}
