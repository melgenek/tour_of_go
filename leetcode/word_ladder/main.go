package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", ladderLength("a", "c", []string{"a", "b", "c"}))
	fmt.Printf("%v\n", 2 == ladderLength("a", "c", []string{"a", "b", "c"}))
	fmt.Printf("%v\n", ladderLength("hot", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Printf("%v\n", 4 == ladderLength("hot", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Printf("%v\n", ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Printf("%v\n", 5 == ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Printf("%v\n", 5 == ladderLength("hit", "cog", []string{"hat", "hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Printf("%v\n", 0 == ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	g := make(map[string][]string)
	for _, word := range wordList {
		n := len(word)
		for i := 0; i < n; i++ {
			mask := fmt.Sprintf("%s*%s", word[0:i], word[i+1:n])
			g[mask] = append(g[mask], word)
			g[word] = append(g[word], mask)
		}
	}

	dist := make(map[string]int)
	q := list.New()

	q.PushFront(beginWord)
	dist[beginWord] = 1

	for q.Len() > 0 {
		el := q.Back()
		q.Remove(el)
		word := el.Value.(string)
		wordDistance := dist[word]

		n := len(word)
		for i := 0; i < n; i++ {
			mask := fmt.Sprintf("%s*%s", word[0:i], word[i+1:n])

			for _, adjacentWord := range g[mask] {
				adjacentDistance, found := dist[adjacentWord]

				if !found || adjacentDistance > wordDistance+1 {
					dist[adjacentWord] = wordDistance + 1
					q.PushFront(adjacentWord)
				}
			}
		}

	}

	result, found := dist[endWord]
	if found {
		return result
	} else {
		return 0
	}
}
