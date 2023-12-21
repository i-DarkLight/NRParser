package main

import (
	//	"os"
	"strings"

	"github.com/i-DarkLight/NRParser/first"
	"github.com/i-DarkLight/NRParser/follow"
)

func main() {
	Str := "A -> aAs | nCb | CcD\nB -> o | wo \nC -> BA | jAu\nD -> hDl | ~"
	lines := strings.Split(Str, "\n")
	//save non terminals as key and their respective options as value in a map
	for _, elem := range lines {
		pos := strings.Index(elem, "->")
		first.MapTerminal[elem[:pos-1]] = elem[pos+3:]
	}
	first.FindAllTogether()
	follow.SetFollow()
	println("================================")
	for key, value := range follow.FollowMap {
		println(key + ": " + value)
	}
	println("================================")
}
