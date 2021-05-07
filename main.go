package main

import (
	"fmt"
	"io"
	"os"

	"github.com/jbarone/debugme/game"
)

const exitFail = 1

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, in io.Reader, out io.Writer) error {
	g := game.NewGame(in, out)
	g.Run()
	return nil
}
