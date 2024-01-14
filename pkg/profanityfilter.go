package pkg

import "strings"

// check the message if it contain the profanity or not

func NoProfanityInMsg(words []string) bool {
	count := 0
	for _, word := range words {
		if IsProfanity(word) {
			count++
		}
	}
	return count == 0
}

func IsProfanity(str string) bool {
	words := []string{"fuck", "bitch", "slut", "cunt", "twat", "bastard", "whore"}
	str2 := strings.ToLower(str)
	for _, word := range words {
		if str2 == word {
			return true
		}
	}
	return false
}
