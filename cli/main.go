package main

import (
	"fmt"
	"strings"

	normalize "github.com/pajlada/lidl-normalize"
)

var tests = map[string]string{
	"a":             "a",
	"Æ":             "AE",
	"æ":             "ae",
	"Å":             "A",
	"å":             "a",
	"Ǻ":             "A",
	"ǻ":             "a",
	"Ḁ":             "A",
	"ḁ":             "a",
	"ẚ":             "a",
	"Ă":             "A",
	"ă":             "a",
	"Ặ":             "A",
	"ặ":             "a",
	"Ắ":             "A",
	"ắ":             "a",
	"Ằ":             "A",
	"ằ":             "a",
	"Ẳ":             "A",
	"ẳ":             "a",
	"Ẵ":             "A",
	"ẵ":             "a",
	"Ȃ":             "A",
	"ȃ":             "a",
	"Â":             "A",
	"â":             "a",
	"Ậ":             "A",
	"ậ":             "a",
	"Ấ":             "A",
	"ấ":             "a",
	"Ầ":             "A",
	"ầ":             "a",
	"Ẫ":             "A",
	"ẫ":             "a",
	"Ẩ":             "A",
	"ẩ":             "a",
	"Ả":             "A",
	"ả":             "a",
	"Ǎ":             "A",
	"ǎ":             "a",
	"Ⱥ":             "A",
	"ⱥ":             "a",
	"Ȧ":             "A",
	"ȧ":             "a",
	"Ǡ":             "A",
	"ǡ":             "a",
	"Ạ":             "A",
	"ạ":             "a",
	"Ä":             "A",
	"ä":             "a",
	"Ǟ":             "A",
	"ǟ":             "a",
	"À":             "A",
	"à":             "a",
	"Ȁ":             "A",
	"ȁ":             "a",
	"Á":             "A",
	"á":             "a",
	"Ā":             "A",
	"ā":             "a",
	"Ā̀":            "A",
	"ā̀":            "a",
	"Ã":             "A",
	"ã":             "a",
	"Ą":             "A",
	"ą":             "a",
	"Ą́":            "A",
	"ą́":            "a",
	"Ą̃":            "A",
	"ą̃":            "a",
	"ᶏ":             "a",
	"Ɑ":             "a",
	"ɑ":             "a",
	"ᶐ":             "a",
	"Ɐ":             "A",
	"ɐ":             "a",
	"Λ":             "A",
	"ʌ":             "A",
	"Ɒ":             "a",
	"ɒ":             "a",
	"ᶛ":             "a",
	"ᴀ":             "A",
	"ᴬ":             "A",
	"ᵃ":             "a",
	"ᵄ":             "a",
	"ₐ":             "a",
	"asd ꬱsd":       "asd asd",
	"asd ⍺ xd":      "asd a xd",
	"abc å def":     "abc a def",
	"ˢᵐᵒˡ ⁿᵃᵗᶦᵒⁿ":   "smol nation", // smol nation
	"Ниг":           "Nig",
	"🇦":             "A",
	"🇺🇦XD":          "UAXD",
	"🆓 ICE":         "FREE ICE",
	"ер":            "er",
	"chocolate 🇳🇮b": "chocolate NIb",
	"🅱lueberry":     "Blueberry",
	"⒝":             "b",
}

var itests = map[string]string{
	"⒜ ⒝ ⒞ ⒟ ⒠ ⒡ ⒢ ⒣ ⒤ ⒥ ⒦ ⒧ ⒨ ⒩ ⒪ ⒫ ⒬ ⒭ ⒮ ⒯ ⒰ ⒱ ⒲ ⒳ ⒴": "a b c d e f g h i j k l m n o p q r s t u v w x y",
}

func main() {
	i := 0
	for in, result := range tests {
		out, err := normalize.Normalize(in)
		if err != nil {
			panic(err)
		}
		if out != result {
			fmt.Printf("(%s) failed - '%s' != '%s' %s:%s\n", in, out, result, in, result)
		}
		i++
	}

	for in, result := range itests {
		out, err := normalize.Normalize(in)
		if err != nil {
			panic(err)
		}

		out = strings.ToLower(out)
		result = strings.ToLower(result)
		if out != result {
			fmt.Printf("(%s) failed - '%s' != '%s' %s:%s\n", in, out, result, in, result)
		}
		i++
	}
}
