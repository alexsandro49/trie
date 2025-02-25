package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/aquasecurity/table"
)

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

	for _, v := range strings.ToUpper(word) {
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

func (t *Trie) collectWordsHelper(node *Node, words *[]string, prefix strings.Builder) {
	if node.IsLeaf {
		*words = append(*words, prefix.String())
	}

	for k, v := range node.Children {
		newPrefix := strings.Builder{}

		newPrefix.WriteString(prefix.String())
		newPrefix.WriteString(string(k))

		t.collectWordsHelper(v, words, newPrefix)
	}
}

func (t *Trie) collectAllWords() []string {
	words := []string{}

	t.collectWordsHelper(t.Root, &words, strings.Builder{})

	return words
}

func (t *Trie) search(prefix string) []string {
	words := []string{}

	for _, c := range strings.ToUpper(prefix) {

		node, ok := t.Root.Children[c]
		if ok {
			newPrefix := strings.Builder{}
			newPrefix.WriteRune(c)

			t.collectWordsHelper(node, &words, newPrefix)
		}
	}

	return words
}

func (t *Trie) Print(searchedWord string) {
	myTable := table.New(os.Stdout)
	myTable.SetHeaders("Registered Words")
	myTable.AddHeaders("ID", "Word", "ID", "Word")
	myTable.SetHeaderAlignment(table.AlignLeft, table.AlignLeft)
	myTable.SetHeaderColSpans(0, 4)

	myTable.SetAlignment(table.AlignLeft)

	var words []string

	if len(searchedWord) > 0 {
		words = t.search(searchedWord)
	} else {
		words = t.collectAllWords()
	}

	for i := 0; i < len(words); i += 2 {
		if i+1 < len(words) {
			myTable.AddRow(
				strconv.Itoa(i), words[i],
				strconv.Itoa(i+1), words[i+1],
			)
		} else {
			myTable.AddRow(strconv.Itoa(i), words[i], "", "")
		}
	}

	myTable.Render()
}
