package main

import (
	"fmt"
)

type Tree struct {
	node *TreeNode
}

func (t *Tree) insert(key int) *Tree  {
	if t.node == nil {
		t.node = &TreeNode {
			key: key,
		}
	} else {
		t.node.insert(key)
	}

	return t
}

type TreeNode struct {
	key int
	left *TreeNode
	right *TreeNode
}

func (node *TreeNode) insert(key int)  {
	if key < node.key {
			if node.left == nil {
				node.left = &TreeNode {
					key: key,
				}
			} else {
				node.left.insert(key)
			}
	} else {
			if node.right == nil {
				node.right = &TreeNode {
					key: key,
				}
			} else {
				node.right.insert(key)				
			}
	}
}

func printTree(node *TreeNode)  {
	 if node == nil {
		 return
	 }

	 fmt.Println(node.key)
	 printTree(node.left)
	 printTree(node.right)
}

func (node *TreeNode) exists(val int) bool  {
	if node == nil {
		return false
	}

	if node.key == val {
		return true
	}

	if node.key > val {
		return node.left.exists(val)
	} else {
		return node.right.exists(val)
	}
}

func BSTOperations() {
	tree := &Tree{}
	tree.insert(5)
	tree.insert(2) 
	tree.insert(21)
	tree.insert(-4)
	tree.insert(3) 
	tree.insert(19)
	tree.insert(25)

	printTree(tree.node)

	fmt.Println(tree.node.exists(23))
	fmt.Println(tree.node.exists(-4))
	fmt.Println(tree.node.exists(3))
}