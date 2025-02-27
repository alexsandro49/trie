package trienode

type Trie struct {
	Root *Node `json:"root"`
}

func NewTrie() Trie {
	return Trie{
		&Node{
			make(map[rune]*Node),
			false,
		},
	}
}
