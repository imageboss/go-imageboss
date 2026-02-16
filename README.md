# @imageboss/go

Go client library for generating [ImageBoss](https://imageboss.me) image URLs. Build URLs with operations (resize, cover, CDN), path-segment options, and optional signing.

## Installation

```bash
go get github.com/imageboss/go
```

## Usage

Create a URL builder with your ImageBoss **source** (configured in your [dashboard](https://imageboss.me/dashboard/sources)), then build URLs with operations and options.

### Basic URL

```go
package main

import (
    "fmt"
    "github.com/imageboss/go"
)

func main() {
    b, err := imageboss.NewURLBuilder("mywebsite-images")
    if err != nil {
        panic(err)
    }

    // CDN (pass-through, no resize)
    url := b.CreateURL("examples/02.jpg", imageboss.CDN())
    fmt.Println(url)
    // https://img.imageboss.me/mywebsite-images/cdn/examples/02.jpg

    // Fixed width (proportional height)
    url = b.CreateURL("examples/02.jpg", imageboss.Width(700))
    // https://img.imageboss.me/mywebsite-images/width/700/examples/02.jpg

    // Fixed height
    url = b.CreateURL("examples/02.jpg", imageboss.Height(700))

    // Cover (exact dimensions, smart crop by default)
    url = b.CreateURL("examples/02.jpg", imageboss.Cover(300, 300))

    // Cover with mode (center, smart, attention, entropy, face, north, etc.)
    url = b.CreateURL("examples/02.jpg", imageboss.CoverMode(320, 320, "center"))

    // With options (path-segment options)
    url = b.CreateURL("examples/02.jpg", imageboss.Width(700), imageboss.Blur(4), imageboss.FormatAuto())
    // https://img.imageboss.me/mywebsite-images/width/700/blur:4/format:auto/examples/02.jpg
}
```

### Builder options

- `imageboss.WithBaseURL("https://custom.cdn.example.com")` – custom base URL.
- `imageboss.WithHTTPS(false)` – use HTTP (default: true).
- `imageboss.WithSecret(secret)` – sign URLs with a `bossToken` query parameter (see [Signed URLs](#signed-urls)).

### Signed URLs

For sources with [signed URLs](https://imageboss.me/docs/security) enabled in the ImageBoss dashboard, pass your secret when creating the builder. The library signs the URL path with HMAC SHA-256 and appends `?bossToken=<hex>`.

```go
b, err := imageboss.NewURLBuilder("mysecureimages", imageboss.WithSecret(os.Getenv("IMAGEBOSS_SECRET")))
if err != nil {
    panic(err)
}
url := b.CreateURL("images/photo.jpg", imageboss.Width(500))
// https://img.imageboss.me/mysecureimages/width/500/images/photo.jpg?bossToken=...
```

### Srcset

Fluid (width-based) srcset:

```go
srcset := b.CreateSrcset("image.png", imageboss.CDN(), nil,
    imageboss.WithMinWidth(100),
    imageboss.WithMaxWidth(380),
    imageboss.WithTolerance(0.08))
// Use in HTML: <img srcset="..." sizes="...">
```

Custom widths:

```go
srcset := b.CreateSrcsetFromWidths("image.jpg", imageboss.Width(100), nil, []int{100, 200, 300, 400})
```

Fixed dimensions (DPR-based, 1x–5x):

```go
srcset := b.CreateSrcset("image.png", imageboss.Width(800), nil)
```

### Helpers

- `imageboss.TargetWidths(minWidth, maxWidth, tolerance)` – list of target widths for custom srcsets.
- `imageboss.MustNewURLBuilder(source, opts...)` – panics on invalid source instead of returning an error.

## Operations and options (ImageBoss API)

- **Operations:** `cdn`, `width`, `height`, `cover` (with optional mode: `center`, `smart`, `attention`, `entropy`, `face`, `north`, etc.).
- **Options (path segments):** e.g. `blur:4`, `format:auto`, `download:1`, `fill-color:ffffff`. Use `imageboss.Opt("key", "value")` or helpers like `imageboss.Blur(4)`, `imageboss.FormatAuto()`, `imageboss.Download()`.

See [ImageBoss docs](https://imageboss.me/docs) for the full API.

## Testing

```bash
go test ./...
```

## Playground

Run the example server to try URLs in the browser:

```bash
go run ./examples/playground
# Open http://127.0.0.1:8080/?source=YOUR_IMAGEBOSS_SOURCE
```

Or print URLs from the basic example:

```bash
go run ./examples/basic
```im

## Releasing (Changesets)

This repo uses [Changesets](https://github.com/changesets/changesets) for versioning and changelogs.

1. **Add a changeset** after making a change: `npx changeset` (choose bump type and add a summary).
2. **Version** when ready to release: `npx changeset version` (updates `package.json` and `CHANGELOG.md`).
3. **Tag the Go module**: e.g. `git tag v1.0.1 && git push origin v1.0.1`.

See [.changeset/README.md](.changeset/README.md) for details.

## License

See [LICENSE](LICENSE) in this repository.
