package router

import "strings"

type Tree struct {
	root *Node
}

type Node struct {
	prefix   string
	children []*Node
}

func NewTree() *Tree {
	return &Tree{
		root: &Node{},
	}
}

func (tree *Tree) Insert(word string) {
	tree.insertNode(tree.root, word)
}

func (tree *Tree) insertNode(node *Node, word string) {
	for _, child := range node.children {
		commonPrefix := tree.getCommonPrefix(child.prefix, word)
		if len(commonPrefix) > 0 {
			if commonPrefix == child.prefix {
				tree.insertNode(child, word[len(commonPrefix):])
			} else {
				newNode := &Node{
					prefix:   commonPrefix,
					children: []*Node{child},
				}

				child.prefix = child.prefix[len(commonPrefix):]
				node.children = append(node.children, newNode)
				node.children = tree.removeNode(node.children, child)
				tree.insertNode(newNode, word[len(commonPrefix):])
			}

			return
		}
	}

	// no common prefix, create a new child
	newNode := &Node{
		prefix: word,
	}
	node.children = append(node.children, newNode)
}

func (tree *Tree) removeNode(nodes []*Node, node *Node) []*Node {
	for i, n := range nodes {
		if n == node {
			return append(nodes[:i], nodes[i+1:]...)
		}
	}

	return nodes
}

// returns the common prefix of two strings.
func (tree *Tree) getCommonPrefix(origin, target string) string {
	minLen := len(origin)
	if len(target) < minLen {
		minLen = len(target)
	}

	for i := 0; i < minLen; i++ {
		if origin[i] != target[i] {
			return origin[:i]
		}
	}

	return origin[:minLen]
}

func (tree *Tree) Search(word string) *Node {
	return tree.searchNode(tree.root, word)
}

func (tree *Tree) searchNode(node *Node, word string) *Node {
	if len(word) == 0 {
		return node
	}

	for _, child := range node.children {
		if strings.HasPrefix(word, child.prefix) {
			return tree.searchNode(child, word[len(child.prefix):])
		}
	}

	return nil
}
