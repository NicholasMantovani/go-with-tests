package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	store   PlayerStore
	in      *bufio.Scanner
	out     io.Writer
	alerter BlindAlerter
}

func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{store: store,
		in:      bufio.NewScanner(in),
		out:     out,
		alerter: alerter}
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)

	numberOfPlayers, _ := strconv.Atoi(c.readLine())
	c.scheduleBlindAlerts(numberOfPlayers)
	userInput := c.readLine()
	c.store.RecordWin(extractWinner(userInput))
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func (c *CLI) scheduleBlindAlerts(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	duration := 0 * time.Second
	for _, b := range blinds {
		c.alerter.ScheduleAlertAt(duration, b)
		duration = duration + blindIncrement
	}
}

func extractWinner(userInpt string) string {
	return strings.Replace(userInpt, " wins", "", 1)
}
