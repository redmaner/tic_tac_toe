package main

// Function that checks if there is a winner
func (g *game) checkWinner() (bool, string) {

	// bi == boardIndex
	boardIndex := 1

	// Check horizontal win conditions
	for i := 1; i <= g.boardSize; i++ {

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
	for i := 1; i <= g.boardSize; i++ {

		// bc == boardColumn == slice of strings with boardsize items
		boardColumn := []string{}
		boardColumnIndex := i
		for cell := 0; cell < g.boardSize; cell++ {
			if cell != 0 {
				boardColumnIndex += g.boardSize
			}
			boardColumn = append(boardColumn, g.board[boardColumnIndex-1])
		}

		isWinner, winnerToken := g.checkIfWinCondition(boardColumn)
		if isWinner {
			return isWinner, winnerToken
		}
	}

	// Check diagonal from left to right
	{
		boardDiagonal := []string{}
		boardDiagonalIndex := 1
		for cell := 0; cell < g.boardSize; cell++ {
			if cell != 0 {
				boardDiagonalIndex += g.boardSize + 1
			}
			boardDiagonal = append(boardDiagonal, g.board[boardDiagonalIndex-1])
		}

		isWinner, winnerToken := g.checkIfWinCondition(boardDiagonal)
		if isWinner {
			return isWinner, winnerToken
		}
	}

	// Check diagonal from right to left
	{
		boardDiagonal := []string{}
		boardDiagonalIndex := 1
		for cell := 0; cell < g.boardSize; cell++ {
			if cell != 0 {
				boardDiagonalIndex += g.boardSize - 1
			}
			boardDiagonal = append(boardDiagonal, g.board[boardDiagonalIndex-1])
		}

		isWinner, winnerToken := g.checkIfWinCondition(boardDiagonal)
		if isWinner {
			return isWinner, winnerToken
		}
	}

	return false, "null"
}

func (g *game) checkIfWinCondition(cellSlice []string) (bool, string) {
	var cellValue string
	for cell := 0; cell < g.boardSize; cell++ {
		if cellValue == "" {
			cellValue = cellSlice[cell]
		} else if cellValue == cellSlice[cell] && cell == g.boardSize-1 {
			return true, cellSlice[cell]
		} else if cellValue == cellSlice[cell] {
			continue
		} else {
			break
		}
	}
	return false, "null"
}
