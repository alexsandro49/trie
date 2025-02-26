package trienode

import "strings"

func (t *Trie) deleteHelper(currentNode *Node, word string, index int) bool {
	if index == len(word) {
		if !currentNode.IsLeaf {
			return false
		}

		currentNode.IsLeaf = false

		return len(currentNode.Children) == 0
	}

	c := word[index]

	node, ok := currentNode.Children[rune(c)]
	if !ok {
		return false
	}

	deleteCurrentNode := t.deleteHelper(node, word, index+1)
	if deleteCurrentNode {
		delete(currentNode.Children, rune(c))
		return len(currentNode.Children) == 0 && !currentNode.IsLeaf
	}

	return false
}

func (t *Trie) Delete(word string) {
	t.deleteHelper(t.Root, strings.ToUpper(word), 0)
}
