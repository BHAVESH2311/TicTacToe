package main

import (
	"fmt"
	"tictactoe/game"
)

func main() {
	fmt.Println("-----X and O game------")
	g1 := game.NewGame("bhavesh","swastik")
	for{
		var cellNo uint
		g1.DisplayBoard()
		fmt.Print("Please enter the cell number:")
		fmt.Scan(&cellNo)
		fmt.Println(g1.Play(cellNo))
		if g1.HasGameEnded() {
			break
		}
	}
}