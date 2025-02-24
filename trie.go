package main

import "strings"

type Trie struct {
	Root *Node
}

func NewTrie() Trie {
	return Trie{
		&Node{
			make(map[rune]*Node),
			false,
		},
	}
}

func (t *Trie) Add(word string) {
	current := t.Root

	for _, v := range word {
		_, ok := current.Children[v]

		if ok {
			current = current.Children[v]

		} else {
			current.Children[v] = &Node{
				make(map[rune]*Node),
				false,
			}

			current = current.Children[v]
		}
	}

	current.IsLeaf = true
}

func (t *Trie) collectAllWords() []string {
	words := []string{}

	t.collectAllWordsHelper(t.Root, &words, strings.Builder{})

	return words
}

func (t *Trie) collectAllWordsHelper(node *Node, words *[]string, prefix strings.Builder) {
	if node.IsLeaf {
		*words = append(*words, prefix.String())
	}

	for k, v := range node.Children {
		newPrefix := strings.Builder{}

		newPrefix.WriteString(prefix.String())
		newPrefix.WriteString(string(k))

		t.collectAllWordsHelper(v, words, newPrefix)
	}
}
