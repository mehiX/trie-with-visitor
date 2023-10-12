package main

// WordsFinder implements a Visitor that registers all the words that it encounters while
// traversing a Trie
type WordsFinder struct {
	words []string
	runes []rune
}

func (w *WordsFinder) push(r rune) {
	w.runes = append(w.runes, r)
}

func (w *WordsFinder) pop() {
	w.runes = w.runes[:len(w.runes)-1]
}

func (w *WordsFinder) markWord() {
	w.words = append(w.words, string(w.runes))
}

func (w *WordsFinder) VisitNode(n *Node) {
	if n.children == nil {
		return
	}

	for r, next := range n.children {
		if r == '*' {
			// mark word
			w.markWord()
			continue
		}
		w.push(r)
		next.Accept(w)
		w.pop()
	}
}

func (w *WordsFinder) Words() []string {
	return w.words
}
