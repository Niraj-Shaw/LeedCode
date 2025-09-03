package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}

func numUniqueEmails(emails []string) int {
	emailsMap := make(map[string]int)

	for _, str := range emails {
		processedStr := processString(str)
		emailsMap[processedStr]++
	}

	return len(emailsMap)
}

func processString(str string) string {
	sepIndex := strings.Index(str, "@")
	endStr := str[sepIndex:]
	tempstr := strings.ToLower(str[:sepIndex])

	tempstr = strings.ReplaceAll(tempstr, ".", "")
	plusIndex := strings.Index(tempstr, "+")
	if plusIndex != -1 {
		tempstr = tempstr[:plusIndex]
	}
	tempstr += endStr

	return tempstr

}
