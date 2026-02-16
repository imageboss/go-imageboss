package imageboss

import (
	"testing"
)

func TestOperation_PathSegment(t *testing.T) {
	tests := []struct {
		op   Operation
		want string
	}{
		{CDN(), "cdn"},
		{Width(700), "width"},
		{Height(500), "height"},
		{Cover(300, 300), "cover"},
		{CoverMode(320, 320, "center"), "cover:center"},
		{CoverMode(320, 320, "smart"), "cover:smart"},
	}
	for _, tt := range tests {
		if got := tt.op.PathSegment(); got != tt.want {
			t.Errorf("%+v.PathSegment() = %s; want %s", tt.op, got, tt.want)
		}
	}
}

func TestOperation_Dimensions(t *testing.T) {
	tests := []struct {
		op   Operation
		want string
	}{
		{CDN(), ""},
		{Width(700), "700"},
		{Height(500), "500"},
		{Cover(300, 300), "300x300"},
		{CoverMode(320, 420, "center"), "320x420"},
	}
	for _, tt := range tests {
		if got := tt.op.Dimensions(); got != tt.want {
			t.Errorf("%+v.Dimensions() = %s; want %s", tt.op, got, tt.want)
		}
	}
}
