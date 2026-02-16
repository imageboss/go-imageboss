package imageboss

import (
	"fmt"
	"strconv"
)

// Operation represents an ImageBoss operation (cdn, width, height, cover).
type Operation struct {
	kind   string // "cdn", "width", "height", "cover"
	width  int    // for width, cover
	height int    // for height, cover
	mode   string // for cover: "", "center", "smart", "attention", "entropy", "face", "north", etc.
}

// PathSegment returns the operation path segment (e.g. "cdn", "width", "cover:center").
func (o Operation) PathSegment() string {
	if o.mode != "" {
		return o.kind + ":" + o.mode
	}
	return o.kind
}

// Dimensions returns the dimensions segment (e.g. "" for cdn, "700" for width, "300x300" for cover).
func (o Operation) Dimensions() string {
	switch o.kind {
	case "cdn":
		return ""
	case "width":
		return strconv.Itoa(o.width)
	case "height":
		return strconv.Itoa(o.height)
	case "cover":
		return fmt.Sprintf("%dx%d", o.width, o.height)
	default:
		return ""
	}
}

// CDN returns a pass-through CDN operation (no resize).
func CDN() Operation {
	return Operation{kind: "cdn"}
}

// Width returns a width operation (fixed width, proportional height).
func Width(w int) Operation {
	return Operation{kind: "width", width: w}
}

// Height returns a height operation (fixed height, proportional width).
func Height(h int) Operation {
	return Operation{kind: "height", height: h}
}

// Cover returns a cover operation (exact width and height, default smart crop).
func Cover(w, h int) Operation {
	return Operation{kind: "cover", width: w, height: h}
}

// CoverMode returns a cover operation with a mode (center, smart, attention, entropy, face, north, etc.).
func CoverMode(w, h int, mode string) Operation {
	return Operation{kind: "cover", width: w, height: h, mode: mode}
}
