package main

import (
	"github.com/i-DarkLight/NRParser/first"
	"github.com/i-DarkLight/NRParser/follow"
)

func main() {
	first.FindAllTogether("A -> aAs | nCb | BcD | BDg\nB -> ko | pAo | ~\nC -> qr | jAu\nD -> l | ~")
	follow.FindFollow()
	println("================================")
	for key, value := range follow.FollowMap {
		println(key + " " + value)
	}
	println("================================")
}
