package main

type Node struct {
	Children map[rune]*Node
	IsLeaf   bool
}
