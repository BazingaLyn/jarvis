package trie

import (
	"fmt"
	"testing"
)

func TestURITrieAdd(t *testing.T) {

	trie := NewTrie()
	trie.addNode("/user/add/1")
	trie.addNode("/user/:id")
	trie.addNode("/user/:id/:id3")
	trie.addNode("/user/:id/:id2")

	fmt.Println(trie.Size)
}

func TestURITrieSearch(t *testing.T) {

	trie := NewTrie()
	trie.addNode("/user/add/1")
	trie.addNode("/user/:id")
	trie.addNode("/user/:id/:id3")
	trie.addNode("/user/:id/:id2")

	_, exist := trie.search("/user/2/3")
	if exist {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}
}
