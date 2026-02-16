package imageboss

import (
	"testing"
)

func TestValidateSource(t *testing.T) {
	valid := []string{"mywebsite-images", "src1", "a", "A1", "source_name"}
	for _, s := range valid {
		got, err := validateSource(s)
		if err != nil {
			t.Errorf("validateSource(%q) err = %v", s, err)
		}
		if got != s {
			t.Errorf("validateSource(%q) = %q", s, got)
		}
	}
}

func TestValidateSource_Invalid(t *testing.T) {
	invalid := []string{"", " ", "has space", "no-dash!"}
	for _, s := range invalid {
		_, err := validateSource(s)
		if err == nil {
			t.Errorf("validateSource(%q) expected error", s)
		}
	}
}

func TestValidateRangeWithTolerance(t *testing.T) {
	_, err := validateRangeWithTolerance(100, 8192, 0.08)
	if err != nil {
		t.Errorf("expected no error: %v", err)
	}
	_, err = validateRangeWithTolerance(200, 100, 0.08)
	if err == nil {
		t.Error("expected error when max < min")
	}
	_, err = validateRangeWithTolerance(100, 200, 0.005)
	if err == nil {
		t.Error("expected error when tolerance < 0.01")
	}
}
