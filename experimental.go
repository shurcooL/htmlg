package htmlg

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Div returns a div element <div>{{range .nodes}}{{.}}{{end}}</div>.
//
// Div is experimental and may be changed or removed.
func Div(nodes ...*html.Node) *html.Node {
	div := &html.Node{
		Type: html.ElementNode, Data: atom.Div.String(),
	}
	for _, n := range nodes {
		div.AppendChild(n)
	}
	return div
}

// DivClass returns a div element <div class="{{.class}}">{{range .nodes}}{{.}}{{end}}</div>.
//
// DivClass is experimental and may be changed or removed.
func DivClass(class string, nodes ...*html.Node) *html.Node {
	div := &html.Node{
		Type: html.ElementNode, Data: atom.Div.String(),
		Attr: []html.Attribute{{Key: atom.Class.String(), Val: class}},
	}
	for _, n := range nodes {
		div.AppendChild(n)
	}
	return div
}

// SpanClass returns a span element <span class="{{.class}}">{{range .nodes}}{{.}}{{end}}</span>.
//
// SpanClass is experimental and may be changed or removed.
func SpanClass(class string, nodes ...*html.Node) *html.Node {
	span := &html.Node{
		Type: html.ElementNode, Data: atom.Span.String(),
		Attr: []html.Attribute{{Key: atom.Class.String(), Val: class}},
	}
	for _, n := range nodes {
		span.AppendChild(n)
	}
	return span
}
