package url_extractor

import "net/url"

func ExtractValidUrls(baseURL string, htmlStr string) ([]string, error) {
	// Parse the base URL
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	// Extract all the links from the HTML string
	links, err := extractLinks(htmlStr)
	if err != nil {
		return nil, err
	}

	validUrls := []string{}

	// Iterate through the links and resolve them as valid child URLs
	for _, link := range links {
		child, err := url.Parse(link)
		if err != nil {
			continue
		}

		validChild := resolveValidChildURL(base, child)
		if validChild != "" {
			validUrls = append(validUrls, validChild)
		}
	}

	return validUrls, nil
}
