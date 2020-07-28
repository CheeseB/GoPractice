package poker_test

import (
	poker "github.com/gypsydave5/learn-go-with-tests/command-line/v3"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record Sonia win from user input", func(t *testing.T) {
		in := strings.NewReader("Sonia wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Sonia")
	})

	t.Run("record Jimo win from user input", func(t *testing.T) {
		in := strings.NewReader("Jimo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Jimo")
	})
}
