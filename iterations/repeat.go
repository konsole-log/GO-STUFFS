package iterations

import "strings"

func Repeat(character string, total int) string {
	var repeated strings.Builder
	for i := 0; i < total; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
