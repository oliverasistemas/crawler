package url_extractor

import (
	"reflect"
	"testing"
)

func TestExtractValidUrls(t *testing.T) {
	baseURL := "https://www.example.com/a"

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
			name:     "Single valid link",
			htmlStr:  `<html><body><a href="/a/link1">Link 1</a></body></html>`,
			expected: []string{"https://www.example.com/a/link1"},
		},
		{
			name:     "Multiple valid links",
			htmlStr:  `<html><body><a href="/a/link1">Link 1</a><a href="/a/link2">Link 2</a></body></html>`,
			expected: []string{"https://www.example.com/a/link1", "https://www.example.com/a/link2"},
		},
		{
			name:     "Invalid links",
			htmlStr:  `<html><body><a href="https://www.anotherdomain.com/link1">Link 1</a><a href="/b/link2">Link 2</a></body></html>`,
			expected: []string{},
		},
		{
			name:     "Mixed valid and invalid links",
			htmlStr:  `<html><body><a href="/a/link1">Link 1</a><a href="https://www.anotherdomain.com/link2">Link 2</a><a href="/b/link3">Link 3</a></body></html>`,
			expected: []string{"https://www.example.com/a/link1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validUrls, err := ExtractValidUrls(baseURL, tt.htmlStr)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(validUrls, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, validUrls)
			}
		})
	}
}
