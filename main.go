package main

import (
	trienode "github.com/alexsandro49/trie/trieNode"
)

func main() {
	var trie trienode.Trie

	storage := NewStorage[trienode.Trie]("trie-data.json")
	if storage.Load(&trie) != nil {
		trie = trienode.NewTrie()
	}

	commandFlags := newCommandFlags()
	commandFlags.Execute(&trie)

	storage.Save(&trie)
}
