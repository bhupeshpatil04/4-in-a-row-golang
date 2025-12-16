package bot

import "fourinarow/internal/game"

func BestMove(g *game.Game) int {
	// Try winning move
	for c := 0; c < game.Cols; c++ {
		copy := *g
		if copy.Drop(c) && copy.Winner == 2 {
			return c
		}
	}
	// Block opponent
	for c := 0; c < game.Cols; c++ {
		copy := *g
		copy.Turn = 1
		if copy.Drop(c) && copy.Winner == 1 {
			return c
		}
	}
	// Fallback
	for c := 0; c < game.Cols; c++ {
		if g.Drop(c) {
			g.Turn = 2
			return c
		}
	}
	return 0
}
