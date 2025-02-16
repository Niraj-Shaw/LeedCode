package main

import (
	"fmt"
	"sort"
)

type Node struct {
	v int
	r int
	c int
}

type Coord struct {
	x   int
	y   int
	val int
}

type pq []*Node

func (h pq) Len() int            { return len(h) }
func (h pq) Less(i, j int) bool  { return h[i].v < h[j].v }
func (h pq) Swap(i, j int)       { h[i].v, h[j].v = h[j].v, h[i].v }
func (h *pq) Push(x interface{}) { *h = append(*h, x.(*Node)) }
func (h *pq) Pop() interface{} {
	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	forest := [][]int{
		{2, 3, 4}, {0, 0, 5}, {8, 7, 6},
	}

	fmt.Print(cutOffTree(forest))
}

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func cutOffTree(forest [][]int) int {
	// nodes := []*Node{}
	// R := len(forest)
	// C := len(forest[0])
	// for rowIndex, rows := range forest {
	// 	for colIndex, val := range rows {
	// 		if val > 1 {
	// 			nodes = append(nodes, &Node{v: val, r: rowIndex, c: colIndex})
	// 		}
	// 	}
	// }
	// sort.Slice(nodes, func(i, j int) bool {
	// 	return nodes[i].v < nodes[j].v
	// })

	// bfs := func(s, t *Node) int {
	// 	visited := make(map[[2]int]bool)
	// 	q := []Node{*s}
	// 	visited[[2]int{s.r, s.c}] = true
	// 	distance := 0

	// 	for len(q) > 0 {
	// 		size := len(q)
	// 		for i := 0; i < size; i++ {
	// 			curr := q[0]
	// 			q = q[1:]
	// 			if curr.r == t.r && curr.c == t.c {
	// 				return distance
	// 			}
	// 			for _, path := range directions {
	// 				x := curr.r + path[0]
	// 				y := curr.c + path[1]
	// 				if x >= 0 && x < R && y >= 0 && y < C && forest[x][y] != 0 && !visited[[2]int{x, y}] {
	// 					q = append(q, Node{v: forest[x][y], r: x, c: y})
	// 					visited[[2]int{x, y}] = true
	// 				}
	// 			}
	// 		}
	// 		distance++
	// 	}
	// 	return -1
	// }

	// ans := 0
	// start := &Node{v: forest[0][0], r: 0, c: 0}
	// for _, node := range nodes {
	// 	res := bfs(start, node)
	// 	if res == -1 {
	// 		return -1
	// 	}
	// 	ans += res
	// 	start = node
	// }
	// return ans
	//----
	// R := len(forest)
	// C := len(forest[0])

	// trees := []Coord{}

	// for i := range forest {
	// 	for j := range forest[i] {
	// 		if forest[i][j] > 1 {
	// 			trees = append(trees, Coord{i, j, forest[i][j]})
	// 		}
	// 	}
	// }

	// sort.Slice(trees, func(i, j int) bool {
	// 	return trees[i].val < trees[j].val
	// })

	// paths := [][]int{
	// 	{1, 0},
	// 	{0, 1},
	// 	{0, -1},
	// 	{-1, 0},
	// }

	// bfs := func(sx, sy, dx, dy int) int {
	// 	row := R
	// 	col := C
	// 	visited := make([][]bool, row)
	// 	rows := make([]bool, row*col)
	// 	for i := range visited {
	// 		visited[i] = rows[i*C : (i+1)*C]
	// 	}

	// 	q := [][3]int{}
	// 	q = append(q, [3]int{sx, sy, 0})
	// 	visited[sx][sy] = true

	// 	for len(q) > 0 {
	// 		curr := q[0]
	// 		q = q[1:]

	// 		if curr[0] == dx && curr[1] == dy {
	// 			return curr[2]
	// 		}

	// 		for _, p := range paths {
	// 			nx := p[0] + curr[0]
	// 			ny := p[1] + curr[1]

	// 			if nx >= 0 && nx < R && ny >= 0 && ny < C && !visited[nx][ny] && forest[nx][ny] != 0 {
	// 				visited[nx][ny] = true
	// 				q = append(q, [3]int{nx, ny, curr[2] + 1})
	// 			}
	// 		}
	// 	}
	// 	return -1
	// }

	// count := 0
	// start := Coord{x: 0, y: 0}
	// for i := 0; i < len(trees); i++ {
	// 	dist := bfs(start.x, start.y, trees[i].x, trees[i].y)
	// 	if dist == -1 {
	// 		return -1
	// 	}
	// 	start.x = trees[i].x
	// 	start.y = trees[i].y
	// 	count += dist
	// }
	// return count
	Nodes := []Node{}
	for row, trees := range forest {
		for col, height := range trees {
			if height > 1 {
				Nodes = append(Nodes, Node{v: height, r: row, c: col})
			}
		}
	}
	sort.Slice(Nodes, func(i, j int) bool {
		return Nodes[i].v < Nodes[j].v
	})
	R := len(forest)
	C := len(forest[0])

	bfs := func(sr, sc, tr, tc int) int {
		visited := make([][]bool, R)
		for i := range visited {
			visited[i] = make([]bool, C)
		}
		q := [][3]int{{0, sr, sc}}
		for len(q) > 0 {
			curr := q[0]
			q = q[1:]
			visited[curr[1]][curr[2]] = true
			if curr[1] == tr && curr[2] == tc {
				return curr[0]
			}
			for _, direction := range directions {
				nr, nc := curr[1]+direction[0], curr[2]+direction[1]
				if nr >= 0 && nr < R && nc >= 0 && nc < C && !visited[nr][nc] && forest[nr][nc] > 1 {
					q = append(q, [3]int{curr[0] + 1, nr, nc})
				}
			}

		}
		return -1

	}
	ans := 0
	sr, sc := 0, 0
	for i := 0; i < len(Nodes); i++ {
		res := bfs(sr, sc, Nodes[i].r, Nodes[i].c)
		if res == -1 {
			return -1
		}
		ans += res
		sr = Nodes[i].r
		sc = Nodes[i].c
	}

	return ans

}
