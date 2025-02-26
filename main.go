package main

import trienode "github.com/alexsandro49/trie/trieNode"

func main() {
	trie := trienode.NewTrie()

	trie.Add("Batata")
	trie.Add("Barco")
	trie.Add("Goiaba")
	trie.Add("Gengibre")
	trie.Add("Motocicleta")
	trie.Add("Fantasma")
	trie.Add("Fantoche")
	trie.Add("Fantastico")

	trie.Print("")
}
