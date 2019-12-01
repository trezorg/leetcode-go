package prefixtree

type TrieDict map[byte]*Trie

type Trie struct {
	stop   bool
	values TrieDict
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{values: make(TrieDict)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		s := word[i]
		_, ok := root.values[s]
		if !ok {
			root.values[s] = &Trie{values: make(TrieDict)}
		}
		root = root.values[s]
	}
	root.stop = true
}

func (this *Trie) _search(word string) *Trie {
	root := this
	var ok bool
	for i := 0; i < len(word); i++ {
		s := word[i]
		root, ok = root.values[s]
		if !ok {
			return nil
		}
	}
	return root
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this._search(word)
	return root != nil && root.stop
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this._search(prefix) != nil
}
