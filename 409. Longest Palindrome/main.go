package main

import "log"

// Given a string which consists of lowercase or uppercase letters, find
// the length of the longest palindromes that can be built with those letters.
// This is case sensitive, for example "Aa" is not considered a palindrome here.
func main() {
	log.Println(longestPalindrome("bb"))             // 2
	log.Println(longestPalindrome("aaabbbbbccccdd")) // 13 = "dccbbaaabbccd"
	log.Println(longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}

func longestPalindrome(s string) int {
	memo := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		memo[s[i]]++
	}

	odds, ret := 0, 0
	for _, cnt := range memo {
		if cnt%2 == 1 {
			odds++
		}

		ret += cnt
	}

	if odds > 0 {
		ret -= odds - 1
	}
	return ret
}
