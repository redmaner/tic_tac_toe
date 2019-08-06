package main

import "fmt"

// Function that checks if there is a winner
func (g *game) checkWinner() (bool, *player) {

	// bi == boardIndex
	boardIndex := 1

	// Check horizontal win conditions
	for i := 0; i < g.boardSize; i++ {

		if i != 1 {
			boardIndex += g.boardSize
		}

		// br == boardRow == slice of strings with boardsize items
		boardRow := []string{}
		for cell := 0; cell < g.boardSize; cell++ {
			boardRow = append(boardRow, g.board[boardIndex+cell-1])
		}

		isWinner, winnerToken := g.checkIfWinCondition(boardRow)
		if isWinner {
			return isWinner, winnerToken
		}
	}

	// Check vertical win conditions
	for i := 0; i < g.boardSize; i++ {

		// bc == boardColumn == slice of strings with boardsize items
		boardColumn := []string{}
		boardColumnIndex := i
		for cell := 0; cell < g.boardSize; cell++ {
			if cell != 0 {
				boardColumnIndex += g.boardSize
			}
			boardColumn = append(boardColumn, g.board[boardColumnIndex])
		}

		isWinner, winnerToken := g.checkIfWinCondition(boardColumn)
		if isWinner {
			return isWinner, winnerToken
		}
	}

	// Check diagonal from left to right
	{
		boardDiagonalLeft := []string{}
		boardDiagonalLeftIndex := 0
		for cell := 0; cell < g.boardSize; cell++ {
			if cell != 0 {
				boardDiagonalLeftIndex += g.boardSize + 1
			}
			boardDiagonalLeft = append(boardDiagonalLeft, g.board[boardDiagonalLeftIndex])
		}

		isWinner, winnerToken := g.checkIfWinCondition(boardDiagonalLeft)
		if isWinner {
			return isWinner, winnerToken
		}
	}

	// Check diagonal from right to left
	{
		boardDiagonalRight := []string{}
		boardDiagonalRightIndex := g.boardSize - 1
		for cell := 0; cell < g.boardSize; cell++ {
			if cell != 0 {
				boardDiagonalRightIndex += g.boardSize - 1
			}
			boardDiagonalRight = append(boardDiagonalRight, g.board[boardDiagonalRightIndex])
		}

		isWinner, winnerToken := g.checkIfWinCondition(boardDiagonalRight)
		if isWinner {
			return isWinner, winnerToken
		}
	}

	return false, nil
}

func (g *game) checkIfWinCondition(cellSlice []string) (bool, *player) {
	var cellValue string
	var winSign string
	for cell := 0; cell < g.boardSize; cell++ {
		if cellValue == "" {
			cellValue = cellSlice[cell]
		} else if cellValue == cellSlice[cell] && cell == g.boardSize-1 {
			winSign = cellSlice[cell]
		} else if cellValue == cellSlice[cell] {
			continue
		} else {
			break
		}
	}

	switch {
	case g.players[0].sign == winSign:
		fmt.Println(cellSlice)
		fmt.Println(g.board)
		return true, g.players[0]
	case g.players[1].sign == winSign:
		fmt.Println(cellSlice)
		fmt.Println(g.board)
		return true, g.players[1]
	default:
		return false, nil
	}
}
