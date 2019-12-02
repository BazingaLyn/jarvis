package trie

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestURITrieAdd(t *testing.T) {

	trie := NewTrie()
	trie.AddNode("/user/add/1", nil)
	trie.AddNode("/user/:id", nil)
	trie.AddNode("/user/:id/:id3", nil)
	trie.AddNode("/user/:id/:id2", nil)

	fmt.Println(trie.Size)
}

func TestURITrieSearch(t *testing.T) {

	trie := NewTrie()
	trie.AddNode("/user/add/1", nil)
	trie.AddNode("/user/:id", nil)
	trie.AddNode("/user/:id/:id3", nil)
	trie.AddNode("/user/:id/:id2", nil)

	node, exist := trie.Search("/user/2/3")
	if exist {
		fmt.Println("exist")
		bytes, err := json.Marshal(node.Params)
		if err == nil {
			fmt.Println(string(bytes))
		}
	} else {
		fmt.Println("not exist")
	}
}
