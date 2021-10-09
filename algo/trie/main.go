package main

import "fmt"

func main() {
	t := Constructor()
	fmt.Printf("%v\n", t.Search("ab") == false)
	t.Insert("abc")
	fmt.Printf("%v\n", t.Search("ab") == false)
	fmt.Printf("%v\n", t.Search("abc") == true)
	fmt.Printf("%v\n", t.StartsWith("ab") == true)
	fmt.Printf("%v\n", t.StartsWith("abc") == true)
	fmt.Printf("%v\n", t.StartsWith("ac") == false)
	t.Insert("cc")
	fmt.Printf("%v\n", t.StartsWith("a") == true)
	fmt.Printf("%v\n", t.StartsWith("c") == true)
}

type Trie struct {
	letters [26]*Trie
	isWord  bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
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

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	res := this.search(word)
	return res != nil && res.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.search(prefix) != nil
}

func (this *Trie) search(prefix string) *Trie {
	t := this
	for i := range prefix {
		v := prefix[i] - 'a'
		if t.letters[v] == nil {
			return nil
		}
		t = t.letters[v]
	}
	return t
}
