package main

import (
	"container/list"
	"fmt"
)

func main() {
	//fmt.Print(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	//fmt.Print(findOrder(2, [][]int{{0, 1}}))
	fmt.Print(findOrder(2, [][]int{{0, 1}, {1, 0}}))
}

// //dfs
// func findOrder(numCourses int, prerequisites [][]int) []int {

// 	if len(prerequisites) == 0 {
// 		result := []int{}
// 		for i := 0; i < numCourses; i++ {
// 			result = append(result, i)
// 		}
// 		return result
// 	}

// 	result := []int{}
// 	edges := make(map[int][]int)
// 	seen := make(map[int]bool)
// 	done := make(map[int]bool)

// 	for i := 0; i < numCourses; i++ {
// 		edges[i] = make([]int, 0)
// 	}

// 	for _, courses := range prerequisites {
// 		edges[courses[0]] = append(edges[courses[0]], courses[1])
// 	}

// 	var dfs func(val int) bool

// 	dfs = func(val int) bool {
// 		if done[val] {
// 			return true
// 		}

// 		if seen[val] {
// 			return false
// 		}
// 		seen[val] = true

// 		for _, course := range edges[val] {
// 			if !dfs(course) {
// 				return false
// 			}
// 		}

// 		seen[val] = false
// 		done[val] = true
// 		result = append(result, val)
// 		return true
// 	}

// 	for key, nodes := range edges {
// 		for _, course := range nodes {
// 			if !dfs(course) {
// 				return []int{}
// 			}

// 		}
// 		if !done[key] {
// 			done[key] = true
// 			result = append(result, key)
// 		}
// 	}

// 	return result

// }

func findOrder(numCourses int, prerequisites [][]int) []int {
	if len(prerequisites) == 0 {
		res := []int{}
		for i := range numCourses {
			res = append(res, i)
		}
		return res
	}

	edges := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	q := list.New()
	res := []int{}

	for _, courses := range prerequisites {
		edges[courses[1]] = append(edges[courses[1]], courses[0])
		inDegree[courses[0]]++
	}

	for key, val := range inDegree {
		if val == 0 {
			q.PushBack(key)
		}
	}

	for q.Len() != 0 {
		element := q.Front()
		for _, val := range edges[element.Value.(int)] {
			inDegree[val]--
			if inDegree[val] == 0 {
				q.PushBack(val)
			}
		}
		res = append(res, element.Value.(int))
		q.Remove(element)

	}

	if len(res) == numCourses {
		return res
	}

	return []int{}
}
