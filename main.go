package main

import (
	"fmt"
	"math/bits"
	"math"
	"sort"
	"net/http"
	"bytes"
	"encoding/json"
	"strings"
	"strconv"
)

func init() {
	http.HandleFunc("/", handler)
}

type result struct {
	BestMatch [][]string `json:"bestMatch"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/api") {
		http.Redirect(w, r, strings.Split(r.URL.String(), r.URL.Path)[0] + "/web/hockeyKidsLinesPage.html", http.StatusPermanentRedirect)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}

	numberOfPlayers, err := strconv.Atoi(r.URL.Query().Get("numberOfPlayers"))
	if err != nil {
		w.WriteHeader(400)
	}
	numberOfPlayersPerLine, err := strconv.Atoi(r.URL.Query().Get("numberOfPlayersPerLine"))
	if err != nil {
		w.WriteHeader(400)
	}
	numberOfLinesPerMatch, err := strconv.Atoi(r.URL.Query().Get("numberOfLinesPerMatch"))
	if err != nil {
		w.WriteHeader(400)
	}
	h := &processingHandler{
		PossibleLinesAsArray: make([]uint16, 0),
		PossibleGames: make([]*game, 0),
	}
	if numberOfPlayers > 16 || numberOfPlayers < 7 || numberOfPlayersPerLine < 3 || numberOfPlayersPerLine > 5 || numberOfLinesPerMatch < 5 || numberOfLinesPerMatch > 16 || numberOfPlayers % numberOfPlayersPerLine == 0 {
		w.WriteHeader(400)
		return
	}
	res := h.process(numberOfPlayers, numberOfLinesPerMatch, numberOfPlayersPerLine)
	resAsBytes, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(resAsBytes))
}

//func main() {
//	fmt.Println((&processingHandler{
//		PossibleLinesAsArray: make([]uint16, 0),
//		PossibleGames: make([]*game, 0),
//	}).process(9, 10, 4))
//}

func (h *processingHandler) process(numberOfPlayers, numberOfLines, lineSize int) result {
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

	game := createGame(uintNumberOfPlayers, uintNumberOfLines)
	game.Lines[0] = (uint16(1) << uintLineSize) - 1
	game.Lines[1] = ((uint16(1) << uintLineSize) - 1) << uintLineSize
	h.buildGame(&game, 2)
	//game.displayLines()
	//game.fillPlayersFromLines()
	//game.displayPlayers()
	sort.Slice(h.PossibleGames, func(i, j int) bool { return h.PossibleGames[i].Score > h.PossibleGames[j].Score })
	bestPossibleGamesNumber := 0
	res := result{
		BestMatch: make([][]string, 0),
	}
	for ; bestPossibleGamesNumber < 5 && bestPossibleGamesNumber < len(h.PossibleGames) && h.PossibleGames[bestPossibleGamesNumber].Score >= h.PossibleGames[0].Score; bestPossibleGamesNumber++ {
		res.BestMatch = append(res.BestMatch, h.gameAsArrayOfPlayers(*h.PossibleGames[bestPossibleGamesNumber]))
	}
	fmt.Printf("Found %d best games\n", bestPossibleGamesNumber)
	return res
}

func (h *processingHandler) buildGame(game *game, currentLine int) {
	if currentLine == len(game.Lines) {
		game.fillPlayersFromLines()
		game.evaluate()
		if game.Score < 10 {
			//fmt.Printf("Score too bad: %f\n", game.Score)
		} else {
			//game.displayPlayers()
			h.PossibleGames = append(h.PossibleGames, game.copy())
		}
		return
	}

	for k := range h.PossibleLinesAsArray {
		if game.Lines[currentLine - 1] & h.PossibleLinesAsArray[k] != 0 {
			continue
		}
		game.Lines[currentLine] = h.PossibleLinesAsArray[k]
		h.buildGame(game, currentLine +1)
	}
}

type processingHandler struct {
	PossibleLinesAsArray []uint16
	PlayerFormat string
	LineFormat string
	MaskOnes uint16
	PossibleGames []*game
}

type game struct {
	Lines []uint16
	Players []uint16
	Score float64
}

func (g game) copy() *game {
	result := &game{
		Lines: make([]uint16, len(g.Lines)),
		Players: make([]uint16, len(g.Players)),
		Score: g.Score,
	}
	copy(result.Lines, g.Lines)
	copy(result.Players, g.Players)
	return result
}

func (g game) fillPlayersFromLines() {
	for i := 0; i < len(g.Players); i++ {
		g.Players[i] = 0
	}
	playerMaskOne := uint16(1)
	for i := 0; i < len(g.Lines); i++ {
		lineMaskOne := uint16(1)
		for j := 0; j < len(g.Players); j++ {
			if lineMaskOne&g.Lines[i] != 0 {
				g.Players[j] += playerMaskOne
			}
			lineMaskOne = lineMaskOne << 1
		}
		playerMaskOne = playerMaskOne << 1
	}
}

func (g *game) evaluate() {
	min := len(g.Lines)
	max := 0
	for i := 0; i < len(g.Players); i++ {
		cpt := bits.OnesCount16(g.Players[i])
		if cpt < min {
			min = cpt
		}
		if cpt > max {
			max = cpt
		}
	}
	g.Score = -math.Pow10(max - min) * 2

	for i := 0; i < len(g.Lines); i++ {
		for j := i + 1; j < len(g.Lines); j++ {
			if g.Lines[i] == g.Lines[j] {
				g.Score += 10
				if i + 2 == j {
					g.Score += 30
				}
			}
		}
	}
}

func createGame(uintNumberOfPlayers, numberOfLines uint) game {
	return game{
		Lines: make([]uint16, numberOfLines),
		Players: make([]uint16, uintNumberOfPlayers),
	}
}

func (h *processingHandler) displayLines(g game) {
	for i := 0; i < len(g.Lines); i++ {
		h.displayLine(g.Lines[i])
	}
	fmt.Println()
}

func (h *processingHandler) displayLine(line uint16) string {
	return fmt.Sprintf(h.LineFormat, line)
}

func (h processingHandler) displayPlayers(g game) string {
	//fmt.Printf("Game score: %f\n", g.Score)
	result := ""
	for i := 0; i < len(g.Players); i++ {
		result += h.displayPlayer(g.Players[i])
	}
	fmt.Println()
	return result
}

func (h processingHandler) writePlayers(builder *bytes.Buffer, g game) {
	for i := 0; i < len(g.Players); i++ {
		builder.WriteString(h.displayPlayer(g.Players[i]))
	}
	builder.WriteString("\n")
}

func (h processingHandler) displayPlayer(player uint16) string {
	return fmt.Sprintf(h.PlayerFormat, player)
}

func (h processingHandler) gameAsArrayOfPlayers(g game) []string {
	res := make([]string, 0)
	for i := 0; i < len(g.Players); i++ {
		res = append(res, h.displayPlayer(g.Players[i]))
	}
	return res
}