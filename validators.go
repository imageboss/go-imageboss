package imageboss

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	sourceRegexp = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_-]*$`)
)

func validateSource(source string) (string, error) {
	s := strings.TrimSpace(source)
	if s == "" {
		return "", errors.New("imageboss: source cannot be empty")
	}
	if !sourceRegexp.MatchString(s) {
		return "", fmt.Errorf("imageboss: invalid source %q", source)
	}
	return s, nil
}

// validateDimension ensures width/height are positive.
func validateDimension(d int) error {
	if d <= 0 {
		return errors.New("imageboss: width and height must be positive")
	}
	return nil
}

type widthRange struct {
	minWidth  int
	maxWidth  int
	tolerance float64
}

func validateRangeWithTolerance(minW, maxW int, tol float64) (widthRange, error) {
	if minW < 0 || maxW < 0 {
		return widthRange{}, errors.New("imageboss: min and max width must be >= 0")
	}
	if maxW < minW {
		return widthRange{}, errors.New("imageboss: maxWidth must be >= minWidth")
	}
	if tol < 0.01 {
		return widthRange{}, errors.New("imageboss: tolerance must be >= 0.01")
	}
	return widthRange{minW, maxW, tol}, nil
}
