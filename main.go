package main

func main() {
	trie := NewTrie()

	trie.Add("Batata")
	trie.Add("Barco")
	trie.Add("Goiaba")
	trie.Add("Gengibre")
	trie.Add("Motocicleta")
	trie.Add("Fantasma")
	trie.Add("Fantoche")
	trie.Add("Fantastico")

	trie.Print()
}
