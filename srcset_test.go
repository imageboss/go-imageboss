package imageboss

import (
	"strings"
	"testing"
)

func TestTargetWidths(t *testing.T) {
	got := TargetWidths(100, 380, 0.08)
	want := []int{100, 116, 135, 156, 181, 210, 244, 283, 328, 380}
	if len(got) != len(want) {
		t.Fatalf("len(TargetWidths(100,380,0.08)) = %d; want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("TargetWidths[%d] = %d; want %d", i, got[i], want[i])
		}
	}
}

func TestTargetWidths_Single(t *testing.T) {
	got := TargetWidths(300, 300, 0.08)
	if len(got) != 1 || got[0] != 300 {
		t.Errorf("TargetWidths(300,300,0.08) = %v; want [300]", got)
	}
}

func TestCreateSrcsetFromWidths(t *testing.T) {
	b := MustNewURLBuilder("demo")
	got := b.CreateSrcsetFromWidths("image.jpg", Width(100), nil, []int{100, 200, 300})
	want := "https://img.imageboss.me/demo/width/100/image.jpg 100w,\n" +
		"https://img.imageboss.me/demo/width/200/image.jpg 200w,\n" +
		"https://img.imageboss.me/demo/width/300/image.jpg 300w"
	if got != want {
		t.Errorf("\ngot:  %s\nwant: %s", got, want)
	}
}

func TestCreateSrcset_Fluid(t *testing.T) {
	b := MustNewURLBuilder("demo")
	got := b.CreateSrcset("image.png", CDN(), nil, WithMinWidth(100), WithMaxWidth(380), WithTolerance(0.08))
	lines := strings.Split(got, ",\n")
	if len(lines) != 10 {
		t.Errorf("CreateSrcset fluid: got %d entries; want 10", len(lines))
	}
	for _, line := range lines {
		if !strings.Contains(line, "width/") {
			t.Errorf("expected width in entry: %s", line)
		}
		if !strings.HasSuffix(strings.TrimSpace(line), "w") {
			t.Errorf("expected Nw descriptor: %s", line)
		}
	}
}

func TestCreateSrcset_FixedWidth(t *testing.T) {
	b := MustNewURLBuilder("demo")
	got := b.CreateSrcset("image.png", Width(800), nil)
	lines := strings.Split(got, ",\n")
	if len(lines) != 5 {
		t.Errorf("CreateSrcset fixed width: got %d entries; want 5 (1x-5x)", len(lines))
	}
	for i, line := range lines {
		if !strings.Contains(line, "width/800") {
			t.Errorf("entry %d: expected width/800: %s", i, line)
		}
	}
}

func TestCreateSrcset_EmptyWidths(t *testing.T) {
	b := MustNewURLBuilder("demo")
	got := b.CreateSrcsetFromWidths("image.jpg", Width(100), nil, nil)
	if got != "" {
		t.Errorf("CreateSrcsetFromWidths(nil widths) = %q; want \"\"", got)
	}
}
