// https://guides.github.com/features/mastering-markdown/

package markdown

import "strings"

const (
	headerToken = '#'
	listToken   = '-'
)

var listTokens = map[rune]struct{}{
	'-': {},
	'*': {},
}

// Node is single markdown syntax representation
// Example: header, list, ...
type Node interface {
	String() string
}

type header struct {
	level int
	text  string
}

func NewHeader(level int, text string) Node {
	return header{
		level: level,
		text:  text,
	}
}

// Example: # Your title
func (h header) String() string {
	var builder strings.Builder

	for i := 0; i < h.level; i++ {
		builder.WriteString(string(headerToken))
	}

	builder.WriteString(" ")

	text := strings.TrimSpace(h.text)
	builder.WriteString(text)

	return builder.String()
}

type listItem struct {
	text string
}

func NewListItem(text string) Node {
	return listItem{
		text: text,
	}
}

func (i listItem) String() string {
	text := strings.TrimSpace(i.text)

	return string(listToken) + " " + text
}

func Equal(n1, n2 Node) bool {
	return n1.String() == n2.String()
}
