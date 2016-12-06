package htmlg

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// H1 returns a h1 element <h1>{{range .nodes}}{{.}}{{end}}</h1>.
func H1(nodes ...*html.Node) *html.Node {
	div := &html.Node{
		Type: html.ElementNode, Data: atom.H1.String(),
	}
	for _, n := range nodes {
		div.AppendChild(n)
	}
	return div
}

// H2 returns a h2 element <h2>{{range .nodes}}{{.}}{{end}}</h2>.
func H2(nodes ...*html.Node) *html.Node {
	div := &html.Node{
		Type: html.ElementNode, Data: atom.H2.String(),
	}
	for _, n := range nodes {
		div.AppendChild(n)
	}
	return div
}

// H3 returns a h3 element <h3>{{range .nodes}}{{.}}{{end}}</h3>.
func H3(nodes ...*html.Node) *html.Node {
	div := &html.Node{
		Type: html.ElementNode, Data: atom.H3.String(),
	}
	for _, n := range nodes {
		div.AppendChild(n)
	}
	return div
}

// H4 returns a h4 element <h4>{{range .nodes}}{{.}}{{end}}</h4>.
func H4(nodes ...*html.Node) *html.Node {
	div := &html.Node{
		Type: html.ElementNode, Data: atom.H4.String(),
	}
	for _, n := range nodes {
		div.AppendChild(n)
	}
	return div
}

// P returns a p element <p>{{range .nodes}}{{.}}{{end}}</p>.
func P(nodes ...*html.Node) *html.Node {
	div := &html.Node{
		Type: html.ElementNode, Data: atom.P.String(),
	}
	for _, n := range nodes {
		div.AppendChild(n)
	}
	return div
}

// DL returns a dl element <dl>{{range .nodes}}{{.}}{{end}}</dl>.
func DL(nodes ...*html.Node) *html.Node {
	div := &html.Node{
		Type: html.ElementNode, Data: atom.Dl.String(),
	}
	for _, n := range nodes {
		div.AppendChild(n)
	}
	return div
}

// DT returns a dt element <dt>{{range .nodes}}{{.}}{{end}}</dt>.
func DT(nodes ...*html.Node) *html.Node {
	div := &html.Node{
		Type: html.ElementNode, Data: atom.Dt.String(),
	}
	for _, n := range nodes {
		div.AppendChild(n)
	}
	return div
}

// DD returns a dd element <dd>{{range .nodes}}{{.}}{{end}}</dd>.
func DD(nodes ...*html.Node) *html.Node {
	div := &html.Node{
		Type: html.ElementNode, Data: atom.Dd.String(),
	}
	for _, n := range nodes {
		div.AppendChild(n)
	}
	return div
}
