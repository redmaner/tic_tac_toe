package main

import (
  "fmt"
  "os"
  "os/exec"
  "runtime"
)

// clearScreen is a cross platform function to clear the screen
func clearScreen() {

  switch runtime.GOOS {
  case "linux":
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
  case "windows":
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
  default:
    for s := 1; s <= 5; s++ {
      fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
    }
  }
}

// Function to show the board of tic tac toe
// This function is universal and is called every round
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

// Function that is shown at the end of the game
// The game can either end with a winner or end in a draw
func showWinner(gameData []string, winnerToken string) {
  clearScreen()
  showBoard(gameData)
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
