package trie

import (
	"fmt"
	"testing"
)

func TestURITrieAdd(t *testing.T) {

	trie := NewTrie()
	trie.addNode("/user/add/1")
	trie.addNode("/user/add/2")
	trie.addNode("/user/add")
	trie.addNode("/user/add/3")

	fmt.Println(trie.Size)
}

func TestURITrieSearch(t *testing.T) {

	trie := NewTrie()
	trie.addNode("/user/add/1")
	trie.addNode("/user/add/2")
	trie.addNode("/user/add")
	trie.addNode("/user/add/3")

	_, exist := trie.search("/user/add/3/hello")
	if exist {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}

}
