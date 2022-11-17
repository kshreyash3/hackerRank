package main

import (
	"log"
	"strings"
)

func main() {
	str := "code swarm"
	words := strings.Split(str, " ")
	key := strings.ToLower(words[0])
	for _, word := range words[1:] {
		key += strings.Title(word)
	}
	log.Println(key)
}
