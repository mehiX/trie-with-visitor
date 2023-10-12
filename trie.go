package main

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

func (t *Trie) Add(words ...string) {
	for _, w := range words {
		t.add(w)
	}
}

func (t *Trie) add(word string) {
	runes := []rune(word)
	curr := t.root

	for _, r := range runes {
		next, ok := curr.children[r]
		if !ok {
			next = NewNode()
			curr.children[r] = next
		}
		curr = next
	}

	curr.children['*'] = nil
}

func (t *Trie) AllWords() []string {
	finder := &WordsFinder{}
	t.root.Accept(finder)
	return finder.Words()
}

func (t *Trie) LettersFrequency() map[rune]int {
	v := NewLettersFreq()
	t.root.Accept(v)

	return v.Letters()
}
