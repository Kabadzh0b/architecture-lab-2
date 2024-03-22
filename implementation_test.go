package lab2

import (
	. "gopkg.in/check.v1"
	"testing"
	"fmt"
)

// Test CI
// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestIsOperand(c *C) {
	c.Assert(isOperand("123"), Equals, true)
	c.Assert(isOperand("abc"), Equals, false)
}

func (s *MySuite) TestIsOperator(c *C) {
	c.Assert(isOperator("+"), Equals, true)
	c.Assert(isOperator("-"), Equals, true)
	c.Assert(isOperator("*"), Equals, true)
	c.Assert(isOperator("/"), Equals, true)
	c.Assert(isOperator("x"), Equals, false)
}

func (s *MySuite) TestCalculate(c *C) {
	case1, err := calculate("+", "1", "2")
	c.Assert(err, IsNil)
	c.Assert(case1, Equals, "3")
	case2, err := calculate("*", "5", "43")
	c.Assert(err, IsNil)
	c.Assert(case2, Equals, "215")
	case3, err := calculate("-", "5", "7")
	c.Assert(err, IsNil)
	c.Assert(case3, Equals, "-2")
	case4, err := calculate("/", "5", "5")
	c.Assert(err, IsNil)
	c.Assert(case4, Equals, "1")
	case5, err := calculate("^", "5", "5")
	c.Assert(err, IsNil)
	c.Assert(case5, Equals, "3125")
	

	_, err = calculate("/", "1", "0")
	c.Assert(err, NotNil)
	_, err = calculate("/", "e", "0")
	c.Assert(err, NotNil)
	_, err = calculate("/", "44", "ttt0")
	c.Assert(err, NotNil)
}

func (s *MySuite) TestPrefixCalculate(c *C) {
	//Simple cases with 2-3 operands
	simpleCase1, err := PrefixCalculate("* + 1 2 5")
	c.Assert(err, IsNil)
	c.Assert(simpleCase1, Equals, "15")
	simpleCase2, err := PrefixCalculate("- / 9 3 1")
	c.Assert(err, IsNil)
	c.Assert(simpleCase2, Equals, "2")
	simpleCase3, err := PrefixCalculate("* * 2 22 2")
	c.Assert(err, IsNil)
	c.Assert(simpleCase3, Equals, "88")
	simpleCase4, err := PrefixCalculate("* /  2 2 55")
	c.Assert(err, IsNil)
	c.Assert(simpleCase4, Equals, "55")
	simpleCase5, err := PrefixCalculate("+ 100 - 2 9")
	c.Assert(err, IsNil)
	c.Assert(simpleCase5, Equals, "93")
	simpleCase6, err := PrefixCalculate("* 100 ^ 2 9")
	c.Assert(err, IsNil)
	c.Assert(simpleCase6, Equals, "51200")

	//Complex cases with 7-10 operands
	complexCase1, err := PrefixCalculate("+ 100 * - 2 9 + / 8 8 9")
	c.Assert(err, IsNil)
	c.Assert(complexCase1, Equals, "30")
	complexCase2, err := PrefixCalculate("- - + + 22 10 * 55 1 / 55 5 - 9 66")
	c.Assert(err, IsNil)
	c.Assert(complexCase2, Equals, "133")
	complexCase3, err := PrefixCalculate("+ * + + / / - - * 10 55 2 -9 2 3 99 -500 1 11")
	c.Assert(err, IsNil)
	c.Assert(complexCase3, Equals, "-298")
	complexCase4, err := PrefixCalculate("^ * + + / / - - * 10 55 2 -9 2 3 99 -500 1 2")
	c.Assert(err, IsNil)
	c.Assert(complexCase4, Equals, "95481")


	//Error cases
	_, err = PrefixCalculate("/ 1 0")
	c.Assert(err, NotNil)
	_, err = PrefixCalculate("/ rr 5")
	c.Assert(err, NotNil)
	_, err = PrefixCalculate("")
	c.Assert(err, NotNil)
	_, err = PrefixCalculate("+ / 1 1 88e")
	c.Assert(err, NotNil)
}

func (s *MySuite) Example(c *C)() {
	anotation := "+ 3 * 3 4"
	result, err := PrefixCalculate(anotation)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Result: ", result)
	//Output:
	//Result: 15
}