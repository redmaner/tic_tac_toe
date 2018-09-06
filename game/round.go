package game

import (
  "fmt"
  "os"
  "strconv"
  "github.com/redmaner/tictacgo/board"
)

// Function that involves the logic of a round which is played a maximum of nine
// times
func round(gameData []string, round int, playerToken string, playerName string) {
  board.Draw(gameData)
  fmt.Println("Round:", round)
  fmt.Println("Player:", playerName)
  fmt.Println("Token:", playerToken)
  fmt.Println("------------------")
  fmt.Println("")

  getPosition(gameData, playerToken)

  // This part determines whether there is a winner or not
  // This part is only ran after round 5, because it takes at least 5 rounds
  // a player can win.
  if round >= 5 {
    isWinner, winnerToken := checkWinner(gameData)
    if isWinner {
      board.End(gameData, winnerToken)
      playAgain()
    } else if !isWinner && round == 9 {
      board.End(gameData, "null")
      playAgain()
    }
  }
}

// Function that asks the player if he wants to play again
func playAgain() {
  var playAgain string
  fmt.Print("\nWould you like to play again? [yes / no] ")
  fmt.Scan(&playAgain)

  if playAgain == "yes" {
    Start()
  } else {
    os.Exit(0)
  }
}

func getMaxRounds(gameData []string) int {
  boardSize, _ := strconv.Atoi(gameData[0])
  return boardSize * boardSize
}
