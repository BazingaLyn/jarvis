package trie

import "strings"

var colonCharacter string = ":"

type Node struct {
	isLeaf   bool
	wildcard bool
	params   map[string]string
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
				wildcardNode, _, ok := getChildWildcard(cur.children)
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
	return strings.TrimPrefix(path, ":")

}

func getChildWildcard(nodes map[string]*Node) (*Node, string, bool) {

	for k, v := range nodes {
		if v.wildcard {
			return v, k, true
		}
	}
	return nil, "", false
}

func isWildcard(path string) bool {
	return strings.HasPrefix(path, colonCharacter)
}

func (urlTrie *UriTrie) search(uri string) (*Node, bool) {

	paths := strings.Split(uri, "/")

	cur := urlTrie.Root
	params := make(map[string]string)

	for _, path := range paths {
		if len(path) > 0 {
			childNode, ok := cur.children[path]
			if !ok {
				wildcardNode, paramName, ok := getChildWildcard(cur.children)
				if !ok {
					return wildcardNode, ok
				} else {
					params[paramName] = path
					cur = wildcardNode
					continue
				}
			}
			cur = childNode
		}
	}
	if len(params) > 0 {
		cur.params = params
	}
	return cur, cur.isLeaf
}
