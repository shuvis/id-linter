package main

import (
	"github.com/shuvis/custom-linters/internal/id"

	"golang.org/x/tools/go/analysis"
)

var AnalyzerPlugin plugin

type plugin struct{}

func (plugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{id.Analyzer}
}
