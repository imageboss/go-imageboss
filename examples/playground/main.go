// Playground demonstrates building ImageBoss URLs and srcsets.
// Run with: go run github.com/imageboss/go/examples/playground
package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"

	"github.com/imageboss/go"
)

const page = `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>ImageBoss Go – Playground</title>
  <style>
    body { font-family: system-ui, sans-serif; max-width: 900px; margin: 2rem auto; padding: 0 1rem; }
    h1 { font-size: 1.5rem; }
    section { margin: 1.5rem 0; padding: 1rem; background: #f5f5f5; border-radius: 8px; }
    code { background: #eee; padding: 2px 6px; border-radius: 4px; word-break: break-all; }
    pre { overflow-x: auto; font-size: 0.85rem; }
    img { max-width: 100%; height: auto; border: 1px solid #ddd; border-radius: 4px; }
    .url { margin: 0.5rem 0; }
  </style>
</head>
<body>
  <h1>ImageBoss Go – Playground</h1>
  <p>Source: <strong>{{.Source}}</strong>. Replace with your ImageBoss source to see live images.</p>

  <section>
    <h2>CDN (pass-through)</h2>
    <div class="url"><code>{{.URLCDN}}</code></div>
    <img src="{{.URLCDN}}" alt="CDN" width="300" loading="lazy" onerror="this.style.display='none'">
  </section>

  <section>
    <h2>Width 400</h2>
    <div class="url"><code>{{.URLWidth}}</code></div>
    <img src="{{.URLWidth}}" alt="Width" width="400" loading="lazy" onerror="this.style.display='none'">
  </section>

  <section>
    <h2>Cover 300×300</h2>
    <div class="url"><code>{{.URLCover}}</code></div>
    <img src="{{.URLCover}}" alt="Cover" width="300" loading="lazy" onerror="this.style.display='none'">
  </section>

  <section>
    <h2>Width 400 + blur:4</h2>
    <div class="url"><code>{{.URLBlur}}</code></div>
    <img src="{{.URLBlur}}" alt="Blur" width="400" loading="lazy" onerror="this.style.display='none'">
  </section>

  <section>
    <h2>Srcset (fluid widths 100–400)</h2>
    <pre>{{.Srcset}}</pre>
    <img src="{{.URLWidth}}"
         srcset="{{.Srcset}}"
         sizes="(max-width: 600px) 100vw, 400px"
         alt="Responsive" loading="lazy" onerror="this.style.display='none'">
  </section>
</body>
</html>
`

func main() {
	source := "mywebsite-images"
	b, err := imageboss.NewURLBuilder(source)
	if err != nil {
		log.Fatal(err)
	}

	imagePath := "examples/02.jpg"

	tpl, err := template.New("").Parse(page)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if s := r.URL.Query().Get("source"); s != "" {
			var err error
			b, err = imageboss.NewURLBuilder(s)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			source = s
		}

		data := map[string]string{
			"Source":     source,
			"URLCDN":     b.CreateURL(imagePath, imageboss.CDN()),
			"URLWidth":   b.CreateURL(imagePath, imageboss.Width(400)),
			"URLCover":   b.CreateURL(imagePath, imageboss.Cover(300, 300)),
			"URLBlur":    b.CreateURL(imagePath, imageboss.Width(400), imageboss.Blur(4)),
			"Srcset":     b.CreateSrcset(imagePath, imageboss.CDN(), nil, imageboss.WithMinWidth(100), imageboss.WithMaxWidth(400)),
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	addr := "127.0.0.1:8080"
	fmt.Printf("Playground: http://%s/?source=YOUR_SOURCE\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
