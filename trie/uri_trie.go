package trie

import "strings"

type Node struct {
	isLeaf   bool
	children map[string]*Node
}

type UriTrie struct {
	Root *Node
	Size int
}

func NewTrie() UriTrie {
	return UriTrie{
		Root: &Node{
			children: make(map[string]*Node),
		},
		Size: 0,
	}
}

func (urlTrie *UriTrie) addNode(uri string) {

	paths := strings.Split(uri, "/")

	cur := urlTrie.Root
	for _, path := range paths {
		if len(path) != 0 {
			childNode, ok := cur.children[path]
			if !ok {
				childNode = &Node{
					children: make(map[string]*Node),
				}
				cur.children[path] = childNode
			}
			cur = childNode
		}
	}

	if !cur.isLeaf {
		cur.isLeaf = true
		urlTrie.Size++
	}

}

func (urlTrie *UriTrie) search(uri string) (*Node, bool) {

	paths := strings.Split(uri, "/")

	cur := urlTrie.Root
	for _, path := range paths {
		if len(path) > 0 {
			childNode, ok := cur.children[path]
			if !ok {
				return childNode, ok
			}
			cur = childNode
		}
	}
	return cur, cur.isLeaf
}
