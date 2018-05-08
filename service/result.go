package service

import (
	"log"
)

type ProcessingResult struct {
	BestMatch [][]string `json:"bestMatch"`
}

func NewProcessingResult(h ProcessingHandler, numberOfPlayers, numberOfLines int) *ProcessingResult {
	bestPossibleGamesNumber := 0
	res := &ProcessingResult{
		BestMatch: make([][]string, 0),
	}
	printer := NewPrinter(numberOfPlayers, numberOfLines)
	for ; bestPossibleGamesNumber < len(h.PossibleGames) && h.PossibleGames[bestPossibleGamesNumber].Score >= h.PossibleGames[0].Score; bestPossibleGamesNumber++ {
		if bestPossibleGamesNumber < 5 {
			res.BestMatch = append(res.BestMatch, printer.GameAsArrayOfPlayers(*h.PossibleGames[bestPossibleGamesNumber]))
		}
	}
	log.Printf("Found %d best games in %d built and returning %d with score: %f\n", bestPossibleGamesNumber, len(h.PossibleGames), len(res.BestMatch), h.PossibleGames[0].Score)
	return res
}
