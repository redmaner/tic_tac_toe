package game

// Initialize and run the game
func Start() {

  // Initialize game data, which corresponds to the 9 cells of a tic tac toe game
  // gameData will be passed around throughout the program as the single source
  // of truth.
  // Because we are dealing with a slice we can pass it around freely without
  // pointers, because it is already a reference type
  gameData := []string{
    "3",
    "1", "2", "3",
    "4", "5", "6",
    "7", "8", "9",
  }

  // This starts the actual game of 9 rounds
  // PlayerOne plays the uneven rounds
  // PlayerTwo plays the even rounds
  maxRounds := getMaxRounds(gameData)
  for r := 1; r <= maxRounds; r++ {
    if r%2 == 0 {
      round(gameData, r, "O", "PlayerTwo")
    } else {
      round(gameData, r, "X", "PlayerOne")
    }
  }
}
