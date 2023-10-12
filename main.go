// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	t := NewTrie()
	t.Add("something", "some", "cat", "dog", "slog", "doggie", "bag")
	fmt.Println(t.AllWords())

	freq := t.LettersFrequency()
	for r, c := range freq {
		fmt.Printf("%c: %d\n", r, c)
	}

}
