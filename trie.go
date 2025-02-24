package main

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
