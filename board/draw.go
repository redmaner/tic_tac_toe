package board

import (
  "fmt"
)

// Function to show the board of tic tac toe
// This function is universal and is called every round
func Draw(gameData []string) {
  clearScreen()
  fmt.Println("------------------")
  fmt.Println("Tic Tac Go:", version)
  fmt.Println("------------------")
  fmt.Println("")
  fmt.Println("  ", gameData[1], "|", gameData[2], "|", gameData[3])
  fmt.Println("  -----------")
  fmt.Println("  ", gameData[4], "|", gameData[5], "|", gameData[6])
  fmt.Println("  -----------")
  fmt.Println("  ", gameData[7], "|", gameData[8], "|", gameData[9])
  fmt.Println("")
  fmt.Println("------------------")
}
