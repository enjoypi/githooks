package aid

import "strings"

type StringSet map[string]bool

func SliceToSet(stringSlice []string, trim bool, lowerCase bool) StringSet {
	set := make(map[string]bool)
	for _, ss := range stringSlice {
		if trim {
			ss = strings.TrimSpace(ss)
		}
		if lowerCase {
			ss = strings.ToLower(ss)
		}
		if len(ss) > 0 {
			set[ss] = true
		}
	}
	return set
}
