package htmlg

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// H1 returns a h1 element <h1>{{range .nodes}}{{.}}{{end}}</h1>.
func H1(nodes ...*html.Node) *html.Node {
	h1 := &html.Node{
		Type: html.ElementNode, Data: atom.H1.String(),
	}
	for _, n := range nodes {
		h1.AppendChild(n)
	}
	return h1
}

// H2 returns a h2 element <h2>{{range .nodes}}{{.}}{{end}}</h2>.
func H2(nodes ...*html.Node) *html.Node {
	h2 := &html.Node{
		Type: html.ElementNode, Data: atom.H2.String(),
	}
	for _, n := range nodes {
		h2.AppendChild(n)
	}
	return h2
}

// H3 returns a h3 element <h3>{{range .nodes}}{{.}}{{end}}</h3>.
func H3(nodes ...*html.Node) *html.Node {
	h3 := &html.Node{
		Type: html.ElementNode, Data: atom.H3.String(),
	}
	for _, n := range nodes {
		h3.AppendChild(n)
	}
	return h3
}

// H4 returns a h4 element <h4>{{range .nodes}}{{.}}{{end}}</h4>.
func H4(nodes ...*html.Node) *html.Node {
	h4 := &html.Node{
		Type: html.ElementNode, Data: atom.H4.String(),
	}
	for _, n := range nodes {
		h4.AppendChild(n)
	}
	return h4
}

// P returns a p element <p>{{range .nodes}}{{.}}{{end}}</p>.
func P(nodes ...*html.Node) *html.Node {
	p := &html.Node{
		Type: html.ElementNode, Data: atom.P.String(),
	}
	for _, n := range nodes {
		p.AppendChild(n)
	}
	return p
}

// DL returns a dl element <dl>{{range .nodes}}{{.}}{{end}}</dl>.
func DL(nodes ...*html.Node) *html.Node {
	dl := &html.Node{
		Type: html.ElementNode, Data: atom.Dl.String(),
	}
	for _, n := range nodes {
		dl.AppendChild(n)
	}
	return dl
}

// DT returns a dt element <dt>{{range .nodes}}{{.}}{{end}}</dt>.
func DT(nodes ...*html.Node) *html.Node {
	dt := &html.Node{
		Type: html.ElementNode, Data: atom.Dt.String(),
	}
	for _, n := range nodes {
		dt.AppendChild(n)
	}
	return dt
}

// DD returns a dd element <dd>{{range .nodes}}{{.}}{{end}}</dd>.
func DD(nodes ...*html.Node) *html.Node {
	dd := &html.Node{
		Type: html.ElementNode, Data: atom.Dd.String(),
	}
	for _, n := range nodes {
		dd.AppendChild(n)
	}
	return dd
}
