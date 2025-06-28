package main

import (
	"fmt"
	"math/rand"
)

var words = []string{
	"Zombie",
	"Vampire",
	"Ghost",
	"Monster",
	"Dragon",
	"Wizard",
	"Elf",
	"Goblin",
	"Orc",
}

func main() {
	target := words[rand.Intn(len(words))]
	// fmt.Println(target)
	guesses := map[rune]bool{}
	guesses[rune(target[0])] = true // Start with the first letter guessed
	guesses[rune(target[len(target)-1])] = true // Start with the last letter guessed

	gameState(target,guesses)
}

func gameState(target string, guesses map[rune]bool){
	for _,c:=range target{
		if guesses[c] == true {
			fmt.Printf("%c", c)
		}else{
			fmt.Printf("_")
		}
		fmt.Print(" ")
	}
	fmt.Println()
}