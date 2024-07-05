## go-fitz
[![Build Status](https://github.com/JoKr4/go-fitz-nolibs/actions/workflows/test.yml/badge.svg)](https://github.com/JoKr4/go-fitz-nolibs/actions)
[![GoDoc](https://godoc.org/github.com/JoKr4/go-fitz-nolibs?status.svg)](https://godoc.org/github.com/JoKr4/go-fitz-nolibs)
[![Go Report Card](https://goreportcard.com/badge/github.com/JoKr4/go-fitz-nolibs?branch=master)](https://goreportcard.com/report/github.com/JoKr4/go-fitz-nolibs)

Go wrapper for [MuPDF](http://mupdf.com/) fitz library that can extract pages from PDF, EPUB and MOBI documents as images, text, html or svg.

### Build tags

* `extlib` - use external MuPDF library
* `static` - build with static external MuPDF library (used with `extlib`)
* `pkgconfig` - enable pkg-config (used with `extlib`)
* `musl` - use musl compiled library
    
### Example
```go
package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/JoKr4/go-fitz-nolibs"
)

func main() {
	doc, err := fitz.New("test.pdf")
	if err != nil {
		panic(err)
	}

	defer doc.Close()

	tmpDir, err := os.MkdirTemp(os.TempDir(), "fitz")
	if err != nil {
		panic(err)
	}

	// Extract pages as images
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.jpg", n)))
		if err != nil {
			panic(err)
		}

		err = jpeg.Encode(f, img, &jpeg.Options{jpeg.DefaultQuality})
		if err != nil {
			panic(err)
		}

		f.Close()
	}
}
```
