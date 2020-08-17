package util

import "strconv"

func MapContainsKey(m map[string]string, s string) bool {
	if _, ok := m[s]; ok {
		return true
	}
	return false
}

func IsNumeric(s string) bool  {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
