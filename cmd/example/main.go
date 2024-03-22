package main

import (
	"flag"
	"strings"
	"io"
	"os"

	lab2 "github.com/cunchyEnjoyers/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File path to read expression from")
	outputFile      = flag.String("o", "", "Path to the file in which output the result")
)

func main() {
	flag.Parse()
	if len(*inputExpression) < 1 && len(*inputFile) < 1 {
		os.Stderr.WriteString("You should provide atleast one flag as arg")
		os.Exit(1)
	}
	var readStream io.Reader
	var writeStream io.Writer
	if len(*inputExpression) > 0 {
		readStream = strings.NewReader(*inputExpression)
	} else {
		file, err := os.Open(*inputFile)
		if err != nil {
			os.Stderr.WriteString("Incorrect file to read")
			os.Exit(1)
		}
		defer file.Close()
		readStream = file
	}
	if len(*outputFile) < 1 {
		writeStream = os.Stdout
	} else {
    file, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			os.Stderr.WriteString("Cannot write to output file")
			os.Exit(1)
		}
		defer file.Close()
		writeStream = file
	}
	handler := &lab2.ComputeHandler{
		R: readStream,
		W: writeStream,
	}
	err := handler.Compute()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}
