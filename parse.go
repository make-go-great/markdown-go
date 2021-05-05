package markdown

import "strings"

const (
	defaultNodeLen = 10
)

// Parse return all markdown nodes from lines
func Parse(lines []string) []Node {
	if len(lines) == 0 {
		return nil
	}

	nodes := make([]Node, 0, defaultNodeLen)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, string(headerToken)) {
			nodes = append(nodes, parseHeader(line))
			continue
		}

		if strings.HasPrefix(line, string(defaultListToken)) ||
			strings.HasPrefix(line, string(alternativeListToken)) {
			nodes = append(nodes, parseListItem(line))
			continue
		}
	}

	return nodes
}

func parseHeader(line string) header {
	level := 0

	for _, c := range line {
		if c != headerToken {
			break
		}

		level++
	}

	line = strings.TrimLeft(line, string(headerToken))
	line = strings.TrimSpace(line)

	return header{
		level: level,
		text:  line,
	}
}

func parseListItem(line string) listItem {
	line = strings.TrimLeft(line, string(defaultListToken))
	line = strings.TrimLeft(line, string(alternativeListToken))
	line = strings.TrimSpace(line)

	return listItem{
		text: line,
	}
}
