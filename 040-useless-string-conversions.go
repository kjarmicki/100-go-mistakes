package main

import (
	"bytes"
	"io"
	"strings"
)

/*
 * Most I/O operations in Go operate on []byte type.
 * Sometimes developers out of a habit convert []byte into a string, and then convert it back for further I/O processing.
 * This can be inefficient, harder to read and unnecessary. `bytes` package supports a lot of the same functions that `strings` does.
 * Before converting []byte to string back and forth, check if `bytes` has the things you need.
 */

func getBytes(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	_ = []byte(sanitizeString(string(b))) // not only looks gross, but is also unnecessary
	return sanitizeBytes(b), nil
}

func sanitizeString(str string) string {
	return strings.TrimSpace(str)
}

func sanitizeBytes(bts []byte) []byte {
	return bytes.TrimSpace(bts)
}
