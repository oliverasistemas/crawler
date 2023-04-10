package url_extractor

import (
	"net/url"
	"testing"
)

func TestResolveValidChildURL(t *testing.T) {

	tests := []struct {
		baseURL  string
		name     string
		childURL string
		expected string
	}{
		{
			baseURL:  "https://www.example.com/a",
			name:     "Valid child URL",
			childURL: "https://www.example.com/a/link1",
			expected: "https://www.example.com/a/link1",
		},
		{
			baseURL:  "https://www.example.com",
			name:     "Valid child path",
			childURL: "/a/bc",
			expected: "https://www.example.com/a/bc",
		},
		{
			baseURL:  "https://www.example.com/a",
			name:     "Invalid child URL",
			childURL: "https://www.example.com/abc",
			expected: "",
		},
		{
			baseURL:  "https://www.example.com/a",
			name:     "Invalid hostname",
			childURL: "https://www.anotherdomain.com/abc",
			expected: "",
		},
		{
			baseURL:  "https://www.example.com/a",
			name:     "Invalid path",
			childURL: "/def",
			expected: "",
		},
		{
			baseURL:  "https://www.example.com/a/b",
			name:     "Valid path 1",
			childURL: "def",
			expected: "https://www.example.com/a/b/def",
		},
		{
			baseURL:  "https://www.example.com",
			name:     "Valid path 2",
			childURL: "def",
			expected: "https://www.example.com/def",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			child, _ := url.Parse(tt.childURL)
			base, _ := url.Parse(tt.baseURL)
			result := resolveValidChildURL(base, child)
			if result != tt.expected {
				t.Errorf("Expected: %q, got: %q", tt.expected, result)
			}
		})
	}
}
