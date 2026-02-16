// Package imageboss provides a Go client for generating ImageBoss image URLs.
//
// ImageBoss is an image resize, compress and CDN service. This library builds
// URLs in the form:
//
//	https://img.imageboss.me/:source/:operation/:dimensions/:options/path/to/image.jpg
//
// See https://imageboss.me/docs for the full API.
package imageboss

import (
	"strings"
)

const (
	// DefaultBaseURL is the default ImageBoss CDN base URL.
	DefaultBaseURL = "https://img.imageboss.me"
	// LibVersion is the library version (for reference and release tracking).
	// It is not appended to generated URLs.
	LibVersion = "go-v1.0.1"
)

// URLBuilder builds ImageBoss image URLs.
type URLBuilder struct {
	baseURL  string
	source   string
	useHTTPS bool
	secret   string // optional; when set, URLs are signed with bossToken (HMAC SHA-256 of path)
}

// BuilderOption configures a URLBuilder.
type BuilderOption func(*URLBuilder)

// NewURLBuilder creates a URLBuilder for the given ImageBoss source.
// The source is the name configured in your ImageBoss dashboard (e.g. "mywebsite-images").
func NewURLBuilder(source string, options ...BuilderOption) (*URLBuilder, error) {
	source, err := validateSource(source)
	if err != nil {
		return nil, err
	}
	b := &URLBuilder{
		baseURL:  DefaultBaseURL,
		source:   source,
		useHTTPS: true,
	}
	for _, fn := range options {
		fn(b)
	}
	return b, nil
}

// MustNewURLBuilder is like NewURLBuilder but panics on error.
func MustNewURLBuilder(source string, options ...BuilderOption) *URLBuilder {
	b, err := NewURLBuilder(source, options...)
	if err != nil {
		panic(err)
	}
	return b
}

// WithBaseURL sets a custom base URL (e.g. for testing or custom CDN).
func WithBaseURL(baseURL string) BuilderOption {
	return func(b *URLBuilder) {
		b.baseURL = strings.TrimSuffix(baseURL, "/")
	}
}

// WithHTTPS sets whether to use HTTPS in generated URLs.
func WithHTTPS(useHTTPS bool) BuilderOption {
	return func(b *URLBuilder) {
		b.useHTTPS = useHTTPS
	}
}

// SetUseHTTPS sets whether to use HTTPS.
func (b *URLBuilder) SetUseHTTPS(use bool) {
	b.useHTTPS = use
}

// WithSecret sets the secret token for signing URLs. When set, generated URLs
// include a bossToken query parameter (HMAC SHA-256 of the path). Enable
// "I want extra security. My URLs need to be signed" for the source in the
// ImageBoss dashboard to get the secret. See https://imageboss.me/docs/security.
func WithSecret(secret string) BuilderOption {
	return func(b *URLBuilder) {
		b.secret = secret
	}
}

// SetSecret sets the secret for signing URLs. See WithSecret.
func (b *URLBuilder) SetSecret(secret string) {
	b.secret = secret
}

// Source returns the configured source name.
func (b *URLBuilder) Source() string {
	return b.source
}

// BaseURL returns the base URL (without path).
func (b *URLBuilder) BaseURL() string {
	return b.baseURL
}

// CreateURL builds an ImageBoss URL for the given image path, operation, and options.
// Path is the image path relative to your source (e.g. "examples/02.jpg").
// Operation is one of: CDN(), Width(w), Height(h), Cover(w, h), or CoverMode(w, h, mode).
// Options are path-segment options like Opt("blur", "4") or Opt("format", "auto").
func (b *URLBuilder) CreateURL(path string, op Operation, options ...Option) string {
	path = sanitizePath(path)
	segments := []string{strings.TrimSuffix(b.baseURL, "/"), b.source, op.PathSegment()}
	if d := op.Dimensions(); d != "" {
		segments = append(segments, d)
	}
	for _, o := range options {
		if s := o.PathSegment(); s != "" {
			segments = append(segments, s)
		}
	}
	segments = append(segments, path)
	urlStr := strings.Join(segments, "/")
	if b.secret != "" {
		pathForSigning := "/" + strings.Join(segments[1:], "/")
		urlStr += "?bossToken=" + signPath(b.secret, pathForSigning)
	}
	return urlStr
}

// CreateURLWithParams builds a URL with CDN operation and the given options.
// For resize operations use CreateURL with Width, Height, or Cover.
func (b *URLBuilder) CreateURLWithParams(path string, params ...Option) string {
	return b.CreateURL(path, CDN(), params...)
}
