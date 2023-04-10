package url_extractor

import (
	"reflect"
	"testing"
)

func TestExtractLinks(t *testing.T) {
	tests := []struct {
		name     string
		htmlStr  string
		expected []string
	}{
		{
			name:     "No links",
			htmlStr:  "<html><body><p>No links here</p></body></html>",
			expected: []string{},
		},
		{
			name:     "Single link",
			htmlStr:  `<html><body><a href="/link1">Link 1</a></body></html>`,
			expected: []string{"/link1"},
		},
		{
			name:     "Multiple links",
			htmlStr:  `<html><body><a href="/link1">Link 1</a><a href="/link2">Link 2</a></body></html>`,
			expected: []string{"/link1", "/link2"},
		},
		{
			name:     "Link with other attributes",
			htmlStr:  `<html><body><a href="/link1" class="link" id="link1">Link 1</a></body></html>`,
			expected: []string{"/link1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			links, err := extractLinks(tt.htmlStr)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(links, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, links)
			}
		})
	}
}
