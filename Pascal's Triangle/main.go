package main

func main() {
	generate(5)
}

func generate(numRows int) [][]int {
	result := [][]int{}
	tempArr := []int{1}
	for i := 0; i < numRows; i++ {
		currArr := []int{}
		for j := 0; j <= i; j++ {

			if j == 0 {
				currArr = append(currArr, tempArr[j])
			} else if j == len(tempArr) {
				currArr = append(currArr, tempArr[j-1])
			} else {
				currArr = append(currArr, tempArr[j-1]+tempArr[j])
			}
		}
		result = append(result, currArr)
		tempArr = currArr

	}
	return result
}
