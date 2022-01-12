package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	// "time"
)

type Board [9][9]int

func pretty(b Board) string {
	var buf bytes.Buffer
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString("+---+---+---+\n")
		}

		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString("|")
			}
			buf.WriteString(strconv.Itoa(b[i][j]))
		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+---+---+---+\n")
	return buf.String()
}

func duplicated(c [10]int) bool {
	for i, v := range c {
		if i == 0 {
			continue
		}
		if v >= 2 {
			return true
		}
	}
	return false
}

func verify(b Board) bool {
	// check row and col
	for i := 0; i < 9; i++ {
		var row [10]int
		var col [10]int
		for j := 0; j < 9; j++ {
			row[b[i][j]]++
			col[b[j][i]]++
		}
		if duplicated(row) || duplicated(col) {
			return false
		}
	}

	// check inner box
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var c [10]int
			for row := i; row < 3; row++ {
				for col := j; col < 3; col++ {
					c[b[row][col]]++
				}
			}
			if duplicated(c) {
				return false
			}
		}
	}

	return true
}

func solved(b Board) bool {
	if !verify(b) {
		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func backtrack(b *Board, depth int) bool {
	// fmt.Printf("Depth: %v\n", depth)
	if solved(*b) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				for c := 9; c >= 1; c-- {
					b[i][j] = c
					// time.Sleep(time.Second * 1)
					// fmt.Printf("Depth: %v: Board[%v][%v]\n", depth, i, j)
					// fmt.Printf("%v\n", pretty(*b))
					if verify(*b) {
						if backtrack(b, depth+1) {
							return true
						}
					}
					b[i][j] = 0
				}
				return false
			}
		}
	}
	return false
}

func short(input string) (*Board, error) {
	if len(input) != 81 {
		return nil, errors.New("Input length must be 81")
	}
	sc := bufio.NewScanner(strings.NewReader(input))
	sc.Split(bufio.ScanRunes)
	var b Board

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !sc.Scan() {
				break
			}

			t := sc.Text()
			if t == "." {
				b[i][j] = 0
			} else {
				n, err := strconv.Atoi(t)
				if err != nil {
					return nil, err
				}
				b[i][j] = n
			}
		}
	}
	return &b, nil
}

func main() {
	b := Board{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Printf("b = %+v\n", pretty(b))
	return
}
