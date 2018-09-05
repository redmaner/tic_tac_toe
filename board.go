package main

import (
  "os"
  "fmt"
)

func clearScreen() {
  for s := 1; s <= 5; s++ {
    fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
  }
}

func showBoard(gameData []string) {
  clearScreen()
  fmt.Println("------------------")
  fmt.Println("Tic Tac Go:", version)
  fmt.Println("------------------")
  fmt.Println("")
  fmt.Println("  ", gameData[0], "|", gameData[1], "|", gameData[2])
  fmt.Println("  -----------")
  fmt.Println("  ", gameData[3], "|", gameData[4], "|", gameData[5])
  fmt.Println("  -----------")
  fmt.Println("  ", gameData[6], "|", gameData[7], "|", gameData[8])
  fmt.Println("")
  fmt.Println("------------------")
}

func showWinner(gameData []string, winnerToken string) {
  clearScreen()
  showBoard(gameData)
  switch winnerToken {
  case "X":
    fmt.Println("\nPlayerOne is the winner!")
    os.Exit(0)
  case "O":
    fmt.Println("\nPlayerTwo is the winner!")
    os.Exit(0)
  case "null":
    fmt.Println("\nThere are no winners, only losers!")
    os.Exit(0)
  default:
    fmt.Println("\nThere are no winners, only losers!")
    os.Exit(0)
  }
}
