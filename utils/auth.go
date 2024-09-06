package TCP_Chat

import "unicode"

func AuthName(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) && c != ' ' {
			return false
		}
	}
	return true
}
