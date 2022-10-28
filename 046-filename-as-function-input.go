package main

import (
	"bufio"
	"io"
	"os"
)

// not a good design, hard to test
// need to create an actual file per test case
func countEmptyLinesInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer func() {
		file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return 1, nil
	}
	return 1, nil
}

// better - can be repurposed for a different data source than a file, also easier and faster to test
func countEmptyLines(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		return 1, nil
	}
	return 1, nil
}
