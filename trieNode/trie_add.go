package trienode

import "strings"

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
