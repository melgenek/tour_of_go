package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", suggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
	fmt.Printf("%v\n", suggestedProducts([]string{"havana"}, "havana"))
	fmt.Printf("%v\n", suggestedProducts([]string{"havana"}, "tatiana"))
}

type Trie struct {
	letters [26]*Trie
	isWord  bool
}

func (this *Trie) add(word string) {
	t := this
	for i := range word {
		v := word[i] - 'a'
		if t.letters[v] == nil {
			t.letters[v] = &Trie{}
		}
		t = t.letters[v]
	}
	t.isWord = true
}

func (this *Trie) search(prefix string, limit int) []string {
	t := this
	for i := range prefix {
		v := prefix[i] - 'a'
		t = t.letters[v]
		if t == nil {
			break
		}
	}

	var b bytes.Buffer
	b.WriteString(prefix)

	if t == nil {
		return []string{}
	} else {
		return suggestion(b, t, []string{}, limit)
	}
}

func suggestion(b bytes.Buffer, t *Trie, arr []string, n int) []string {
	currentLength := b.Len()
	res := arr
	if t.isWord {
		res = append(res, b.String())
	}
	if len(arr) == n {
		return arr
	}

	for i, next := range t.letters {
		if next != nil {
			b.WriteByte('a' + byte(i))
			res = suggestion(b, next, res, n)
			b.Truncate(currentLength)
		}
	}
	return res

}

func suggestedProducts(products []string, searchWord string) [][]string {
	t := &Trie{}
	for _, v := range products {
		t.add(v)
	}

	res := make([][]string, len(searchWord))

	for i := 1; i <= len(searchWord); i++ {
		prefix := searchWord[:i]
		res[i-1] = t.search(prefix, 3)
	}

	return res
}
