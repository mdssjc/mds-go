package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)

	//var repeated strings.Builder
	//for i := 0; i < repeatCount; i++ {
	//	repeated.WriteString(character)
	//}
	//return repeated.String()
}
