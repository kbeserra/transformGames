package reelGames

import (
	"fmt"
	"strings"

	"github.com/kbeserra/transformGames/representation"
)

type Board struct {
	Symbols [][]string
}

const (
	BoardSymbolStringMaxLength = 8
	BoardEmptyCellSymbol       = "#"
	BoardCellRow               = 1
	BoardCellColumn            = 0
	BoardEmptyCellString       = ""
)

func (b Board) String() string {
	var sb strings.Builder

	n := 0
	for _, col := range b.Symbols {
		if len(col) > n {
			n = len(col)
		}
	}

	var err error
	for i := 0; i < n; i++ {
		for j, col := range b.Symbols {
			if i < len(col) {
				if len(col[i]) < BoardSymbolStringMaxLength {
					_, err = sb.WriteString(strings.Repeat(" ", BoardSymbolStringMaxLength-len(col[i])))
				}
				if err != nil {
					panic(err)
				}
				if len(col[i]) < BoardSymbolStringMaxLength {
					_, err = sb.WriteString(col[i])
				} else {
					_, err = sb.WriteString(col[i][:BoardSymbolStringMaxLength])
				}
			} else {
				_, err = sb.WriteString(strings.Repeat(BoardEmptyCellSymbol, BoardSymbolStringMaxLength))
			}
			if err != nil {
				panic(err)
			}
			if j+1 < len(b.Symbols) {
				_, err = sb.WriteString("\t")
			}
			if err != nil {
				panic(err)
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (b Board) Copy() representation.GameState {
	symbs := make([][]string, len(b.Symbols))
	for i, col := range b.Symbols {
		symbs[i] = make([]string, len(col))
		copy(symbs[i], col)
	}
	return &Board{
		Symbols: symbs,
	}
}

func (b *Board) FromShape(shape []uint) error {
	b.Symbols = make([][]string, len(shape))
	for i, n := range shape {
		b.Symbols[i] = make([]string, n)
	}
	return nil
}

func (b *Board) FillFromReels(reels [][]string, stops []int) error {
	if len(stops) < len(b.Symbols) {
		return fmt.Errorf("length of stops too short for board")
	}
	if len(reels) < len(b.Symbols) {
		return fmt.Errorf("length of stops too short for reels")
	}

	for j, col := range b.Symbols {
		reel := reels[j]
		if len(reel) == 0 && len(col) != 0 {
			return fmt.Errorf("reel %d is empty while column %d is not", j, j)
		}
		s := stops[j] % len(reel)
		for i := range col {
			col[i] = reel[s]
			s = (s + 1) % len(reel)
		}
	}

	return nil
}

func (b *Board) RemoveCells(cells [][2]int) error {
	for _, c := range cells {
		row, col := c[BoardCellRow], c[BoardCellColumn]
		if col >= len(b.Symbols) {
			return fmt.Errorf("cell column index, %d, is out of range", col)
		}
		if row >= len(b.Symbols[col]) {
			return fmt.Errorf("cell row index, %d, is out of range", row)
		}
		b.Symbols[col][row] = BoardEmptyCellString
	}
	return nil
}

func (b *Board) ProjectToLine(rows []int) ([]string, error) {
	n := len(rows)
	if n > len(b.Symbols) {
		n = len(b.Symbols)
	}
	rtn := make([]string, n)
	for i := 0; i < n; i++ {
		rtn[i] = b.Symbols[i][rows[i]]
	}
	return rtn, nil
}

func (b *Board) RemoveRow(row int) error {
	for r, col := range b.Symbols {
		if row+1 == len(col) {
			b.Symbols[r] = col[:row]
		} else if row+1 < len(col) {
			b.Symbols[r] = append(col[:row], col[row+1:]...)
		}
	}

	return nil
}
