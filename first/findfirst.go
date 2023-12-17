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
			if unicode.IsLower(rune(elem[i])) {
				firstCollection = append(firstCollection, string(elem[i]))
				break
			} else if elem == "~" {
				firstCollection = append(firstCollection, "~")
			} else {
				//if a non terminal is reached, find first for that non terminal
				newLine := MapTerminal[string(elem[i])]
				nextterm := string(elem[i])
				str := eachLine(nextterm, newLine)
				//if non terminal does not have lamda as an option go to next option in main non terminal
				if !strings.Contains(str, "~") {
					toSlice := strings.Split(str, ",")
					firstCollection = append(firstCollection, toSlice...)
					continue out
				} else if strings.Contains(str, "~") {
					//if non terminal does contaion lamda, add the terminal of after that non terminal
					newLine := MapTerminal[string(elem[i])]
					nextterm := string(elem[i])
					str := eachLine(nextterm, newLine)
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
func FindAllTogether(str string) {
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
	finalCheck()
	for key, value := range FirstList {
		println(key + "  " + value)
	}
}
func testLamda(str string) bool {
	for key, value := range MapTerminal {
		if str == key {
			temp := strings.Split(value, " | ")
			for _, elem := range temp {
				if ReturnLamda(elem) || elem == "~" {
					return true
				}
			}
		}
	}
	return false
}

// function to debug calculation of lamda
func finalCheck() {
	var temp string
	var res string
	//loop to iterate over first list and delete duplicate or unwanted elementes
	for key, value := range FirstList {
		for _, elem := range value {
			if elem == ',' {
				continue
			}
			temp += string(elem)
			//if term is already in first list do not add it to the list again
			if !strings.Contains(res, string(elem)) && elem != '~' {
				res += " " + string(elem)
			} else {
				continue
			}
		}
		if testLamda(key) {
			res += " ~"
		}
		FirstList[key] = res
		res, temp = "", ""
	}
}

// To calculate first of one given option
func FindOne(str string) string {
	var temp string
	for _, letter := range str {
		if unicode.IsLower(letter) {
			temp += string(letter)
			break
		} else {
			temp += FirstList[string(letter)]
			if strings.Contains(temp, "~") {
				continue
			} else {
				break
			}
		}
	}
	ret := temp
	return ret
}

// function checks if a option can generate lamda
func ReturnLamda(str string) bool {
	flag := true
	for _, letter := range str {
		if unicode.IsLower(letter) || (unicode.IsUpper(letter) && !strings.Contains(FirstList[string(letter)], "~")) {
			flag = false
			break
		}
	}
	return flag
}
