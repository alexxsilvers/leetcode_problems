package ctci

func isUniqueChars(s string) bool {
	charsMap := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if charsMap[s[i]] == true {
			return false
		}
		charsMap[s[i]] = true
	}
	return true
}
