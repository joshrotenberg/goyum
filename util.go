package goyum

import (
	"unicode"
)

func downcase(target string) string {
	a := []rune(target)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}
