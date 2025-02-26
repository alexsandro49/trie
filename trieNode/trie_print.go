package trienode

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aquasecurity/table"
)

func (t *Trie) Print(searchedWord string) {
	myTable := table.New(os.Stdout)
	myTable.SetHeaders("Registered Words")
	myTable.AddHeaders("ID", "Word", "ID", "Word")
	myTable.SetHeaderAlignment(table.AlignLeft, table.AlignLeft)
	myTable.SetHeaderColSpans(0, 4)

	myTable.SetAlignment(table.AlignLeft)

	var words []string

	fmt.Println(len(searchedWord))
	if len(searchedWord) > 0 {
		words = t.search(searchedWord)
	} else {
		words = t.collectAllWords()
	}

	for i := 0; i < len(words); i += 2 {
		if i+1 < len(words) {
			myTable.AddRow(
				strconv.Itoa(i), words[i],
				strconv.Itoa(i+1), words[i+1],
			)
		} else {
			myTable.AddRow(strconv.Itoa(i), words[i], "", "")
		}
	}

	myTable.Render()
}
