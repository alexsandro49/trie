package trienode

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
