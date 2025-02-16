package main

import "fmt"

func main() {
	fmt.Print(canConstruct("rat", "car"))
}

func canConstruct(ransomNote string, magazine string) bool {
	mag := []rune(magazine)
	myMap := make(map[rune]int)
	for i := 0; i < len(mag); i++ {
		if _, ok := myMap[mag[i]]; !ok {
			myMap[mag[i]] = 1

		} else {
			myMap[mag[i]]++
		}
	}
	for _, val := range ransomNote {
		if count, ok := myMap[val]; ok {
			if count > 0 {
				myMap[val]--
			} else {
				return false
			}

		} else {
			return false
		}

	}
	return true

}
