package src

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Player struct{
	Marker string
}

func NewPlayer(m string) *Player{
	return &Player{m}
}

func (p *Player) Play(b *GameBoard) error{
	tilenum, err := intakeInput(b)
	if err != nil{
		return err
	}
	for !b.EditSlot((tilenum-1)/3, (tilenum-1)%3, p.Marker){
		fmt.Println("Cannot edit over modified tile!\nTry again!")
		tilenum, err = intakeInput(b)
		if err != nil{
			return err
		}
	}
	return nil
}

func intakeInput(b *GameBoard) (int,error){
	b.Display()
	fmt.Print("Enter tile number: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("Something went wrong while reading input. Please try running program again!")
		return -1, errors.New("input failure")
	}
	position := strings.TrimSuffix(input,"\n")
	if len(position) > 1{
		fmt.Println("Inputed more than one tile number!")
		return -1, errors.New("invalid argument")
	}
	tilenum, err := strconv.Atoi(position)
	if err != nil{
		return -1, errors.New("string to int conversion error")
	}
	if tilenum <0 || tilenum >9 {
		fmt.Println("Inputted a number less than 0 or greater than 9!")
		return -1, errors.New("input out of bounds")
	}
	return tilenum, nil
}
