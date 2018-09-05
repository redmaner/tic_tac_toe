package main

const (
  version string = "v1.0"
  maxRounds int = 9
)

func main() {
  initGame()
}

// Initialize the game
func initGame() {

  // Initialize game data, which corresponds to the 9 cells of tic tac toe
  gameData := []string{
    "1", "2", "3",
    "4", "5", "6",
    "7", "8", "9",
  }

  for r := 1; r <= maxRounds; r++ {
    if r%2 == 0 {
      doTurn(gameData, r, "O", "PlayerTwo")
    } else {
      doTurn(gameData, r, "X", "PlayerOne")
    }
  }
}
