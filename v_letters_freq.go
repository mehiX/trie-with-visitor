package main

// lettersFreq traverses a Trie and records the frequency with which letters appear.
// The values returned show how often a letter appears in the Trie, not in the words
// in the dictionary. Comparing these values with the frequencies of the letters in
// the words show how efficiently the Trie stores the dictionary.
type lettersFreq struct {
	freq map[rune]int
}

func NewLettersFreq() *lettersFreq {
	return &lettersFreq{freq: make(map[rune]int)}
}

func (l *lettersFreq) VisitNode(n *Node) {

	if n == nil {
		return
	}

	for r, next := range n.children {
		if r != '*' {
			l.freq[r]++
			next.Accept(l)
		}
	}
}

func (l *lettersFreq) Letters() map[rune]int {
	return l.freq
}
