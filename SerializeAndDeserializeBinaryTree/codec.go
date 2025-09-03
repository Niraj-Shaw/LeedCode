package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	// if root == nil {
	// 	return ""
	// }
	// serialArray := []string{}
	// q := []*TreeNode{root}
	// for len(q) != 0 {
	// 	item := q[0]
	// 	q = q[1:]
	// 	if item != nil {
	// 		q = append(q, item.left)
	// 		q = append(q, item.right)
	// 		serialArray = append(serialArray, strconv.Itoa(item.val))
	// 	} else {
	// 		serialArray = append(serialArray, "null")
	// 	}
	// }
	// i := len(serialArray) - 1
	// for ; i > 0; i-- {
	// 	if serialArray[i] != "null" {
	// 		break
	// 	}
	// }
	// serialArray = serialArray[0 : i+1]

	// return strings.Join(serialArray, " ")
	var f func(root *TreeNode) string
	f = func(root *TreeNode) string {
		if root == nil {
			return "Null,"
		}
		str := strconv.Itoa(root.val) + ","
		str += f(root.left)
		str += f(root.right)
		return str
	}
	return f(root)

}

func (this *Codec) deserialize(data string) *TreeNode {
	// q := []*TreeNode{}
	// root := &TreeNode{}
	// index := 1

	// if len(data) == 0 {
	// 	return nil
	// }

	// nodes := strings.Split(data, " ")

	// if nodes[0] == "null" {
	// 	return nil
	// }

	// nodeVal, err := strconv.Atoi(nodes[0])
	// if err == nil {
	// 	root.val = nodeVal
	// }
	// q = append(q, root)

	// for len(q) != 0 {
	// 	node := q[0]
	// 	q = q[1:]
	// 	if index < len(nodes) {
	// 		if nodes[index] == "null" {
	// 			node.left = nil
	// 		} else {
	// 			nodeVal, err := strconv.Atoi(nodes[index])
	// 			if err == nil {
	// 				newNode := &TreeNode{}
	// 				newNode.val = nodeVal
	// 				node.left = newNode
	// 				q = append(q, newNode)
	// 			}
	// 		}
	// 		index++
	// 	}
	// 	if index < len(nodes) {
	// 		if nodes[index] == "null" {
	// 			node.right = nil
	// 		} else {
	// 			nodeVal, err := strconv.Atoi(nodes[index])
	// 			if err == nil {
	// 				newNode := &TreeNode{}
	// 				newNode.val = nodeVal
	// 				node.right = newNode
	// 				q = append(q, newNode)
	// 			}
	// 		}
	// 		index++
	// 	}
	// }

	// return root
	dataArray := strings.Split(data, ",")
	var f func(*[]string) *TreeNode

	f = func(s *[]string) *TreeNode {
		if len(*s) == 0 {
			return nil
		}
		val := (*s)[0]
		*s = (*s)[1:] // Advance the slice
		if val == "Null" {
			return nil
		}
		v, _ := strconv.Atoi(val)
		root := &TreeNode{val: v}
		root.left = f(s)
		root.right = f(s)
		return root
	}

	return f(&dataArray)
}
