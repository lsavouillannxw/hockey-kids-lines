package service

type Game struct {
	Players []uint16
	Lines   []uint16
	Score   float64
}

func NewGame(numberOfPlayers, numberOfLines int) *Game {
	return &Game{
		Players: make([]uint16, numberOfPlayers),
		Lines:   make([]uint16, numberOfLines),
	}
}

func (g Game) Clone() *Game {
	result := NewGame(len(g.Players), len(g.Lines))
	result.Score = g.Score
	copy(result.Lines, g.Lines)
	copy(result.Players, g.Players)
	return result
}

func (g Game) FillPlayersFromLines() {
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

func (g *Game) Evaluate() {
	g.Score = 0
	for i := 0; i < len(g.Lines); i++ {
		for j := i + 1; j < len(g.Lines); j++ {
			if g.Lines[i] == g.Lines[j] {
				g.Score += 10
				if i+2 == j {
					g.Score += 30
				}
			}
		}
	}
}
