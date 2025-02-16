package main

import "fmt"

func main() {
	surveyData := []int{0, 4, 2, 9, 2, 3, 5, 0}
	fmt.Print(findPeak(surveyData))
}

func findPeak(surveyData []int) []int {
	result := []int{}
	max := 0
	if len(surveyData) > 2 {

		for i := 1; i < len(surveyData); i++ {
			if (surveyData[i] > surveyData[i-1]) && (surveyData[i] > surveyData[i+1]) {
				if surveyData[i] > max {
					max = surveyData[i]
				}
				result = append(result, surveyData[i])
				i++

			}
		}
	}
	return result

}
