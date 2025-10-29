package main

import (
	"fmt"

	"CleanMyWorkspace/Clean"

	"github.com/Mentor-Paris/CleanMyWorkspace"
)

func main() {
	workspace := CleanMyWorkspace.GenererateWorkSpace()

	fmt.Println("Souk Avant le passage du nettoyeur 2000")
	for _, row := range *workspace {
		fmt.Print("|")
		for _, cell := range row {
			if cell == nil {
				fmt.Print("nil|")
			} else {
				fmt.Printf("%s|", *cell)
			}
		}
		fmt.Println()
	}

	cleanedWorkspace := Clean.CleanWorkSpace(workspace)

	fmt.Println("\nSouk nettoyé par le nettoyeur 2000")
	for _, row := range *cleanedWorkspace {
		fmt.Print("|")
		for _, cell := range row {
			if cell == nil {
				fmt.Print("nil|")
			} else {
				fmt.Printf("%s|", *cell)
			}
		}
		fmt.Println()
	}

	fmt.Println("\nLes mentors sont contents j'ai jeté ma canette dans la bonne poubelle !")
}
