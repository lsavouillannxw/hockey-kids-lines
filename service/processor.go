package service

import (
	"math/bits"
	"sort"
)

type ProcessingHandler struct {
	PossibleLinesAsArray []uint16
	PossibleGames        []*Game
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
	//fmt.Printf("%0.16b\n", maskOnes)
	//fmt.Printf("%0.16b\n", maskZeros)

	var possibleLinesAsMap = make(map[uint16]bool, 0)
	max := uint16(1)
	max = max << uintNumberOfPlayers
	min := uint16(1)
	min = min << uintLineSize
	for i := min; i < max; i++ {
		val := i & maskOnes
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
	return h
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
		if game.Lines[currentLine-1]&h.PossibleLinesAsArray[k] != 0 {
			continue
		}
		game.Lines[currentLine] = h.PossibleLinesAsArray[k]
		h.buildAllPossibleGames(game, currentLine+1)
	}
}
