package strDiff

import (
	"unicode"
)

func IsHiragana(s string) bool {
	for _, r := range s {
		if ok := unicode.In(r, unicode.Hiragana); !ok {
			return false
		}
	}
	return true
}
