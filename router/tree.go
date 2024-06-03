/*
*
TODO:
- have to evaluate when it starts with '/:' such case is when there is only one first path
- have to be able to resolve /:var1/:var2 current behaviour stack overflow
*
*/
package router

import "strings"

type Tree[T any] struct {
	root *Node[T]
}

type Node[T any] struct {
	prefix   string
	children []*Node[T]
}

func NewTree[T any]() *Tree[T] {
	return &Tree[T]{
		root: &Node[T]{},
	}
}

func (tree *Tree[T]) Insert(word string) {
	tree.insertNode(tree.root, word)
}

func (tree *Tree[T]) insertNode(node *Node[T], word string) {
	for _, child := range node.children {
		commonPrefix := tree.getCommonPrefix(child.prefix, word)
		if len(commonPrefix) > 0 {
			if commonPrefix == child.prefix {
				tree.insertNode(child, word[len(commonPrefix):])
			} else {
				newNode := &Node[T]{
					prefix:   commonPrefix,
					children: []*Node[T]{child},
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
	newNode := &Node[T]{
		prefix: word,
	}
	node.children = append(node.children, newNode)
}

func (tree *Tree[T]) removeNode(nodes []*Node[T], node *Node[T]) []*Node[T] {
	for i, n := range nodes {
		if n == node {
			return append(nodes[:i], nodes[i+1:]...)
		}
	}

	return nodes
}

// returns the common prefix of two strings.
func (tree *Tree[T]) getCommonPrefix(origin, target string) string {
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

func (tree *Tree[T]) Search(word string) *Node[T] {
	return tree.searchNode(tree.root, word)
}

func (tree *Tree[T]) searchNode(node *Node[T], word string) *Node[T] {
	if len(word) == 0 {
		return node
	}

	for _, child := range node.children {
		if strings.HasPrefix(word, child.prefix) {
			return tree.searchNode(child, word[len(child.prefix):])
		}

		if strings.HasPrefix(child.prefix, ":") {
			nextPath := strings.Index(word, "/")
			if nextPath == -1 {
				return child
			} else {
				prefix := ""
				val := strings.Index(child.prefix, "/")
				if val == -1 {
					prefix = child.prefix
				} else {
					prefix = child.prefix[:val]
				}

				return tree.searchNode(node, prefix+word[nextPath:])
			}
		}
	}

	return nil
}
