package util

import "sort"

/*
 * General utility packages are suboptimal because their name is meaningless and they lack cohesion.
 * Try to find patterns within these packages and extract their parts info more meaningful packages.
 * (see ../stringset)
 */

func NewStringSet(strings ...string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, str := range strings {
		set[str] = struct{}{}
	}
	return set
}

func SortStringSet(set map[string]struct{}) []string {
	keys := make([]string, 0)
	for key, _ := range set {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
