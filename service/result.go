package service

import "log"

type ProcessingResult struct {
	BestMatch [][]string `json:"bestMatch"`
}

func NewProcessingResult(h ProcessingHandler) *ProcessingResult {
	bestPossibleGamesNumber := 0
	res := &ProcessingResult{
		BestMatch: make([][]string, 0),
	}
	printer := NewPrinter(h.NumberOfPlayers, h.NumberOfLines)
	for ; bestPossibleGamesNumber < len(h.PossibleGames) && h.PossibleGames[bestPossibleGamesNumber].Score >= h.PossibleGames[0].Score; bestPossibleGamesNumber++ {
		if bestPossibleGamesNumber < 20 {
			res.BestMatch = append(res.BestMatch, printer.GameAsArrayOfPlayers(*h.PossibleGames[bestPossibleGamesNumber]))
		}
	}
	if len(h.PossibleGames) == 0 {
		log.Printf("No best games found in %d built\n", len(h.PossibleGames))
	} else {
		log.Printf("Found %d best games in %d built and returning %d with score: %f\n", bestPossibleGamesNumber, len(h.PossibleGames), len(res.BestMatch), h.PossibleGames[0].Score)
	}
	return res
}
