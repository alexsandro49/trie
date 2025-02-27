package trienode

type Node struct {
	Children map[rune]*Node `json:"children"`
	IsLeaf   bool           `json:"is_leaf"`
}
