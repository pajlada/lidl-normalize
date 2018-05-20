package normalize

import (
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func stripAccents(s string) (string, error) {
	b := make([]byte, len(s))

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	nDst, _, err := t.Transform(b, []byte(s), true)
	if err != nil {
		return "", err
	}

	return string(b[:nDst]), nil
}

// Normalize hehe
func Normalize(s string) (string, error) {
	if s == "" {
		return s, nil
	}

	for i, w := 0, 0; i < len(s); i += w {
		char, width := utf8.DecodeRuneInString(s[i:])
		replacement, exists := confusableTable[char]
		if exists {
			xd := string(replacement)
			s = s[:i] + xd + s[i+width:]
			w = len(xd)
		} else {
			w = width
		}
	}

	var err error
	s, err = stripAccents(s)
	if err != nil {
		return "", err
	}

	return s, nil
}
