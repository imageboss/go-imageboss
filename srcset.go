package imageboss

import (
	"math"
	"strconv"
	"strings"
)

const (
	defaultMinWidth   = 100
	defaultMaxWidth   = 8192
	defaultTolerance  = 0.08
)

// DefaultWidths is the default list of widths from TargetWidths(100, 8192, 0.08).
var DefaultWidths = []int{
	100, 116, 135, 156, 181, 210, 244, 283, 328, 380, 441, 512,
	594, 689, 799, 927, 1075, 1247, 1446, 1678, 1946, 2257, 2619, 3038,
	3524, 4087, 4741, 5500, 6380, 7401, 8192,
}

// SrcsetOptions holds options for srcset generation.
type SrcsetOptions struct {
	MinWidth       int
	MaxWidth       int
	Tolerance      float64
	VariableQuality bool
}

// SrcsetOption configures srcset generation.
type SrcsetOption func(*SrcsetOptions)

// WithMinWidth sets the minimum width for fluid srcset.
func WithMinWidth(w int) SrcsetOption {
	return func(o *SrcsetOptions) {
		o.MinWidth = w
	}
}

// WithMaxWidth sets the maximum width for fluid srcset.
func WithMaxWidth(w int) SrcsetOption {
	return func(o *SrcsetOptions) {
		o.MaxWidth = w
	}
}

// WithTolerance sets the width tolerance (e.g. 0.08 for 8%).
func WithTolerance(t float64) SrcsetOption {
	return func(o *SrcsetOptions) {
		o.Tolerance = t
	}
}

// WithVariableQuality enables variable quality per DPR for fixed-width srcset.
func WithVariableQuality(v bool) SrcsetOption {
	return func(o *SrcsetOptions) {
		o.VariableQuality = v
	}
}

// TargetWidths returns a slice of target widths between min and max with the given tolerance.
func TargetWidths(minWidth, maxWidth int, tolerance float64) []int {
	r, err := validateRangeWithTolerance(minWidth, maxWidth, tolerance)
	if err != nil {
		return DefaultWidths
	}
	if r.minWidth == r.maxWidth {
		return []int{r.minWidth}
	}
	var widths []int
	start := float64(r.minWidth)
	for int(start) < r.maxWidth && int(start) < defaultMaxWidth {
		widths = append(widths, int(math.Round(start)))
		start = start * (1.0 + r.tolerance*2.0)
	}
	if len(widths) > 0 && widths[len(widths)-1] < r.maxWidth {
		widths = append(widths, r.maxWidth)
	}
	return widths
}

// CreateSrcset generates a srcset attribute string.
// If op has fixed dimensions (Width, Height, or Cover), a DPR-based srcset is generated.
// Otherwise a fluid width-based srcset is generated.
func (b *URLBuilder) CreateSrcset(path string, op Operation, options []Option, srcsetOpts ...SrcsetOption) string {
	opts := SrcsetOptions{
		MinWidth:        defaultMinWidth,
		MaxWidth:        defaultMaxWidth,
		Tolerance:       defaultTolerance,
		VariableQuality: true,
	}
	for _, fn := range srcsetOpts {
		fn(&opts)
	}

	// Fixed dimensions → DPR-based srcset
	if op.kind == "width" && op.width > 0 {
		return b.buildSrcsetDPR(path, Width(op.width), options, opts.VariableQuality)
	}
	if op.kind == "height" && op.height > 0 {
		return b.buildSrcsetDPR(path, Height(op.height), options, opts.VariableQuality)
	}
	if op.kind == "cover" && op.width > 0 && op.height > 0 {
		return b.buildSrcsetDPR(path, op, options, opts.VariableQuality)
	}

	// Fluid width-based srcset (use Width as kind; actual widths from TargetWidths)
	widths := TargetWidths(opts.MinWidth, opts.MaxWidth, opts.Tolerance)
	return b.CreateSrcsetFromWidths(path, Width(opts.MinWidth), options, widths)
}

// CreateSrcsetFromWidths builds a srcset with the given widths (fluid-width).
func (b *URLBuilder) CreateSrcsetFromWidths(path string, op Operation, options []Option, widths []int) string {
	if len(widths) == 0 {
		return ""
	}
	var entries []string
	for _, w := range widths {
		var segOp Operation
		switch op.kind {
		case "width":
			segOp = Width(w)
		case "height":
			segOp = Height(w)
		case "cover":
			segOp = Cover(op.width, op.height)
			if op.mode != "" {
				segOp = CoverMode(op.width, op.height, op.mode)
			}
		default:
			segOp = Width(w)
		}
		urlStr := b.CreateURL(path, segOp, options...)
		entries = append(entries, urlStr+" "+strconv.Itoa(w)+"w")
	}
	return strings.Join(entries, ",\n")
}

func (b *URLBuilder) buildSrcsetDPR(path string, op Operation, options []Option, variableQuality bool) string {
	dprQualities := map[string]string{"1": "75", "2": "50", "3": "35", "4": "23", "5": "20"}
	var entries []string
	for i := 1; i <= 5; i++ {
		ratio := strconv.Itoa(i)
		var opts []Option
		opts = append(opts, options...)
		if variableQuality {
			opts = append(opts, Opt("quality", dprQualities[ratio]))
		}
		urlStr := b.CreateURL(path, op, opts...)
		entries = append(entries, urlStr+" "+ratio+"x")
	}
	return strings.Join(entries, ",\n")
}
