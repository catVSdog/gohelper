package str

import "unicode"

func IsUpper(a string) bool {
	area := map[string]struct{}{
		"A": {}, "B": {}, "C": {}, "D": {}, "E": {}, "F": {}, "G": {}, "H": {}, "I": {}, "J": {}, "K": {}, "L": {},
		"M": {}, "N": {}, "O": {}, "P": {}, "Q": {}, "R": {}, "S": {}, "T": {}, "U": {}, "V": {}, "W": {}, "X": {},
		"Y": {}, "Z": {},
	}
	_, ok := area[a]
	return ok
}

func IsLower(a string) bool {
	area := map[string]struct{}{
		"a": {}, "b": {}, "c": {}, "d": {}, "e": {}, "f": {}, "g": {}, "h": {}, "i": {}, "j": {}, "k": {}, "l": {},
		"m": {}, "n": {}, "o": {}, "p": {}, "q": {}, "r": {}, "s": {}, "t": {}, "u": {}, "v": {}, "w": {}, "x": {},
		"y": {}, "z": {},
	}
	_, ok := area[a]
	return ok
}

func IsAlphanumeric(a string) bool {
	aRune := []rune(a)
	for _, v := range aRune {
		if !unicode.IsLetter(v) && !unicode.IsNumber(v) {
			return false
		}
	}
	return true
}
