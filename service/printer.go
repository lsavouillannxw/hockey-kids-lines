package service

import (
	"bytes"
	"fmt"
)

type Printer struct {
	PlayerFormat string
	LineFormat   string
}

func NewPrinter(numberOfPlayers, numberOfLines int) *Printer {
	return &Printer{
		PlayerFormat: fmt.Sprintf("%%0%db", numberOfLines),
		LineFormat:   fmt.Sprintf("%%0%db", numberOfPlayers),
	}
}

func (h Printer) displayPlayers(g Game) string {
	//fmt.Printf("Game score: %f\n", g.Score)
	result := ""
	for i := 0; i < len(g.Players); i++ {
		result += h.displayPlayer(g.Players[i])
	}
	return result
}

func (h Printer) writePlayers(builder *bytes.Buffer, g Game) {
	for i := 0; i < len(g.Players); i++ {
		builder.WriteString(h.displayPlayer(g.Players[i]))
	}
	builder.WriteString("\n")
}

func (h Printer) displayPlayer(player uint16) string {
	return fmt.Sprintf(h.PlayerFormat, player)
}

func (h Printer) GameAsArrayOfPlayers(g Game) []string {
	res := make([]string, 0)
	for i := 0; i < len(g.Players); i++ {
		res = append(res, h.displayPlayer(g.Players[i]))
	}
	return res
}
