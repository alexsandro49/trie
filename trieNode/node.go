package trienode

type Node struct {
	Children map[rune]*Node
	IsLeaf   bool
}
