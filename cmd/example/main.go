package main

import (
	"flag"
	"fmt"

	lab2 "github.com/cunchyEnjoyers/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	res, _ := lab2.PrefixCalculate("+ * + + / / - - * 10 55 2 -9 2 3 99 -500 1 11")
	fmt.Println(res)
}
