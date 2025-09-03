package main

type FreqStack struct {
	freqMap     map[int]int
	recentMap   map[int][]int
	currMaxFreq int
}

func Constructor() FreqStack {
	return FreqStack{
		freqMap:   make(map[int]int),
		recentMap: make(map[int][]int),
	}
}

func (fs *FreqStack) Push(val int) {
	fs.freqMap[val]++
	item := fs.freqMap[val]
	if item > fs.currMaxFreq {
		fs.currMaxFreq = item
	}
	if _, ok := fs.recentMap[item]; !ok {
		fs.recentMap[item] = make([]int, 0)
	}
	fs.recentMap[item] = append(fs.recentMap[item], val)

}

func (fs *FreqStack) Pop() int {
	Len := len(fs.recentMap[fs.currMaxFreq]) - 1
	x := fs.recentMap[fs.currMaxFreq][Len]
	fs.recentMap[fs.currMaxFreq] = fs.recentMap[fs.currMaxFreq][:Len]
	fs.freqMap[x]--
	if len(fs.recentMap[fs.currMaxFreq]) == 0 {
		fs.currMaxFreq--
	}

	return x
}
