package config

import (
	"flag"
	"fmt"
)

type Params struct {
	InputFile  string
	OutputFile string
	WindowSize int
}

func ParseParams() (*Params, error) {
	params := &Params{}
	flag.StringVar(&params.InputFile, "in", "", "Input file to be processed (required)")
	flag.StringVar(&params.OutputFile, "out", "", "Output file (required)")
	flag.IntVar(&params.WindowSize, "size", 10, "sliding window size")

	flag.Parse()

	if params.InputFile == "" || params.OutputFile == "" {
		return nil, fmt.Errorf("both `in` and `out` must be defined. Run with `-h` to see all possible parameters")
	}

	return params, nil
}
