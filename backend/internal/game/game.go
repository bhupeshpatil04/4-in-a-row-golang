package game

const (
	Rows = 6
	Cols = 7
)

type Game struct {
	Board  [Rows][Cols]int
	Turn   int
	Winner int
}

func NewGame() *Game {
	return &Game{Turn: 1}
}

func (g *Game) Drop(col int) bool {
	if col < 0 || col >= Cols || g.Winner != 0 {
		return false
	}
	for r := Rows - 1; r >= 0; r-- {
		if g.Board[r][col] == 0 {
			g.Board[r][col] = g.Turn
			if g.checkWin(r, col) {
				g.Winner = g.Turn
			}
			g.Turn = 3 - g.Turn
			return true
		}
	}
	return false
}

func (g *Game) checkWin(r, c int) bool {
	dirs := [][]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}
	for _, d := range dirs {
		count := 1
		for i := 1; i < 4; i++ {
			nr, nc := r+d[0]*i, c+d[1]*i
			if nr < 0 || nr >= Rows || nc < 0 || nc >= Cols || g.Board[nr][nc] != g.Board[r][c] {
				break
			}
			count++
		}
		for i := 1; i < 4; i++ {
			nr, nc := r-d[0]*i, c-d[1]*i
			if nr < 0 || nr >= Rows || nc < 0 || nc >= Cols || g.Board[nr][nc] != g.Board[r][c] {
				break
			}
			count++
		}
		if count >= 4 {
			return true
		}
	}
	return false
}
