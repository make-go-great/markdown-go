package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderString(t *testing.T) {
	tests := []struct {
		name   string
		header Header
		want   string
	}{
		{
			name: "level 1",
			header: Header{
				level: 1,
				text:  "abc",
			},
			want: "# abc",
		},
		{
			name: "level 3",
			header: Header{
				level: 3,
				text:  "xyz",
			},
			want: "### xyz",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.header.String()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestListItemString(t *testing.T) {
	tests := []struct {
		name     string
		listItem ListItem
		want     string
	}{
		{
			name: "normal",
			listItem: ListItem{
				text: "abc",
			},
			want: "- abc",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.listItem.String()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		n1   Node
		n2   Node
		want bool
	}{
		{
			name: "header",
			n1:   NewHeader(0, "CHANGELOG"),
			n2:   NewHeader(0, "Changelog"),
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Equal(tc.n1, tc.n2)
			assert.Equal(t, tc.want, got)
		})
	}
}
