package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func main() {

}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	seen := make(map[*Node]*Node)
	return dfs(node, seen)
}

func dfs(node *Node, seen map[*Node]*Node) *Node {
	if node == nil {
		return node
	}

	if v, ok := seen[node]; ok {
		return v
	}

	clonedNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}

	seen[node] = clonedNode

	for i, neighbor := range node.Neighbors {
		clonedNode.Neighbors[i] = dfs(neighbor, seen)
	}

	return clonedNode

}
