package ctci

func isPermutation(first, second string) bool {
	if len(first) != len(second) {
		return false
	}

	firstChars := make(map[byte]int)
	for i := 0; i < len(first); i++ {
		firstChars[first[i]]++
	}

	for i := 0; i < len(second); i++ {
		if _, exist := firstChars[second[i]]; exist {
			firstChars[second[i]]--
		}

		if firstChars[second[i]] == 0 {
			delete(firstChars, second[i])
		}
	}

	return len(firstChars) == 0
}
