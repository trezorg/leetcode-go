package prefixtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixTree(t *testing.T) {
	trie := Constructor()
	asserts := assert.New(t)
	trie.Insert("apple")
	asserts.True(trie.Search("apple"))
	asserts.False(trie.Search("app"))
	asserts.True(trie.StartsWith("app"))
	trie.Insert("app")
	asserts.True(trie.Search("app"))
	asserts.False(trie.Search("ap"))
	asserts.False(trie.Search("a"))
	asserts.False(trie.Search(""))
	asserts.True(trie.StartsWith("a"))
	asserts.False(trie.StartsWith("b"))
	trie.Insert("banana")
	asserts.True(trie.StartsWith("b"))
	asserts.True(trie.StartsWith("ban"))
	asserts.False(trie.Search("ban"))
	asserts.True(trie.Search("banana"))
}
