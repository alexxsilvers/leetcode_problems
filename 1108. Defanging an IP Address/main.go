package main

import (
	"log"
	"strings"
)

// Given a valid (IPv4) IP address, return a defanged version of that IP address.
//
// A defanged IP address replaces every period "." with "[.]".
func main() {
	log.Println(defangIPaddr("1.1.1.1"))      // 1[.]1[.]1[.]1
	log.Println(defangIPaddr("255.100.50.0")) // 255[.]100[.]50[.]0
}

func defangIPaddr(address string) string {
	sb := strings.Builder{}
	sb.Grow(len(address) + 12)
	for _, c := range []byte(address) {
		if c == '.' {
			sb.WriteString("[.]")
		} else {
			sb.WriteByte(c)
		}
	}

	return sb.String()
}
