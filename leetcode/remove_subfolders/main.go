package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Printf("%v\n", removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))
	fmt.Printf("%v\n", removeSubfolders([]string{"/a", "/a/b/c", "/a/b/d"}))
	fmt.Printf("%v\n", removeSubfolders([]string{"/a/b/c", "/a/b/ca", "/a/b/d"}))
}

type Trie struct {
	prefixes map[string]*Trie
	folder   bool
}

func NewTrie() *Trie {
	return &Trie{
		make(map[string]*Trie),
		false,
	}
}

func (this *Trie) Add(str string) {
	t := this

	parts := strings.Split(str, "/")
	for i := 1; i < len(parts); i++ {
		part := parts[i]
		if t.prefixes[part] == nil {
			t.prefixes[part] = NewTrie()
		}
		t = t.prefixes[part]
		if t.folder {
			return
		}
	}
	t.folder = true
}

func (this *Trie) Folders() []string {
	return trimmedFolders(this, []string{})
}

func trimmedFolders(t *Trie, parts []string) []string {
	if len(t.prefixes) == 0 {
		return []string{fmt.Sprintf("/%s", strings.Join(parts, "/"))}
	} else {
		res := []string{}
		for part, subTrie := range t.prefixes {
			res = append(res, trimmedFolders(subTrie, append(parts, part))...)
		}
		return res
	}
}

func removeSubfolders(folders []string) []string {
	sort.Strings(folders)
	t := NewTrie()
	for _, folder := range folders {
		t.Add(folder)
	}
	return t.Folders()
}
