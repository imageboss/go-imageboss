package imageboss

import (
	"strings"
	"testing"
)

func TestNewURLBuilder(t *testing.T) {
	b, err := NewURLBuilder("mywebsite-images")
	if err != nil {
		t.Fatal(err)
	}
	if b.Source() != "mywebsite-images" {
		t.Errorf("Source() = %s; want mywebsite-images", b.Source())
	}
	if b.BaseURL() != DefaultBaseURL {
		t.Errorf("BaseURL() = %s; want %s", b.BaseURL(), DefaultBaseURL)
	}
}

func TestNewURLBuilder_InvalidSource(t *testing.T) {
	_, err := NewURLBuilder("")
	if err == nil {
		t.Error("expected error for empty source")
	}
	_, err = NewURLBuilder("invalid source!")
	if err == nil {
		t.Error("expected error for invalid source")
	}
}

func TestCreateURL_CDN(t *testing.T) {
	b := MustNewURLBuilder("mywebsite-images")
	got := b.CreateURL("examples/02.jpg", CDN())
	want := "https://img.imageboss.me/mywebsite-images/cdn/examples/02.jpg"
	if got != want {
		t.Errorf("\ngot:  %s\nwant: %s", got, want)
	}
}

func TestCreateURL_Width(t *testing.T) {
	b := MustNewURLBuilder("mywebsite-images")
	got := b.CreateURL("examples/02.jpg", Width(700))
	want := "https://img.imageboss.me/mywebsite-images/width/700/examples/02.jpg"
	if got != want {
		t.Errorf("\ngot:  %s\nwant: %s", got, want)
	}
}

func TestCreateURL_Height(t *testing.T) {
	b := MustNewURLBuilder("mywebsite-images")
	got := b.CreateURL("examples/02.jpg", Height(700))
	want := "https://img.imageboss.me/mywebsite-images/height/700/examples/02.jpg"
	if got != want {
		t.Errorf("\ngot:  %s\nwant: %s", got, want)
	}
}

func TestCreateURL_Cover(t *testing.T) {
	b := MustNewURLBuilder("mywebsite-images")
	got := b.CreateURL("examples/02.jpg", Cover(300, 300))
	want := "https://img.imageboss.me/mywebsite-images/cover/300x300/examples/02.jpg"
	if got != want {
		t.Errorf("\ngot:  %s\nwant: %s", got, want)
	}
}

func TestCreateURL_CoverMode(t *testing.T) {
	b := MustNewURLBuilder("mywebsite-images")
	got := b.CreateURL("examples/02.jpg", CoverMode(320, 320, "center"))
	want := "https://img.imageboss.me/mywebsite-images/cover:center/320x320/examples/02.jpg"
	if got != want {
		t.Errorf("\ngot:  %s\nwant: %s", got, want)
	}
}

func TestCreateURL_WithOptions(t *testing.T) {
	b := MustNewURLBuilder("mywebsite-images")
	got := b.CreateURL("examples/02.jpg", Width(700), Opt("blur", "4"), Opt("format", "auto"))
	want := "https://img.imageboss.me/mywebsite-images/width/700/blur:4/format:auto/examples/02.jpg"
	if got != want {
		t.Errorf("\ngot:  %s\nwant: %s", got, want)
	}
}

func TestCreateURL_BaseURL(t *testing.T) {
	b := MustNewURLBuilder("mywebsite-images", WithBaseURL("https://custom.cdn.example.com"))
	got := b.CreateURL("img.jpg", CDN())
	want := "https://custom.cdn.example.com/mywebsite-images/cdn/img.jpg"
	if got != want {
		t.Errorf("\ngot:  %s\nwant: %s", got, want)
	}
}

func TestCreateURL_PathNormalized(t *testing.T) {
	b := MustNewURLBuilder("mywebsite-images")
	got := b.CreateURL("/examples/02.jpg", CDN())
	// Path is escaped and leading slash trimmed in segment join
	want := "https://img.imageboss.me/mywebsite-images/cdn/examples/02.jpg"
	if got != want {
		t.Errorf("\ngot:  %s\nwant: %s", got, want)
	}
}

func TestCreateURL_Signed(t *testing.T) {
	b := MustNewURLBuilder("mysecureimages", WithSecret("mysecret"))
	got := b.CreateURL("01.jpg", Width(500))
	if !strings.Contains(got, "?bossToken=") {
		t.Errorf("expected bossToken in URL: %s", got)
	}
	if len(got) < 64+len("?bossToken=") {
		t.Errorf("bossToken should be 64 hex chars: %s", got)
	}
	// Same path + secret must yield same token
	got2 := b.CreateURL("01.jpg", Width(500))
	if got != got2 {
		t.Errorf("signed URL should be deterministic: %s vs %s", got, got2)
	}
	// No secret → no token
	bNoSecret := MustNewURLBuilder("mywebsite-images")
	unsigned := bNoSecret.CreateURL("01.jpg", Width(500))
	if strings.Contains(unsigned, "bossToken=") {
		t.Errorf("expected no bossToken when secret not set: %s", unsigned)
	}
}

