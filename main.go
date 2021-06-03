package main

import (
	game "TicTacTerminal/src"
	"fmt"
	"log"
)
var winner int
func main(){
	b := game.NewBoard()
	fmt.Println("Welcome to TicTacTerminal, a 2-player TicTacToe game ported to the terminal!\nPlayer 1 is X and Player 2 is 0!\nEnter them in the tiles on the board shown below!")
	b.DisplayExample()
	fmt.Println("----BEGIN GAME----")
	p1 := game.NewPlayer("X")
	p2 := game.NewPlayer("O")
	for turns:=0; turns<8; turns++ {
		if turn(p1, b){
			break
		}
		if turn(p2, b){
			break
		}
	}
	if winner == 1{
		fmt.Println("Player 1 wins!")
	} else if winner == 2{
		fmt.Println("Player 2 wins!")
	} else {
		fmt.Println("Draw!")
	}
	b.Display()
}

func turn(p *game.Player, b *game.GameBoard) bool{
	err := p.Play(b)
	 if err !=nil{
	 	log.Fatal(err)
	 }
	 condition := b.CheckBoard()
	 if condition > 0{
	 	winner = condition
	 	return true
	 }
	 return false
}



