package main

import (
	"io"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
)

/*
 * iotest helps with testing Reader/Writer implementations
 * https://pkg.go.dev/testing/iotest
 */

type LowerCaseReader struct {
	reader io.Reader
}

func TestLowerCaseReader(t *testing.T) {
	err := iotest.TestReader(
		strings.NewReader("abcd"), // any custom reader implementation
		[]byte("abcd"),
	)
	assert.NoError(t, err)
}
