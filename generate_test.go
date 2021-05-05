package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name  string
		nodes []Node
		want  string
	}{
		{
			name: "normal",
			nodes: []Node{
				header{
					level: 1,
					text:  "header",
				},
				listItem{
					text: "item 1",
				},
				listItem{
					text: "item 2",
				},
			},
			want: "# header\n\n- item 1\n\n- item 2\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GenerateText(tc.nodes)
			assert.Equal(t, tc.want, got)
		})
	}
}
