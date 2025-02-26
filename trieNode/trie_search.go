package trienode

import "strings"

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

func (t *Trie) searchHelper(currentNode *Node, words *[]string, prefix strings.Builder) {
	if currentNode.IsLeaf {
		*words = append(*words, strings.ToUpper(prefix.String()))
	}

	for c, child_node := range currentNode.Children {
		newPrefix := strings.Builder{}
		newPrefix.WriteString(prefix.String())
		newPrefix.WriteRune(c)

		t.searchHelper(child_node, words, newPrefix)
	}

}

func (t *Trie) search(prefix string) []string {
	words := []string{}
	currentNode := t.Root

	for _, c := range strings.ToUpper(prefix) {

		_, ok := currentNode.Children[c]
		if !ok {
			words = append(words, "Not Found")
			return words
		}

		currentNode = currentNode.Children[c]
	}

	newPrefix := strings.Builder{}
	newPrefix.WriteString(prefix)

	t.searchHelper(currentNode, &words, newPrefix)

	return words
}
