package table

import (
	"fmt"
	"github.com/i-DarkLight/NRParser/first"
	"strings"
)

var Terms []string

func SaveToTable(r, c int, flag bool) {
	table := make([][]string, r+1)
	for v := range table {
		table[v] = make([]string, c+2)
	}
	i := 1
	for key := range first.MapTerminal {
		table[i][0] = key
		i++
	}
	for x := 1; x < c+2; x++ {
		if x > c {
			table[0][x] = "$"
			continue
		}
		table[0][x] = Terms[x-1]
	}
	for nonTerm := 1; nonTerm < r+1; nonTerm++ {
		for k := 1; k < c+2; k++ {
			if table[0][k] == "$" {
				table[nonTerm][k] = "~"
			}
			table[nonTerm][k] = returnOptionStats(table[0][k], table[nonTerm][0])
		}
	}
	if flag {
		fmt.Printf("%v", table)
	}
}

func returnOptionStats(x, NonTerm string) string {
	for _, value := range first.AllOptions {
		if strings.Contains(value.First, x) && value.NonTerm == NonTerm {
			return value.Option
		}
	}
	return ""
}
