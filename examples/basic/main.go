// Basic example: build ImageBoss URLs and print them.
// Run from repo root: go run ./examples/basic
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/imageboss/go"
)

func main() {
	b, err := imageboss.NewURLBuilder("mywebsite-images")
	if err != nil {
		log.Fatal(err)
	}

	path := "examples/02.jpg"

	fmt.Println("CDN:", b.CreateURL(path, imageboss.CDN()))
	fmt.Println("Width 700:", b.CreateURL(path, imageboss.Width(700)))
	fmt.Println("Cover 300x300:", b.CreateURL(path, imageboss.Cover(300, 300)))
	fmt.Println("Cover center 320x320:", b.CreateURL(path, imageboss.CoverMode(320, 320, "center")))
	fmt.Println("Width 700 + blur:", b.CreateURL(path, imageboss.Width(700), imageboss.Blur(4)))

	srcset := b.CreateSrcset(path, imageboss.CDN(), nil, imageboss.WithMinWidth(100), imageboss.WithMaxWidth(400))
	fmt.Println("Srcset (first 3):")
	for i, line := range strings.Split(srcset, ",\n") {
		if i >= 3 {
			break
		}
		fmt.Println(" ", strings.TrimSpace(line))
	}
}
