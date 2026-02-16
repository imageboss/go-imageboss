package imageboss

import (
	"testing"
)

func TestParam_PathSegment(t *testing.T) {
	tests := []struct {
		opt  Option
		want string
	}{
		{Param("blur", "4"), "blur:4"},
		{Param("format", "auto"), "format:auto"},
		{Opt("blur", "4"), "blur:4"},
		{FormatAuto(), "format:auto"},
		{Blur(4), "blur:4"},
		{Download(), "download:1"},
		{DownloadFilename("my-image.png"), "download:my-image.png"},
	}
	for _, tt := range tests {
		if got := tt.opt.PathSegment(); got != tt.want {
			t.Errorf("PathSegment() = %s; want %s", got, tt.want)
		}
	}
}
