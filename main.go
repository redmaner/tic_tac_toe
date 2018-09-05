package main


// maxRounds is the amount of rounds of a tic tac toe game
const (
  version string = "v2.0 alpha1"
  maxRounds int = 9
)

func main() {
  initGame()
}

// Initialize and run the game
func initGame() {

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
  for r := 1; r <= maxRounds; r++ {
    if r%2 == 0 {
      doTurn(gameData, r, "O", "PlayerTwo")
    } else {
      doTurn(gameData, r, "X", "PlayerOne")
    }
  }
}
