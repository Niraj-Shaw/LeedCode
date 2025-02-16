package main

func main() {

}

func firstUniqChar(s string) int {
	myMap := make(map[rune]int)
	chars := []rune(s)
	for i := 0; i < len(chars); i++ {
		if _, ok := myMap[chars[i]]; !ok {
			myMap[chars[i]] = 1
		} else {
			myMap[chars[i]]++
		}
	}
	for i := 0; i < len(chars); i++ {
		if myMap[chars[i]] == 1 {
			return i
		}
	}
	return -1

}
