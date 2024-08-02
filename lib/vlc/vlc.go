package vlc

import (
	"strings"
	"unicode"
)

func Encode(str string)string{
	str = prepareText(str)
	return ""
}

// prepareText prepares text to be fit in encode:
// changes upper case latters to: ! + lower case latter
// i.g.: My name is Ted -> !my name is !ted
func prepareText(str string) string{
	var buf strings.Builder
	
	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}
	
	return buf.String()
}