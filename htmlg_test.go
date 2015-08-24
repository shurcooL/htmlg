package htmlg_test

import (
	"fmt"
	"os"

	"github.com/shurcooL/htmlg"
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
