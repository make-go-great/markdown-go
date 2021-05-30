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

		isListToken := false
		for listTok := range listTokens {
			if strings.HasPrefix(line, string(listTok)) {
				isListToken = true
				break
			}
		}

		if isListToken {
			nodes = append(nodes, parseListItem(line))
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
	for listTok := range listTokens {
		line = strings.TrimLeft(line, string(listTok))
	}

	line = strings.TrimSpace(line)

	return listItem{
		text: line,
	}
}
