package trie

import "strings"

var colonCharacter string = ":"

type Node struct {
	isLeaf   bool
	wildcard bool
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

			// 通配符处理 简单实现
			isWildcard := isWildcard(path)
			if isWildcard {
				wildcardNode, ok := getChildWildcard(cur.children)
				if !ok {
					wildcardNode = &Node{
						wildcard: true,
						children: make(map[string]*Node),
					}
					cur.children[getCommonPath(path)] = wildcardNode
				}
				cur = wildcardNode

			} else {
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
	}

	if !cur.isLeaf {
		cur.isLeaf = true
		urlTrie.Size++
	}

}

func getCommonPath(path string) string {
	return strings.TrimSuffix(path, ":")

}

func getChildWildcard(nodes map[string]*Node) (*Node, bool) {

	for _, v := range nodes {
		if v.wildcard {
			return v, true
		}
	}
	return nil, false
}

func isWildcard(path string) bool {
	return strings.HasPrefix(path, colonCharacter)
}

func (urlTrie *UriTrie) search(uri string) (*Node, bool) {

	paths := strings.Split(uri, "/")

	cur := urlTrie.Root
	for _, path := range paths {
		if len(path) > 0 {
			childNode, ok := cur.children[path]
			if !ok {
				wildcardNode, ok := getChildWildcard(cur.children)
				if !ok {
					return wildcardNode, ok
				} else {
					cur = wildcardNode
					continue
				}
			}
			cur = childNode
		}
	}
	return cur, cur.isLeaf
}
