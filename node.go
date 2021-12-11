// https://guides.github.com/features/mastering-markdown/

package markdown

import "strings"

const (
	headerToken = '#'
	listToken   = '-'
)

var listTokens = map[rune]struct{}{
	listToken: {},
	'*':       {},
}

// Node is single markdown syntax representation
// Example: header, list, ...
type Node interface {
	String() string
}

type Header struct {
	level int
	text  string
}

func NewHeader(level int, text string) Header {
	text = strings.TrimSpace(text)

	return Header{
		level: level,
		text:  text,
	}
}

// Example: # Your title
func (h Header) String() string {
	var builder strings.Builder

	for i := 0; i < h.level; i++ {
		builder.WriteString(string(headerToken))
	}

	builder.WriteString(" ")
	builder.WriteString(h.text)

	return builder.String()
}

type ListItem struct {
	text string
}

func NewListItem(text string) ListItem {
	text = strings.TrimSpace(text)

	return ListItem{
		text: text,
	}
}

func (i ListItem) String() string {
	return string(listToken) + " " + i.text
}

func Equal(n1, n2 Node) bool {
	return strings.EqualFold(n1.String(), n2.String())
}
