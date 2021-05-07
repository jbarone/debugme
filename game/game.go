// Package game provides a broken implementation of the
// game Rock, Paper, Scissors.
package game

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

// A Move in the game
type Move int

// Valid moves
const (
	Rock Move = iota + 1
	Paper
	Scissors
)

func (move Move) String() string {
	switch move {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		return `¯\_(ツ)_/¯`
	}
}

// Compare moves to see who won.
func (move Move) Compare(move2 Move) int {
	switch {
	case move == move2:
		return 0
	case move == Rock && move2 == Paper:
		return 2
	case move == Rock && move2 == Scissors:
		return 1
	case move == Paper && move2 == Rock:
		return 1
	case move == Paper && move2 == Scissors:
		return 2
	case move == Scissors && move2 == Paper:
		return 1
	case move == Scissors && move2 == Rock:
		return 2
	}

	// shouldn't reach here
	return -1
}

// NewGame creates a new game
func NewGame(in io.Reader, out io.Writer) Game {
	rand.Seed(time.Now().UnixNano())
	return Game{
		in:  in,
		out: out,
		getGuess: func() int {
			return rand.Intn(3)
		},
	}
}

// Game of Rock, Paper, Scissors.
type Game struct {
	in       io.Reader
	out      io.Writer
	getGuess func() int
}

// Run the game.
func (g Game) Run() {
	userChoice := g.GetChoice()
	computerChoice := Move(g.getGuess())
	fmt.Fprintf(g.out, "Computer chooses: %s\n", computerChoice)

	switch userChoice.Compare(computerChoice) {
	case 0:
		fmt.Fprintln(g.out, "It's a Tie!")
	case 1:
		fmt.Fprintln(g.out, "You WIN!!")
	case 2:
		fmt.Fprintln(g.out, "You LOSE!!!")
	}
}

// GetChoice of the player.
func (g Game) GetChoice() Move {
	var move Move
	for move < Scissors && move > Rock {
		fmt.Fprint(g.out, "Pick [r]ock, [p]aper, or [s]cissors: ")
		var input string
		fmt.Fscanln(g.in, &input)
		switch input {
		case "r", "rock":
			move = Rock
		case "p", "paper":
			move = Paper
		case "s", "scissors":
			move = Scissors
		default:
			fmt.Fprintln(g.out, "I didn't understand your choice, please retry")
		}
	}
	return move
}
