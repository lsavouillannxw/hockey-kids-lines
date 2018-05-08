package service

import (
	"log"
	"math/bits"
	"sort"
)

type ProcessingHandler struct {
	// Input
	NumberOfPlayers int
	NumberOfLines   int
	LineSize        int
	// Computed
	PossibleLinesAsArray []uint16
	AppearancesCounter   [][]int
	MaxAppearances       int
	// Output
	PossibleGames []*Game
	MaxScore      float64
}

func NewProcessingHandler(numberOfPlayers, numberOfLines, lineSize int) *ProcessingHandler {
	return &ProcessingHandler{
		NumberOfPlayers:      numberOfPlayers,
		NumberOfLines:        numberOfLines,
		LineSize:             lineSize,
		PossibleLinesAsArray: make([]uint16, 0),
		PossibleGames:        make([]*Game, 0),
	}
}

func (h *ProcessingHandler) Process() *ProcessingHandler {
	uintNumberOfPlayers := uint(h.NumberOfPlayers)
	uintLineSize := uint(h.LineSize)
	intLineSize := int(uintLineSize)

	maskOnes := uint16(1) << uintNumberOfPlayers
	maskOnes--
	//log.Printf("%0.16b\n", maskOnes)
	//log.Printf("%0.16b\n", maskZeros)

	h.MaxAppearances = h.LineSize * h.NumberOfLines / h.NumberOfPlayers
	if h.LineSize*h.NumberOfLines%h.NumberOfPlayers > 0 {
		h.MaxAppearances++
	}
	log.Printf("Max appearance: %d", h.MaxAppearances)
	h.AppearancesCounter = make([][]int, h.NumberOfLines)
	for i := 0; i < h.NumberOfLines; i++ {
		h.AppearancesCounter[i] = make([]int, h.NumberOfPlayers)
	}

	var possibleLinesAsMap = make(map[uint16]bool, 0)
	max := uint16(1)
	max = max << uintNumberOfPlayers
	min := uint16(1)
	min = min << uintLineSize
	for i := min; i < max; i++ {
		val := i & maskOnes
		if bits.OnesCount16(val) == intLineSize {
			possibleLinesAsMap[val] = true
			//log.Printf("%0.16b\n", val)
		}
	}
	log.Printf("Number of possible lines: %d\n", len(possibleLinesAsMap))
	for k := range possibleLinesAsMap {
		h.PossibleLinesAsArray = append(h.PossibleLinesAsArray, k)
	}

	game := NewGame(h.NumberOfPlayers, h.NumberOfLines)
	game.Lines[0] = (uint16(1) << uintLineSize) - 1
	for i := 0; i < h.LineSize; i++ {
		h.AppearancesCounter[0][i]++
		h.AppearancesCounter[1][i]++
	}
	game.Lines[1] = ((uint16(1) << uintLineSize) - 1) << (uintNumberOfPlayers - uintLineSize)
	for i := 0; i < h.LineSize; i++ {
		h.AppearancesCounter[1][h.NumberOfPlayers-1-i]++
	}
	h.buildAllPossibleGames(game, 2)
	sort.Slice(h.PossibleGames, func(i, j int) bool { return h.PossibleGames[i].Score > h.PossibleGames[j].Score })
	return h
}

func (h *ProcessingHandler) buildAllPossibleGames(game *Game, currentLine int) {
	if currentLine == h.NumberOfLines {
		game.Evaluate()
		if game.Score >= h.MaxScore {
			h.MaxScore = game.Score
			h.PossibleGames = append(h.PossibleGames, game.Clone())
		}
		return
	}

	for k := range h.PossibleLinesAsArray {
		if game.Lines[currentLine-1]&h.PossibleLinesAsArray[k] != 0 {
			continue // A player can't play twice in a row
		}
		cpt := uint16(1)
		invalid := false
		for i := 0; i < h.NumberOfPlayers; i++ {
			if cpt&h.PossibleLinesAsArray[k] != 0 {
				h.AppearancesCounter[currentLine][i] = h.AppearancesCounter[currentLine-1][i] + 1
				if h.AppearancesCounter[currentLine][i] > h.MaxAppearances {
					invalid = true
					break // A player can't play more than MaxAppearances
				}
			} else {
				h.AppearancesCounter[currentLine][i] = h.AppearancesCounter[currentLine-1][i]
				if h.AppearancesCounter[currentLine][i]+((h.NumberOfLines-currentLine)/2) < h.MaxAppearances-1 {
					invalid = true
					break // A player must play at least MaxAppearances - 1
				}
			}
			cpt = cpt * 2
		}
		if invalid {
			continue
		}
		game.Lines[currentLine] = h.PossibleLinesAsArray[k]
		h.buildAllPossibleGames(game, currentLine+1)
	}
}
