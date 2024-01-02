package follow

import (
	"strings"

	"github.com/i-DarkLight/NRParser/first"
)

var FollowMap = make(map[string]string)
var AllOptions []string

func SetFollow() {
	for _, value := range first.MapTerminal {
		AllOptions = append(AllOptions, strings.Split(value, " | ")...)
	}
	for key := range first.MapTerminal {
		FindFollow(key)
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
func FindFollow(nonterm string) string {
	var NormFollow []string
	for _, option := range AllOptions {
		for i, letter := range option {
			if string(letter) == nonterm {
				if nonterm == "A" {
					temp := FollowMap["A"]
					NormFollow = append(NormFollow, temp+"$")
				}
				temp := FollowMap[string(letter)]
				temp += first.FindOne(option[i+1:])
				NormFollow = append(NormFollow, temp)
				if byte(letter) == option[len(option)-1] || first.ReturnLamda(option[i+1:]) {
					nonTerm, check := returnNonterm(option)
					if check {
						temp := FollowMap[string(letter)] + FindFollow(nonTerm)
						NormFollow = append(NormFollow, temp)
					}
				}
			}
		}
	}
	ls := strings.Join(NormFollow, " ")
	FollowMap[nonterm] = ls
	delDupes()
	return ls
}
func delDupes() {
	var res string
	for key, value := range FollowMap {
		for _, elem := range value {
			if elem == ',' || elem == '~' {
				continue
			}
			if !strings.Contains(res, string(elem)) && elem != ' ' {
				res = res + " " + string(elem)
			} else {
				continue
			}
		}
		FollowMap[key] = res
		res = ""
	}
}
