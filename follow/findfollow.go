package follow

import (
	"github.com/i-DarkLight/NRParser/first"
	"strings"
	"unicode"
)

var FollowMap = make(map[string]string)
var AllOptions []string

func setOptions() {
	for _, value := range first.MapTerminal {
		AllOptions = append(AllOptions, strings.Split(value, " | ")...)
	}
}
func returnNonterm(str string) (string, bool) {
	for key, value := range first.MapTerminal {
		if strings.Contains(value, str) {
			return key, true
		}
	}
	return "", false
}
func FindFollow() {
	setOptions()
	for _, option := range AllOptions {
		for i, letter := range option {
			if unicode.IsUpper(letter) {
				if byte(letter) != option[len(option)-1] || !first.ReturnLamda(option[i+1:]) {
					temp := FollowMap[string(letter)]
					FollowMap[string(letter)] = temp + first.FindOne(option[i+1:])
				}
			}
		}
	}
	for _, option := range AllOptions {
		for i, letter := range option {
			if unicode.IsUpper(letter) {
				if byte(letter) == option[len(option)-1] || first.ReturnLamda(option[i+1:]) {
					nonterm, check := returnNonterm(option)
					if check {
						temp := FollowMap[string(letter)]
						FollowMap[string(letter)] = temp + FollowMap[nonterm]
					} else {
						continue
					}
				}
			}
		}
	}
	delDupes()
}
func delDupes() {
	var temp string
	var res string
	for key, value := range FollowMap {
		for _, elem := range value {
			if elem == ',' || elem == '~' {
				continue
			}
			temp += string(elem)
			if !strings.Contains(res, string(elem)) {
				res+=" "+string(elem)
			} else {
				continue
			}
		}
		FollowMap[key] = res
		res, temp = "", ""
	}
}
