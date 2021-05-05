package markdown

import "strings"

// GenerateText return single string which represents all markdown nodes
func GenerateText(nodes []Node) string {
	lines := make([]string, len(nodes))
	for i, node := range nodes {
		lines[i] = node.String()
	}

	result := strings.Join(lines, string(NewlineToken)+string(NewlineToken))
	// Fix no newline at end of file
	result += string(NewlineToken)
	return result
}
