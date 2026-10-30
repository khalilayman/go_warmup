package main

import (
	"fmt"
	"os"
)

const MaxDisplay = 10

func main() {

	words := os.Args[1:]

	if len(words) == 0 {
		fmt.Fprintln(os.Stderr, "Erreur saisir un mot")
		os.Exit(1)
	}

	fmt.Printf("Nombre total de mots : %d\n", len(words))

	fmt.Println("Mots de longeur > 4 ")
	for _, w := range words {

		
		if len(w) > 4 {
			fmt.Println(w)
		}
	}

}
