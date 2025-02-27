package main

import (
	"flag"
	"fmt"

	trienode "github.com/alexsandro49/trie/trieNode"
)

type CommandFlags struct {
	Add    string
	Delete string
	Search string
	List   bool
}

func newCommandFlags() *CommandFlags {
	commandFlags := CommandFlags{}

	flag.StringVar(&commandFlags.Add, "add", "", "Add a new word")
	flag.StringVar(&commandFlags.Delete, "delete", "", "Delete a word")
	flag.StringVar(&commandFlags.Search, "search", "", "Search a word")
	flag.BoolVar(&commandFlags.List, "list", false, "List all registered words")

	flag.Parse()

	return &commandFlags
}

func (cF *CommandFlags) Execute(trie *trienode.Trie) {
	switch {
	case cF.Add != "":
		trie.Add(cF.Add)
	case cF.Delete != "":
		trie.Delete(cF.Delete)
	case cF.Search != "":
		trie.Print(cF.Search)
	case cF.List:
		trie.Print("")
	default:
		fmt.Println("Error: Invalid command!")
	}
}
