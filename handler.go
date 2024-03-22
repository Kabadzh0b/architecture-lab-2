package lab2

import "io"

type ComputeHandler struct {
	R io.Reader
	W io.Writer
}

func (ch *ComputeHandler) Compute() error {
	var input []byte
	buffer := make([]byte, 1024)
	for {
		n, readErr := ch.R.Read(buffer)
		input = append(input, buffer[:n]...)
		if readErr != nil {
			if readErr == io.EOF {
				break
			} else {
				return readErr
			}
		}
	}
	res, err := PrefixCalculate(string(input))
	if err != nil {
		return err
	}
	_, writeErr := ch.W.Write([]byte(res))
	return writeErr
}