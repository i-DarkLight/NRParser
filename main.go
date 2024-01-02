package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/i-DarkLight/NRParser/first"
	"github.com/i-DarkLight/NRParser/follow"
	"github.com/i-DarkLight/NRParser/table"
)

func newGrammar() {
	println("\n\ninput a grammar(shouldn't have Left Recursive problem):")
	input := bufio.NewScanner(os.Stdin)
	var lines []string
	for {
		input.Scan()
		line := input.Text()
		if line == "END" {
			break
		}
		lines = append(lines, line)
	}
	temp := strings.Join(lines, "\n")
	var terms []string
	for _, v := range temp {
		if unicode.IsLower(v) {
			str := strings.Join(terms, "")
			if !strings.Contains(str, string(v)) {
				terms = append(terms, string(v))
			}
		}
	}
	println(len(terms))
	table.Terms = terms
	grammar := strings.Split(temp, "\n")
	for _, elem := range grammar {
		pos := strings.Index(elem, "->")
		if pos == -1 {
			print("not valid grammar")
			fmt.Printf("\033[1A\033[K")
			time.Sleep(2 * time.Second)
			newGrammar()
		} else {
			first.MapTerminal[elem[:pos-1]] = elem[pos+3:]
		}
	}
	print("your grammar is successfully saved! :)\nwhat's next?\n1- see options for your grammar\n2- change grammar\n")
	first.SetFirst()
	first.FirstOptions()
	follow.SetFollow()
	rows := len(first.MapTerminal)
	cols := len(terms)
	table.SaveToTable(rows, cols,false)
	var inp string
	fmt.Scan(&inp)
	switch inp {
	case "1":
		var inpu string
		println("1- see firsts\n2- see follows\n3- see table\n4- test a string\n")
		fmt.Scan(&inpu)
		switch inpu {
		case "1":
			for key, value := range first.FirstList {
				println(key + ": " + value)
			}
		case "2":
			for key, value := range follow.FollowMap {
				println(key + ": " + value)
			}
		case "3":
			table.SaveToTable(rows, cols,true)
		case "4":
			//see if a given string is accepted by grammar
		}
	case "2":
		newGrammar()
	}
	//	fmt.Printf("\033[1A\033[K")
}
func main() {
	print("Welcome to my small project :)\n\nthis project is a NonRecursiveParser!")
	newGrammar()

}
