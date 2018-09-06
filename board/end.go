package board

import (
  "fmt"
)

// Function that is shown at the end of the game
// The game can either end with a winner or end in a draw
func End(gameData []string, winnerToken string) {
  clearScreen()
  Draw(gameData)
  switch winnerToken {
  case "X":
    fmt.Println("\nPlayerOne is the winner!")
  case "O":
    fmt.Println("\nPlayerTwo is the winner!")
  case "null":
    fmt.Println("\nThere are no winners, only losers!")
  default:
    fmt.Println("\nThere are no winners, only losers!")
  }
}
