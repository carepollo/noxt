package router

import (
	"regexp"
	"strings"
)

type Tree[T any] struct {
	root *Node[T]
}

type Node[T any] struct {
	prefix   string
	children []*Node[T]
	value    T
}

func NewTree[T any]() *Tree[T] {
	return &Tree[T]{
		root: &Node[T]{},
	}
}

func (tree *Tree[T]) Insert(word string, value T) {
	tree.insertNode(tree.root, word, value)
}

func (tree *Tree[T]) insertNode(node *Node[T], word string, value T) {
	for _, child := range node.children {
		commonPrefix := tree.getCommonPrefix(child.prefix, word)
		if len(commonPrefix) > 0 {
			if commonPrefix == child.prefix {
				tree.insertNode(child, word[len(commonPrefix):], value)
			} else {
				newNode := &Node[T]{
					prefix:   commonPrefix,
					children: []*Node[T]{child},
					value:    value,
				}

				child.prefix = child.prefix[len(commonPrefix):]
				node.children = append(node.children, newNode)
				node.children = tree.removeNode(node.children, child)
				tree.insertNode(newNode, word[len(commonPrefix):], value)
			}

			return
		}
	}

	// no common prefix, create a new child
	newNode := &Node[T]{
		prefix: word,
		value:  value,
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
		// static path
		if strings.HasPrefix(word, child.prefix) {
			return tree.searchNode(child, word[len(child.prefix):])
		}

		// dynamic path
		if strings.Contains(child.prefix, ":") {
			pattern, err := regexp.Compile("^[a-zA-Z0-9]+$")
			if err != nil {
				panic("invalid regexp")
			}

			varStart, varEnd := 0, 0
			replace := word
			for i := 0; i < len(child.prefix); i++ {
				letter := child.prefix[i]
				if letter == ':' {
					varStart = i
					continue
				}

				if !pattern.MatchString(string(letter)) || i == len(child.prefix)-1 || i == '/' {
					varEnd = strings.Index(word, string(letter))
					if varEnd < 0 {
						return nil
					}

					replace = strings.Replace(word, replace[varStart:varEnd], child.prefix[varStart:i], 1)
					varStart, varEnd = 0, 0
				}
			}

			return tree.searchNode(child, replace)
		}
	}

	return nil
}

func (node *Node[T]) GetValue() T {
	return node.value
}
