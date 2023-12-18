package main

import (
	"strings"
	"github.com/i-DarkLight/NRParser/first"
	"github.com/i-DarkLight/NRParser/follow"
)

func catchError() bool {
	for key:=range first.MapTerminal{
		first.FindErr(key)
	}
	for key, value := range first.FirstErr {
		if strings.Contains(value, key) {
			return true
		}
	}
	return false
}
func main() {
	str := "A -> aAs | nCb | BcD\nB -> xo | yo | ~\nC -> Br | jAu\nD -> l | ~"
	lines := strings.Split(str, "\n")
	//save non terminals as key and their respective options as value in a map
	for _, elem := range lines {
		pos := strings.Index(elem, "->")
		first.MapTerminal[elem[:pos-1]] = elem[pos+3:]
	}
	catchError()
	flag := false
	for range lines {
		if catchError() {
			flag = true
		}
	}
	if !flag {
		first.FindAllTogether()
	} else {
		println("there is a first/first problem in your grammar")
	}
	follow.FindFollow()
	println("================================")
	for key, value := range follow.FollowMap {
		println(key + " " + value)
	}
	println("================================")
}
