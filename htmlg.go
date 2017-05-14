// Package htmlg contains helper funcs for generating HTML nodes and rendering them.
// Context-aware escaping is done just like in html/template, making it safe against code injection.
package htmlg

import (
	"bytes"
	"fmt"
	"io"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Text returns a plain text node.
func Text(s string) *html.Node {
	return &html.Node{
		Type: html.TextNode, Data: s,
	}
}

// Strong returns a strong text node.
func Strong(s string) *html.Node {
	strong := &html.Node{
		Type: html.ElementNode, Data: atom.Strong.String(),
	}
	strong.AppendChild(Text(s))
	return strong
}

// A returns an anchor element <a href="{{.href}}">{{.s}}</a>.
func A(s, href string) *html.Node {
	a := &html.Node{
		Type: html.ElementNode, Data: atom.A.String(),
		Attr: []html.Attribute{{Key: atom.Href.String(), Val: href}},
	}
	a.AppendChild(Text(s))
	return a
}

// Render renders HTML nodes, returning result as a string.
// Context-aware escaping is done just like in html/template when rendering nodes.
func Render(nodes ...*html.Node) string {
	var buf bytes.Buffer
	for _, node := range nodes {
		err := html.Render(&buf, node)
		if err != nil {
			// html.Render should only return a non-nil error if there's a problem writing to the supplied io.Writer.
			// We don't expect that to ever be the case (unless there's not enough memory), so panic.
			// If this ever happens in other situations, it's a bug in this library that should be reported and fixed.
			panic(fmt.Errorf("internal error: html.Render returned non-nil error, this is not expected to happen: %v", err))
		}
	}
	return buf.String()
}

// Component is anything that can render itself into HTML nodes.
type Component interface {
	Render() []*html.Node
}

// RenderComponents renders components into HTML, writing result to w.
// Context-aware escaping is done just like in html/template when rendering nodes.
func RenderComponents(w io.Writer, components ...Component) error {
	for _, c := range components {
		for _, node := range c.Render() {
			err := html.Render(w, node)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// RenderComponentsString renders components into HTML, returning result as a string.
// Context-aware escaping is done just like in html/template when rendering nodes.
func RenderComponentsString(components ...Component) string {
	var buf bytes.Buffer
	for _, c := range components {
		for _, node := range c.Render() {
			err := html.Render(&buf, node)
			if err != nil {
				// html.Render should only return a non-nil error if there's a problem writing to the supplied io.Writer.
				// We don't expect that to ever be the case (unless there's not enough memory), so panic.
				// If this ever happens in other situations, it's a bug in this library that should be reported and fixed.
				panic(fmt.Errorf("internal error: html.Render returned non-nil error, this is not expected to happen: %v", err))
			}
		}
	}
	return buf.String()
}
