package functions

import "unicode"

func inPrintable(Message string) bool {
	for _,Rune := range Message {
		if !unicode.IsPrint(Rune) {
			return false
		}
	}
	return true
}