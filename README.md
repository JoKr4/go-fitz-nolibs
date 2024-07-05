## go-fitz-nolibs

Go wrapper for [MuPDF](http://mupdf.com/) fitz library that can extract pages from PDF, EPUB and MOBI documents as images, text, html or svg.

### Build

TODO msys2, mingw etc
    
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
