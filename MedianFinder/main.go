package main

import "fmt"

func main() {
	medianFinder := Constructor()
	medianFinder.addNum(1)
	//fmt.Println(medianFinder.nums)
	medianFinder.addNum(2)
	//fmt.Println(medianFinder.nums)
	fmt.Println(medianFinder.findMedian())
	medianFinder.addNum(3)
	//fmt.Println(medianFinder.nums)
	fmt.Println(medianFinder.findMedian())
}

// Testing the implementation
// func main() {
// 	mf := Constructor()
// 	mf.AddNum(1)
// 	mf.AddNum(2)
// 	fmt.Println(mf.FindMedian()) // Output: 1.5

// 	mf.AddNum(3)
// 	fmt.Println(mf.FindMedian()) // Output: 2
// }
