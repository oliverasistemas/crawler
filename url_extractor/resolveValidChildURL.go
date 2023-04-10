package url_extractor

import (
	"net/url"
	"strings"
)

func resolveValidChildURL(base, child *url.URL) string {
	if child.IsAbs() {
		if base.Host != child.Host || !strings.HasPrefix(child.Path, base.Path) {
			return ""
		}

		if base.Path != "/" && !strings.HasSuffix(base.Path, "/") {
			if child.Path == base.Path || strings.HasPrefix(child.Path, base.Path+"/") {
				return child.String()
			}
			return ""
		}

		return child.String()
	} else {
		if !strings.HasPrefix(child.Path, "/") {
			base.Path = strings.TrimSuffix(base.Path, "/") + "/"
		}
		resolvedURL := base.ResolveReference(child)
		if strings.HasPrefix(resolvedURL.Path, base.Path) {
			return resolvedURL.String()
		}
		return ""
	}
}
