package main

import (
	"flag"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/quadrowin/exhaustivenew/pkg/analyzer"
)

func main() {
	flag.Bool("unsafeptr", false, "")

	a, err := analyzer.NewAnalyzer([]string{}, []string{})
	if err != nil {
		panic(err)
	}

	singlechecker.Main(a)
}
