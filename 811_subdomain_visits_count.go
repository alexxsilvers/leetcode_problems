package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	domainCounts := make(map[string]int)
	for _, addr := range cpdomains {
		spl := strings.Split(addr, " ")
		cnt, _ := strconv.Atoi(spl[0])
		domain := spl[1]
		for len(domain) > 0 {
			domainCounts[domain] += cnt
			nextIndex := strings.Index(domain, ".")
			if nextIndex == -1 {
				break
			}

			domain = domain[nextIndex + 1:]
		}
	}

	result := make([]string, 0, len(domainCounts))
	for domain, counter := range domainCounts {
		result = append(result, strconv.Itoa(counter) + " " + domain)
	}
	return result
}

func main() {
	fmt.Println("result ", subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}