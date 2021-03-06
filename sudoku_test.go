package main

import (
	"fmt"
	"testing"
)

func TestDuplicated(t *testing.T) {
	if duplicated([10]int{0, 0, 0, 0, 0, 0, 0, 0, 0}) {
		t.Fatal("not duplicated but failed")
	}

	if duplicated([10]int{0, 0, 1, 0, 0, 0, 0, 0, 0}) {
		t.Fatal("not duplicated but failed")
	}

	if !duplicated([10]int{0, 2, 0, 0, 0, 0, 0, 0, 0}) {
		t.Fatal("it is duplicated")
	}
}

func TestVerify(t *testing.T) {
	cases := []struct {
		b        Board
		expected bool
	}{
		{
			b: Board{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: true,
		},
		{
			b: Board{
				{1, 0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: false,
		},
		{
			b: Board{
				{1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: false,
		},
		{
			b: Board{
				{1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: false,
		},
	}

	for i, v := range cases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			actual := verify(v.b)
			if actual != v.expected {
				t.Errorf("expeted %v, got %v", v.expected, actual)
			}
		})
	}
}

func TestBacktrack(t *testing.T) {
	b := Board{
		{0, 5, 0, 0, 8, 3, 0, 1, 7},
		{0, 0, 0, 1, 0, 0, 4, 0, 0},
		{3, 0, 4, 0, 0, 5, 6, 0, 8},

		{0, 0, 0, 0, 3, 0, 0, 0, 9},
		{0, 9, 0, 8, 2, 4, 5, 0, 0},
		{0, 0, 6, 0, 0, 0, 0, 7, 0},

		{0, 0, 9, 0, 0, 0, 0, 5, 0},
		{0, 0, 7, 2, 9, 0, 0, 8, 6},
		{1, 0, 3, 6, 0, 7, 2, 0, 4},
	}

	expected := Board{
		{6, 5, 2, 4, 8, 3, 9, 1, 7},
		{9, 7, 8, 1, 6, 2, 4, 3, 5},
		{3, 1, 4, 9, 7, 5, 6, 2, 8},

		{8, 2, 5, 7, 3, 6, 1, 4, 9},
		{7, 9, 1, 8, 2, 4, 5, 6, 3},
		{4, 3, 6, 5, 1, 9, 8, 7, 2},

		{2, 6, 9, 3, 4, 8, 7, 5, 1},
		{5, 4, 7, 2, 9, 1, 3, 8, 6},
		{1, 8, 3, 6, 5, 7, 2, 9, 4},
	}

	backtrack(&b, 1)
	// fmt.Printf("b = \n%+v\n", pretty(b))

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] != expected[i][j] {
				t.Errorf("expected %v in %d, %d, got %v", expected[i][j], i, j, b[i][j])
			}
		}
	}
}

func TestShort(t *testing.T) {
	t.Run("case Convertable", func(t *testing.T) {
		b, err := short("2.6.3......1.65.7..471.8.5.5......29..8.194.6...42...1....428..6.93....5.7.....13")
		if err != nil {
			t.Fatalf("Failed to parse %s", err)
		}

		expected := Board{
			{2, 0, 6, 0, 3, 0, 0, 0, 0},
			{0, 0, 1, 0, 6, 5, 0, 7, 0},
			{0, 4, 7, 1, 0, 8, 0, 5, 0},
			{5, 0, 0, 0, 0, 0, 0, 2, 9},
			{0, 0, 8, 0, 1, 9, 4, 0, 6},
			{0, 0, 0, 4, 2, 0, 0, 0, 1},
			{0, 0, 0, 0, 4, 2, 8, 0, 0},
			{6, 0, 9, 3, 0, 0, 0, 0, 5},
			{0, 7, 0, 0, 0, 0, 0, 1, 3},
		}

		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if b[i][j] != expected[i][j] {
					t.Errorf("expected %v in %d, %d, got %v", expected[i][j], i, j, b[i][j])
				}
			}
		}
	})
	t.Run("case Input too short", func(t *testing.T) {
		_, err := short("2.6.3....")
		if err == nil {
			t.Fatal("Expected to return error")
		}
	})
	t.Run("case Input too long", func(t *testing.T) {
		_, err := short("2.6.3......1.65.7..471.8.5.5.............................29..8.194.6...42...1....428..6.93....5.7.....13")
		if err == nil {
			t.Fatal("Expected to return error")
		}
	})
}
