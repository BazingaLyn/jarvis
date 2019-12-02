package trie

import (
	"net/http"
	"strings"
)

var colonCharacter string = ":"

type Handle func(http.ResponseWriter, *http.Request)

type Node struct {
	IsLeaf   bool
	Wildcard bool
	Handle   Handle
	Params   map[string]string
	Children map[string]*Node
}

type UriTrie struct {
	Root *Node
	Size int
}

func NewTrie() UriTrie {
	return UriTrie{
		Root: &Node{
			Children: make(map[string]*Node),
		},
		Size: 0,
	}
}

func (urlTrie *UriTrie) AddNode(uri string, handler Handle) {

	paths := strings.Split(uri, "/")

	cur := urlTrie.Root
	for _, path := range paths {
		if len(path) != 0 {

			// 通配符处理 简单实现
			isWildcard := isWildcard(path)
			if isWildcard {
				wildcardNode, _, ok := getChildWildcard(cur.Children)
				if !ok {
					wildcardNode = &Node{
						Wildcard: true,
						Children: make(map[string]*Node),
					}
					cur.Children[getCommonPath(path)] = wildcardNode
				}
				cur = wildcardNode

			} else {
				childNode, ok := cur.Children[path]
				if !ok {
					childNode = &Node{
						Children: make(map[string]*Node),
					}
					cur.Children[path] = childNode
				}
				cur = childNode
			}
		}
	}

	if !cur.IsLeaf {
		cur.IsLeaf = true
		cur.Handle = handler
		urlTrie.Size++
	}

}

func getCommonPath(path string) string {
	return strings.TrimPrefix(path, ":")

}

func getChildWildcard(nodes map[string]*Node) (*Node, string, bool) {

	for k, v := range nodes {
		if v.Wildcard {
			return v, k, true
		}
	}
	return nil, "", false
}

func isWildcard(path string) bool {
	return strings.HasPrefix(path, colonCharacter)
}

func (urlTrie *UriTrie) Search(uri string) (*Node, bool) {

	paths := strings.Split(uri, "/")

	cur := urlTrie.Root
	params := make(map[string]string)

	for _, path := range paths {
		if len(path) > 0 {
			childNode, ok := cur.Children[path]
			if !ok {
				wildcardNode, paramName, ok := getChildWildcard(cur.Children)
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
		cur.Params = params
	}
	return cur, cur.IsLeaf
}
