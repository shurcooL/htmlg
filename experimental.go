package htmlg

import (
	"context"
	"io"

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

// Span returns a span element <span>{{range .nodes}}{{.}}{{end}}</span>.
//
// Span is experimental and may be changed or removed.
func Span(nodes ...*html.Node) *html.Node {
	span := &html.Node{
		Type: html.ElementNode, Data: atom.Span.String(),
	}
	for _, n := range nodes {
		span.AppendChild(n)
	}
	return span
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

// ComponentContext is anything that can render itself into HTML nodes.
//
// ComponentContext is experimental and may be changed or removed.
type ComponentContext interface {
	Render(ctx context.Context) []*html.Node
}

// RenderComponentsContext renders components into HTML, writing result to w.
// Context-aware escaping is done just like in html/template when rendering nodes.
//
// RenderComponentsContext is experimental and may be changed or removed.
func RenderComponentsContext(ctx context.Context, w io.Writer, components ...ComponentContext) error {
	for _, c := range components {
		for _, node := range c.Render(ctx) {
			err := html.Render(w, node)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
