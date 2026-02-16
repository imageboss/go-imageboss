package imageboss

import (
	"net/url"
	"strings"
)

// sanitizePath normalizes the image path for use in the URL.
func sanitizePath(path string) string {
	path = strings.TrimSpace(path)
	if path == "" {
		return path
	}
	path = strings.TrimPrefix(path, "/")
	// Path-escape each segment to handle spaces and special chars.
	parts := strings.Split(path, "/")
	for i, p := range parts {
		parts[i] = url.PathEscape(p)
	}
	return strings.Join(parts, "/")
}
