package src

import "fmt"

type GameBoard struct{
	Board [3][3]string
	Full bool
}

func NewBoard() *GameBoard{
	return &GameBoard{[3][3]string{{" "," ", " "}, {" "," ", " "},{" "," ", " "}}, false}
}

func (b *GameBoard) EditSlot(r,c int, marker string) bool{
	if b.Board[r][c] == " " {
		b.Board[r][c] = marker
		return true
	}
	return false
}

func (b *GameBoard) CheckBoard() int{
		//check rows
		for r := 0; r < 3; r++ {
			if b.Board[r][0] == b.Board[r][1] && b.Board[r][0] == b.Board[r][2] {
				winner := b.findWinnerFromBoardSlot(r,0)
				if winner > 0{
					return winner
				}
			}
		}
		//check cols
		for c := 0; c < 3; c++ {
			if b.Board[0][c] == b.Board[1][c] && b.Board[0][c] == b.Board[2][c] {
				winner := b.findWinnerFromBoardSlot(0,c)
				if winner > 0 {
					return winner
				}
			}
		}
		//check diagonals
		if b.Board[0][0] == b.Board[1][1] && b.Board[0][0] == b.Board[2][2]{
			winner := b.findWinnerFromBoardSlot(0,0)
			if winner > 0 {
				return winner
			}
		} else if b.Board[2][0] == b.Board[1][1] && b.Board[1][1] == b.Board[0][2] {
			winner := b.findWinnerFromBoardSlot(1,1)
			if winner > 0 {
				return winner
			}
		} else if b.isFull(){ //draw case = 3
			return 3
		}
		return 0
}

func (b *GameBoard) isFull() bool{
	for r:=0;r<3;r++{
		for c:=0;c<3;c++{
			if b.Board[r][c] == " "{
				return false
			}
		}
	}
	return true
}

func (b *GameBoard) findWinnerFromBoardSlot(r,c int) int{
	if b.Board[r][c] == "X" {
		return 1
	} else if b.Board[r][c] == "O" {
		return 2
	}
	return 0
}

func (b *GameBoard) Display() {
	for r:=0; r<3; r++{
		fmt.Println(b.Board[r][0] + " | " + b.Board[r][1] + " | " + b.Board[r][2])
		if r < 2 {
			fmt.Println("---------")
		}
	}
}

func (b *GameBoard) DisplayExample() {
	i:=1
	for r:=0; r<3; r++{
		fmt.Println(i,"|",i+1,"|",i+2)
		i+=3
		if r < 2 {
			fmt.Println("---------")
		}
	}
}