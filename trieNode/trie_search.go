package trienode

import "strings"

func (t *Trie) allWordsHelper(node *Node, words *[]string, prefix strings.Builder) {
	if node.IsLeaf {
		*words = append(*words, prefix.String())
	}

	for k, v := range node.Children {
		newPrefix := strings.Builder{}

		newPrefix.WriteString(prefix.String())
		newPrefix.WriteString(string(k))

		t.allWordsHelper(v, words, newPrefix)
	}
}

func (t *Trie) allWords() []string {
	words := []string{}

	t.allWordsHelper(t.Root, &words, strings.Builder{})

	return words
}

func (t *Trie) searchWordsHelper(currentNode *Node, words *[]string, prefix strings.Builder) {
	if currentNode.IsLeaf {
		*words = append(*words, strings.ToUpper(prefix.String()))
	}

	for c, child_node := range currentNode.Children {
		newPrefix := strings.Builder{}
		newPrefix.WriteString(prefix.String())
		newPrefix.WriteRune(c)

		t.searchWordsHelper(child_node, words, newPrefix)
	}

}

func (t *Trie) searchWords(prefix string) []string {
	words := []string{}
	currentNode := t.Root

	for _, c := range strings.ToUpper(prefix) {

		_, ok := currentNode.Children[c]
		if !ok {
			return words
		}

		currentNode = currentNode.Children[c]
	}

	newPrefix := strings.Builder{}
	newPrefix.WriteString(prefix)

	t.searchWordsHelper(currentNode, &words, newPrefix)

	return words
}
