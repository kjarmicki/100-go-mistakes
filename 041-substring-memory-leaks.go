package main

import "errors"

/*
 * Substring leaks are similar to slice leaks.
 * If a substring is created from a large string, it may internally hold a reference to a large string despite it being not accessible anymore.
 */

type store struct {
}

func (s *store) leakyHandleLog(log string) error { // assume log starts with 36 char UUID and is really large afterwards
	if len(log) < 36 {
		return errors.New("incorrect log format")
	}
	uuid := log[:36] // this will still keep the reference to the entire log string in the memory
	_ = uuid
	return nil
}

func (s *store) handleLog(log string) error {
	if len(log) < 36 {
		return errors.New("incorrect log format")
	}
	uuid := string([]byte(log[:36])) // this makes a copy of the substring, so internal byte slice references a new backing array
	_ = uuid
	return nil
}
