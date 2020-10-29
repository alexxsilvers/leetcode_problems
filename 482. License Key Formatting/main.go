package main

import "log"

// You are given a license key represented as a string S which consists only alphanumeric
// character and dashes. The string is separated into N+1 groups by N dashes.
// Given a number K, we would want to reformat the strings such that each group
// contains exactly K characters, except for the first group which could be shorter
// than K, but still must contain at least one character. Furthermore, there must be a
// dash inserted between two groups and all lowercase letters should be converted to uppercase.
// Given a non-empty string S and a number K, format the string according to the rules described above.
func main() {
	log.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4)) // 5F3Z-2E9W
	log.Println(licenseKeyFormatting("2-5g-3-J", 2)) // 2-5G-3J
}

func licenseKeyFormatting(S string, K int) string {
	upper := map[byte]byte{
		'a': 'A', 'b': 'B', 'c': 'C', 'd': 'D', 'e': 'E', 'f': 'F', 'g': 'G', 'h': 'H', 'i': 'I', 'j': 'J', 'k': 'K',
		'l': 'L', 'm': 'M', 'n': 'N', 'o': 'O', 'p': 'P', 'q': 'Q', 'r': 'R', 's': 'S', 't': 'T', 'u': 'U', 'v': 'V',
		'w': 'W', 'x': 'X', 'y': 'Y', 'z': 'Z',
	}
	t := K
	tempString := make([]byte, 0, len(S))

	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '-' {
			continue
		}

		if t == 0 {
			tempString = append(tempString, '-')
			t = K
		}

		if S[i] >= 'a' && S[i] <= 'z' {
			tempString = append(tempString, upper[S[i]])
		} else {
			tempString = append(tempString, S[i])
		}
		t--
	}

	for i := 0; i < len(tempString)/2; i++ {
		tempString[i], tempString[len(tempString)-i-1] = tempString[len(tempString)-i-1], tempString[i]
	}

	return string(tempString)
}
