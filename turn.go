package main

import (
  "fmt"
  "os"
)

func doTurn(gameData []string, round int, playerToken string, playerName string) {
  showBoard(gameData)
  fmt.Println("Round:", round)
  fmt.Println("Player:", playerName)
  fmt.Println("Token:", playerToken)
  fmt.Println("------------------")
  fmt.Println("")

  getPosition(gameData, playerToken)

  if round >= 5 {
    isWinner, winnerToken := checkWinner(gameData)
    if isWinner {
      showWinner(gameData, winnerToken)
      playAgain()
    } else if !isWinner && round == 9 {
      showWinner(gameData, "null")
      playAgain()
    }
  }
}

func checkWinner(gameData []string) (bool, string) {
  switch {
  // Cell 1, 2, and 3 are equal
  case gameData[0] == gameData[1] && gameData[1] == gameData[2]:
    return true, gameData[1]

  // Cell 4, 5, and 6 are equal
  case gameData[3] == gameData[4] && gameData[4] == gameData[5]:
    return true, gameData[4]

  // Cell 7, 8, and 9 are equal
  case gameData[6] == gameData[7] && gameData[7] == gameData[8]:
    return true, gameData[7]

  // Cell 1, 4, and 7 are equal
  case gameData[0] == gameData[3] && gameData[3] == gameData[6]:
    return true, gameData[3]

  // Cell 2, 5, and 8 are equal
  case gameData[1] == gameData[4] && gameData[4] == gameData[7]:
    return true, gameData[4]

  // Cell 3, 6, and 9 are equal
  case gameData[2] == gameData[5] && gameData[5] == gameData[8]:
    return true, gameData[5]

  // Cell 1, 5, and 9 are equal
  case gameData[0] == gameData[4] && gameData[4] == gameData[8]:
    return true, gameData[4]

  // Cell 3, 5, and 7 are equal
  case gameData[2] == gameData[4] && gameData[4] == gameData[6]:
    return true, gameData[4]

  default:
    return false, "null"
  }
}

func playAgain() {
  var playAgain string
  fmt.Print("\nWould you like to play again? [yes / no] ")
  fmt.Scan(&playAgain)

  if playAgain == "yes" {
    initGame()
  } else {
    os.Exit(0)
  }
}
