package trie

// Trie is a tree data structure that stores strings
type Trie struct {
	children map[rune]*Trie
	isWord   bool
}

func NewTrie() *Trie {
	return &Trie{children: make(map[rune]*Trie)}
}

// Insert inserts a string into the trie
func (t *Trie) Insert(s string) {
	curr := t
	for _, c := range s {
		if _, ok := curr.children[c]; !ok {
			curr.children[c] = NewTrie()
		}
		curr = curr.children[c]
	}
	curr.isWord = true
}

// Search returns true if the string is in the trie
func (t *Trie) Search(s string) bool {
	curr := t
	for _, c := range s {
		if _, ok := curr.children[c]; !ok {
			return false
		}
		curr = curr.children[c]
	}
	return curr.isWord
}

// StartsWith returns true if there is any string in the trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	curr := t
	for _, c := range prefix {
		if _, ok := curr.children[c]; !ok {
			return false
		}
		curr = curr.children[c]
	}
	return true
}
