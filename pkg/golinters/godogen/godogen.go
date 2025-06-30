package godogen

import (
	"github.com/lukasngl/godogen"

	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(godogen.Analyzer).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
