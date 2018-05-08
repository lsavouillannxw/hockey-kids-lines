package service

import (
	"log"
	"math/bits"
	"sort"
)

type ProcessingHandler struct {
	PossibleLinesAsArray []uint16
	PossibleGames        []*Game
	AppearancesCounter   [][]int
	MaxAppearances       int
	MaxScore             float64
}

func NewProcessingHandler() *ProcessingHandler {
	return &ProcessingHandler{
		PossibleLinesAsArray: make([]uint16, 0),
		PossibleGames:        make([]*Game, 0),
	}
}

func (h *ProcessingHandler) Process(numberOfPlayers, numberOfLines, lineSize int) *ProcessingHandler {
	uintNumberOfPlayers := uint(numberOfPlayers)
	uintLineSize := uint(lineSize)
	intLineSize := int(uintLineSize)

	maskOnes := uint16(1) << uintNumberOfPlayers
	maskOnes--
	//log.Printf("%0.16b\n", maskOnes)
	//log.Printf("%0.16b\n", maskZeros)

	h.MaxAppearances = lineSize * numberOfLines / numberOfPlayers
	if lineSize*numberOfLines%numberOfPlayers > 0 {
		h.MaxAppearances++
	}
	log.Printf("Max appearance: %d", h.MaxAppearances)
	h.AppearancesCounter = make([][]int, numberOfLines)
	for i := 0; i < numberOfLines; i++ {
		h.AppearancesCounter[i] = make([]int, numberOfPlayers)
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

	game := NewGame(numberOfPlayers, numberOfLines)
	game.Lines[0] = (uint16(1) << uintLineSize) - 1
	for i := 0; i < lineSize; i++ {
		h.AppearancesCounter[0][i]++
		h.AppearancesCounter[1][i]++
	}
	game.Lines[1] = ((uint16(1) << uintLineSize) - 1) << (uintNumberOfPlayers - uintLineSize)
	for i := 0; i < lineSize; i++ {
		h.AppearancesCounter[1][numberOfPlayers - 1 - i]++
	}
	h.buildAllPossibleGames(game, 2)
	sort.Slice(h.PossibleGames, func(i, j int) bool { return h.PossibleGames[i].Score > h.PossibleGames[j].Score })
	return h
}

func (h *ProcessingHandler) buildAllPossibleGames(game *Game, currentLine int) {
	if currentLine == len(game.Lines) {
		//game.FillPlayersFromLines()
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
		for i := 0; i < len(h.AppearancesCounter[currentLine]); i++ {
			if cpt&h.PossibleLinesAsArray[k] != 0 {
				h.AppearancesCounter[currentLine][i] = h.AppearancesCounter[currentLine - 1][i] + 1
				if h.AppearancesCounter[currentLine][i] > h.MaxAppearances {
					invalid = true
					break // A player can't play more than MaxAppearances
				}
			} else {
				h.AppearancesCounter[currentLine][i] = h.AppearancesCounter[currentLine - 1][i]
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
