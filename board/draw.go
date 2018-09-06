package board

import (
  "fmt"
  "strconv"
)

// Function to show the board of tic tac toe
// This function is universal and is called every round
func Draw(gameData []string) {
  boardSize, _ := strconv.Atoi(gameData[0])
  maxCells := boardSize * boardSize

  // Draw the header
  clearScreen()
  fmt.Println("---------------------")
  fmt.Println("Tic Tac Go:", version)
  fmt.Println("---------------------")
  fmt.Print("\n  ")

  // Draw the board dynamicly, based on the size of the board (boardSize)
  for i := 1; i <= maxCells; i++ {
    fmt.Print(gameData[i])
    if i % boardSize == 0 && i != maxCells {
      printBoardLine(boardSize)
    } else if i % boardSize == 0 && i == maxCells {
      fmt.Print("\n")
    } else {
      fmt.Print(" | ")
    }
  }

  // Draw the bottom
  fmt.Println("")
  fmt.Println("------------------")
}


func printBoardLine(boardSize int) {
  fmt.Print("\n -")
  for c := 1; c <= boardSize; c++ {
    fmt.Print("---")
  }
  fmt.Print("-\n  ")
}
