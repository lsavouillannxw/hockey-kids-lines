package service

import (
	"fmt"
	"math/bits"
	"sort"
	"bytes"
)

type ProcessingResult struct {
	BestMatch [][]string `json:"bestMatch"`
}

type ProcessingHandler struct {
	PossibleLinesAsArray []uint16
	PlayerFormat string
	LineFormat string
	MaskOnes uint16
	PossibleGames []*Game
}

func NewProcessingHandler() *ProcessingHandler {
	return &ProcessingHandler {
	PossibleLinesAsArray: make([]uint16, 0),
		PossibleGames: make([]*Game, 0),
	}
}

func (h *ProcessingHandler) Process(numberOfPlayers, numberOfLines, lineSize int) ProcessingResult {
	uintNumberOfPlayers := uint(numberOfPlayers)
	uintNumberOfLines := uint(numberOfLines)
	uintLineSize := uint(lineSize)
	intLineSize := int(uintLineSize)

	h.LineFormat = fmt.Sprintf("%%0%db", uintNumberOfPlayers)
	h.PlayerFormat = fmt.Sprintf("%%0%db", uintNumberOfLines)

	h.MaskOnes = uint16(1) << uintNumberOfPlayers
	h.MaskOnes--
	//fmt.Printf("%0.16b\n", maskOnes)
	//fmt.Printf("%0.16b\n", maskZeros)

	var possibleLinesAsMap = make(map[uint16]bool, 0)
	max := uint16(1)
	max = max << uintNumberOfPlayers
	min := uint16(1)
	min = min << uintLineSize
	for i := min; i < max; i++ {
		val := i & h.MaskOnes
		if bits.OnesCount16(val) == intLineSize {
			possibleLinesAsMap[val] = true
			//fmt.Printf("%0.16b\n", val)
		}
	}
	//fmt.Printf("Number of possible lines: %d\n", len(possibleLines))
	for k := range possibleLinesAsMap {
		h.PossibleLinesAsArray = append(h.PossibleLinesAsArray, k)
	}

	game := NewGame(numberOfPlayers, numberOfLines)
	game.Lines[0] = (uint16(1) << uintLineSize) - 1
	game.Lines[1] = ((uint16(1) << uintLineSize) - 1) << (uintNumberOfPlayers - uintLineSize)
	h.buildAllPossibleGames(game, 2)
	sort.Slice(h.PossibleGames, func(i, j int) bool { return h.PossibleGames[i].Score > h.PossibleGames[j].Score })
	bestPossibleGamesNumber := 0
	res := ProcessingResult{
		BestMatch: make([][]string, 0),
	}
	for ; bestPossibleGamesNumber < 5 && bestPossibleGamesNumber < len(h.PossibleGames) && h.PossibleGames[bestPossibleGamesNumber].Score >= h.PossibleGames[0].Score; bestPossibleGamesNumber++ {
		res.BestMatch = append(res.BestMatch, h.gameAsArrayOfPlayers(*h.PossibleGames[bestPossibleGamesNumber]))
	}
	fmt.Printf("Found %d best games\n", bestPossibleGamesNumber)
	return res
}

func (h *ProcessingHandler) buildAllPossibleGames(game *Game, currentLine int) {
	if currentLine == len(game.Lines) {
		game.FillPlayersFromLines()
		game.Evaluate()
		if game.Score < 10 {
			//fmt.Printf("Score too bad: %f\n", Game.Score)
		} else {
			//Game.displayPlayers()
			h.PossibleGames = append(h.PossibleGames, game.clone())
		}
		return
	}

	for k := range h.PossibleLinesAsArray {
		if game.Lines[currentLine - 1] & h.PossibleLinesAsArray[k] != 0 {
			continue
		}
		game.Lines[currentLine] = h.PossibleLinesAsArray[k]
		h.buildAllPossibleGames(game, currentLine +1)
	}
}

func (h *ProcessingHandler) displayLines(g Game) {
	for i := 0; i < len(g.Lines); i++ {
		h.displayLine(g.Lines[i])
	}
}

func (h *ProcessingHandler) displayLine(line uint16) string {
	return fmt.Sprintf(h.LineFormat, line)
}

func (h ProcessingHandler) displayPlayers(g Game) string {
	//fmt.Printf("Game score: %f\n", g.Score)
	result := ""
	for i := 0; i < len(g.Players); i++ {
		result += h.displayPlayer(g.Players[i])
	}
	return result
}

func (h ProcessingHandler) writePlayers(builder *bytes.Buffer, g Game) {
	for i := 0; i < len(g.Players); i++ {
		builder.WriteString(h.displayPlayer(g.Players[i]))
	}
	builder.WriteString("\n")
}

func (h ProcessingHandler) displayPlayer(player uint16) string {
	return fmt.Sprintf(h.PlayerFormat, player)
}

func (h ProcessingHandler) gameAsArrayOfPlayers(g Game) []string {
	res := make([]string, 0)
	for i := 0; i < len(g.Players); i++ {
		res = append(res, h.displayPlayer(g.Players[i]))
	}
	return res
}