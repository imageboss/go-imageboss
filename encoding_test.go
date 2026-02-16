package imageboss

import (
	"testing"
)

func TestSanitizePath(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"examples/02.jpg", "examples/02.jpg"},
		{"/examples/02.jpg", "examples/02.jpg"},
		{"path/to/image.jpg", "path/to/image.jpg"},
	}
	for _, tt := range tests {
		got := sanitizePath(tt.in)
		if got != tt.want {
			t.Errorf("sanitizePath(%q) = %q; want %q", tt.in, got, tt.want)
		}
	}
}
