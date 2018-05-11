package main

import "strings"

func transliterateString(phrase string) string {
	var charMap = map[rune]rune{
		'2': 'ء',   // figure out how to do various hamzas
		'a': 0x627, // bare alif -- need more ideas on vowels
		'b': 0x628,
		'p': 0x628, // loan words or Farsi
		't': 0x62A,
		'j': 0x62C,
		'g': 0x62C, // Egyptian mostly; will have to think about ق substitution for others
		'3': 0x639,
		'm': 0x645,
		'r': 0x631,
		'7': 'ح',
		'z': 0x632,
		's': 0x633,
	}

	phrase = strings.ToLower(phrase)
	arabizi := []rune(phrase)
	var arabic []rune

	for _, letter := range arabizi {
		if val, ok := charMap[letter]; ok {
			arabic = append(arabic, val)
		} else {
			// arabic = append(arabic, letter)
		}
	}
	//reverse(arabic)
	return string(arabic)
}

func reverse(r []rune) {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}
