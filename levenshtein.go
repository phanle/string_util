package string_util

import (
	. "github.com/phanle/array_util"
)

func MinimumDistance(s1 string, s2 string) int {
	if len(s2) == 0 {
		return len(s1)
	}

	if len(s1) == 0 {
		return len(s2)
	}
	var cost int;

	if s1[len(s1)-1] == s2[len(s2)-1] {
		cost = 0
	} else {
		cost = 1
	}
	return Min(MinimumDistance(s1[0:len(s1)-1], s2) + 1,
			   MinimumDistance(s1, s2[0:len(s2)-1])+ 1,
			   MinimumDistance(s1[0:len(s1)-1], s2[0:len(s2)-1]) + cost)
}
