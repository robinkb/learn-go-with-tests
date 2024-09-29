package poker

import (
	"bufio"
	"io"
	"strings"
)

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		store: store,
		stdin: bufio.NewScanner(in),
	}
}

type CLI struct {
	store PlayerStore
	stdin *bufio.Scanner
}

func (c *CLI) PlayPoker() {
	c.store.RecordWin(extractWinner(c.readLine()))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.stdin.Scan()
	return c.stdin.Text()
}
