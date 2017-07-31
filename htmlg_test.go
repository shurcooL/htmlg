package htmlg_test

import (
	"fmt"
	"os"

	"github.com/shurcooL/component"
	"github.com/shurcooL/htmlg"
	"golang.org/x/net/html"
)

func Example() {
	// Context-aware escaping is done just like in html/template.
	html := htmlg.Render(
		htmlg.Text("Hi & how are you, "),
		htmlg.A("Gophers", "https://golang.org/"),
		htmlg.Text("? <script> is a cool gopher."),
	)

	fmt.Fprintln(os.Stdout, html)

	// Output:
	// Hi &amp; how are you, <a href="https://golang.org/">Gophers</a>? &lt;script&gt; is a cool gopher.
}

func ExampleAppendChildren() {
	div := htmlg.Div(
		htmlg.Text("Go "),
	)
	htmlg.AppendChildren(div, component.Link{
		Text:   "there",
		URL:    "https://golang.org",
		NewTab: true,
	}.Render()...)
	div.AppendChild(
		htmlg.Text("for more!"),
	)

	err := html.Render(os.Stdout, div)
	if err != nil {
		panic(err)
	}

	// Output:
	// <div>Go <a href="https://golang.org" target="_blank">there</a>for more!</div>
}
