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
  maxCellSize := len(strconv.Itoa(maxCells))

  // Draw the header
  clearScreen()
  fmt.Println("---------------------")
  fmt.Println("Tic Tac Go:", version)
  fmt.Println("---------------------")
  fmt.Print("\n  ")

  // Draw the board dynamicly, based on the size of the board (boardSize)
  for i := 1; i <= maxCells; i++ {
    cellSize := len(strconv.Itoa(i))
    fmt.Print(gameData[i])
    correctCellSize(maxCellSize, cellSize)
    if i % boardSize == 0 && i != maxCells {
      printBoardLine(boardSize, maxCellSize)
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


func printBoardLine(boardSize int, maxCellSize int) {
  fmt.Print("\n ")
  for c := 1; c <= boardSize; c++ {
    switch maxCellSize {
    case 1:
      fmt.Print("----")
    case 2:
      fmt.Print("-----")
    case 3:
      fmt.Print("------")
    }
  }
  fmt.Print("\n  ")
}

func correctCellSize(mcs int, cs int) {
  if mcs == cs {
    return
  } else {
    d := mcs - cs
    for i := 1; i <= d; i++ {
      fmt.Print(" ")
    }
  }
}
