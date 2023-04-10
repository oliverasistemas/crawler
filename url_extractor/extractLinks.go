package url_extractor

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

func extractLinks(htmlStr string) ([]string, error) {
	links := []string{}

	// Create a new tokenizer with the HTML string as input
	tokenizer := html.NewTokenizer(strings.NewReader(htmlStr))

	for {
		tokenType := tokenizer.Next()

		// Break the loop if we reach the end of the document
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			return nil, err
		}

		// If the token is a start tag and the tag is an <a> tag
		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "a" {
				// Iterate over the attributes of the <a> tag
				for _, attr := range token.Attr {
					// If the attribute is "href", add its value to the links slice
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}

	return links, nil
}
