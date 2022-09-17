package main

// go build -buildmode=plugin plugin/exhaustivenew.go

import (
	"github.com/quadrowin/exhaustivenew/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

// GetAnalyzers must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	a, err := analyzer.NewAnalyzer([]string{}, []string{})
	if err != nil {
		panic(err)
	}

	return []*analysis.Analyzer{a}
}

// AnalyzerPlugin must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin
