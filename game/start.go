package game

import (
  "strconv"
)

// Initialize and run the game
func Start(boardSize int) {

  // Initialize game data, which corresponds to the cells of a tic tac toe game
  // gameData will be passed around throughout the program as the single source
  // of truth.
  // Because we are dealing with a slice we can pass it around freely without
  // pointers, because it is already a reference type
  gameData := initGameData(boardSize)

  // This starts the actual game of 9 rounds
  // PlayerOne plays the uneven rounds
  // PlayerTwo plays the even rounds
  maxRounds, _ := getData(gameData)
  for r := 1; r <= maxRounds; r++ {
    if r%2 == 0 {
      round(gameData, r, "O", "PlayerTwo")
    } else {
      round(gameData, r, "X", "PlayerOne")
    }
  }
}

func initGameData(boardSize int) []string {
  if boardSize < 3 {
    boardSize = 3
  }
  maxCells := boardSize * boardSize
  gameData := make([]string, maxCells + 1)
  gameData[0] = strconv.Itoa(boardSize)
  for c := 1; c <= maxCells; c++ {
    gameData[c] = strconv.Itoa(c)
  }
  return gameData
}
