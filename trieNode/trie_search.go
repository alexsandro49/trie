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

func (t *Trie) search(prefix string) []string {
	words := []string{}

	for _, c := range strings.ToUpper(prefix) {

		node, ok := t.Root.Children[c]
		if ok {
			newPrefix := strings.Builder{}
			newPrefix.WriteRune(c)

			t.collectWordsHelper(node, &words, newPrefix)
		}
	}

	if len(words) == 0 {
		words = append(words, "Not Found")
	}

	return words
}
