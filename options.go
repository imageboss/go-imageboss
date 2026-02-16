package imageboss

import (
	"fmt"
	"strings"
)

// Option is a path-segment option (e.g. blur:4, format:auto).
type Option interface {
	PathSegment() string
}

// Param is an Option built from key and value(s). Implements Option.
func Param(key string, values ...string) Option {
	return paramOption{key: key, values: values}
}

type paramOption struct {
	key    string
	values []string
}

func (p paramOption) PathSegment() string {
	if p.key == "" {
		return ""
	}
	if len(p.values) == 0 {
		return p.key
	}
	return p.key + ":" + strings.Join(p.values, ",")
}

// Opt creates a single key:value path segment for ImageBoss options.
func Opt(key, value string) Option {
	return Param(key, value)
}

// FormatAuto adds format:auto option.
func FormatAuto() Option {
	return Opt("format", "auto")
}

// Blur adds blur:N option (0–40, 2–4 recommended).
func Blur(n int) Option {
	return Opt("blur", fmt.Sprintf("%d", n))
}

// Download adds download option (force download). Use DownloadFilename for custom name.
func Download() Option {
	return Opt("download", "1")
}

// DownloadFilename adds download:filename option.
func DownloadFilename(filename string) Option {
	return Opt("download", filename)
}
