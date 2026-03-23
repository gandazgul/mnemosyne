package chunker

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

// ParseMarkdown converts a markdown string into a goldmark AST.
func ParseMarkdown(source []byte) ast.Node {
	reader := text.NewReader(source)
	parser := goldmark.DefaultParser()
	return parser.Parse(reader)
}
