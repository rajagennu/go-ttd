package iteration

import "strings"

func Repeat(char string, count int) string {
	if char == "" {
		return ""
	}
	repeatedWord := ""
	for i := 0; i < count; i++ {
		repeatedWord += char
	}
	return repeatedWord
}

func RepeatV2(char string, count int) string {
	return strings.Repeat(char, count)
}
