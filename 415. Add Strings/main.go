package main

import (
	"log"
)

/*
	Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.
	The length of both num1 and num2 is < 5100.
	Both num1 and num2 contains only digits 0-9.
	Both num1 and num2 does not contain any leading zero.
	You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/
func main() {
	log.Println(addStrings("1", "999"))       // 1000
	log.Println(addStrings("99999", "11111")) // 111110
	log.Println(addStrings("1", "9"))         // 10
	log.Println(addStrings("3", "5"))         // 8
}

func addStrings(num1 string, num2 string) string {
	var add byte

	l1, l2 := len(num1)-1, len(num2)-1
	if l1 > l2 {
		l1, l2 = l2, l1
		num1, num2 = num2, num1
	}

	res := make([]byte, len(num2)+1)
	k := l2 + 1

	for l1 >= 0 && l2 >= 0 {
		n1 := num1[l1] - '0'
		l1--

		n2 := num2[l2] - '0'
		l2--

		sum := n1 + n2 + add

		add = sum / 10
		res[k] = sum%10 + '0'
		k--
	}

	for l2 >= 0 {
		sum := num2[l2] - 48 + add
		add = sum / 10
		res[k] = sum%10 + '0'
		l2--
		k--
	}

	if add > 0 {
		res[0] = '1'
		return string(res)
	}

	return string(res[1:])
}
