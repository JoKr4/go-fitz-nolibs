//go:build !extlib

package fitz

/*
#cgo CFLAGS: -I/mingw64/include
#cgo windows,amd64 LDFLAGS: -L/mingw64/lib/ -lmupdf -lmupdf-third
*/
import "C"
