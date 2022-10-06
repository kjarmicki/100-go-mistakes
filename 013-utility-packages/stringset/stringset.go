package stringset

import "sort"

/*
 * (continued from ../util)
 * Way better:
 * 1. The package name is meaningful for what the package does
 * 2. Package contents are coherent
 */

type StringSet map[string]struct{}

func New(strings ...string) StringSet {
	set := make(map[string]struct{})
	for _, str := range strings {
		set[str] = struct{}{}
	}
	return set
}

func (s StringSet) Sort() []string {
	keys := make([]string, 0)
	for key, _ := range s {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
