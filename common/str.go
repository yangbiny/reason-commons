package common

import "unicode"

func IsEmptyStr(str *string) bool {
	return str == nil || len(*str) == 0
}

func IsBlankStr(str *string) bool {
	if IsEmptyStr(str) {
		return true
	}
	for _, char := range *str {
		if !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}

func IsNotBlankStr(str *string) bool {
	return !IsBlankStr(str)
}
