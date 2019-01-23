// implements conversions from markdown string to a structure,
// so that people could convert it into other forms.
package mdconv

import (
	"gopkg.in/russross/blackfriday.v2"
)

type Markdown struct {
	Option Option
}

func (m *Markdown) Parse(input []byte) *Node {
	p.block(input)
	// Walk the tree and finish up some of unfinished blocks
	for p.tip != nil {
		p.finalize(p.tip)
	}
	// Walk the tree again and process inline markdown in each block
	p.doc.Walk(func(node *Node, entering bool) WalkStatus {
		if node.Type == Paragraph || node.Type == Heading || node.Type == TableCell {
			p.inline(node, node.content)
			node.content = nil
		}
		return GoToNext
	})
	p.parseRefsToAST()
	return p.doc
}





func MarkdownProc(input []byte, opts ...blackfriday.Option) (ast *blackfriday.Node) {
	optList := []blackfriday.Option{
		blackfriday.WithExtensions(
			blackfriday.CommonExtensions,
		),
	}
	optList = append(optList, opts...)
	parser := blackfriday.New(optList...)
	ast = parser.Parse(input)
	return ast
}
