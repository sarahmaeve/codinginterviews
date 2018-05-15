package main

import "testing"

func TestTransliterate(t *testing.T) {
	wordlist := map[string]string{
		// arabizi, Arabic
		"6alib":   "طالب",
		"mar7aba": "مرحبا", // perhaps additonal spellings and diacritic
		"a7mar":   "أحمر",
	}

	for arabizi, arabic := range wordlist {
		transl := transliterateString(arabizi)
		if transl != arabic {
			t.Errorf("Expected %s for %s, but received %s", arabic, arabizi, transl)
		}
	}
	return
}
