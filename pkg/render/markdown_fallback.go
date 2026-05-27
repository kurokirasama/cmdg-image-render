package render

import "html"

func markdownPlainTextHTML(md []byte) []byte {
	return []byte("<pre>" + html.EscapeString(string(md)) + "</pre>")
}
