package first

import (
	"strings"
	"unicode"
)

var MapTerminal = make(map[string]string)
var FirstList = make(map[string]string)
func eachLine(nonterm string, options string) string {
	optionsList := strings.Split(options, " | ")
	var firstCollection []string
out:
//Iterate over options of a given non terminal
	for _, elem := range optionsList {
		//Iterate over letters of a option in case an option starts with another non terminal
		for i := 0; i < len(elem); i++ {
			if unicode.IsLower(rune(elem[i])) || elem[i] == '~' {
				firstCollection = append(firstCollection, string(elem[i]))
				break
			}	else {
				//if a non terminal is reached, find first for that non terminal
				newLine := MapTerminal[string(elem[i])]
				nextterm := string(elem[i])
				str := eachLine(nextterm,newLine)
				//if non terminal does not have lamda as an option go to next option in main non terminal
				if !strings.Contains(str, "~") {
					toSlice := strings.Split(str, ",")
					firstCollection = append(firstCollection, toSlice...)
					break out
				} else if strings.Contains(str, "~") {
					//if non terminal does contaion lamda, add the terminal of after that non terminal
					newLine := MapTerminal[string(elem[i])]
					nextterm:=string(elem[i])
					str := eachLine(nextterm,newLine)
					toSlice := strings.Split(str, ",")
					firstCollection = append(firstCollection, toSlice...)
				}
			}
		}
	}
	//gather up all first terminals and append to FirstList
	list := strings.Join(firstCollection, ",")
	FirstList[nonterm] = list
	return list
}
func Find(str string) {
	lines := strings.Split(str, "\n")
	//save non terminals as key and their respective options as value in a map
	for _, elem := range lines {
		pos := strings.Index(elem, "->")
		MapTerminal[elem[:pos-1]] = elem[pos+3:]
	}
	//calculate all firsts
	for key, value := range MapTerminal {
		eachLine(key, value)
	}
	for key, value := range FirstList {
		println(key + "  " + value)
	}
}
