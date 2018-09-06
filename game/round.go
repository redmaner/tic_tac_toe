package game

import (
  "fmt"
  "os"
  "strconv"
  "time"
  "github.com/redmaner/tictacgo/AI"
  "github.com/redmaner/tictacgo/board"
)

// Function that involves the logic of a round which is played a maximum of nine
// times
func round(gameData []string, round int, playerToken string, opponentToken string, playerName string) {
  maxRounds, boardSize := getData(gameData)

  board.Draw(gameData)
  fmt.Print("Board: ", boardSize, "x", boardSize, "\n")
  fmt.Println("Round:", round)
  fmt.Println("Player:", playerName)
  fmt.Println("Token:", playerToken)
  fmt.Println("------------------")
  fmt.Println("")

  if playerName == "AI" {
    AI.Play(gameData, opponentToken, playerToken)
    time.Sleep(500 * time.Millisecond)
  } else {
    getPosition(gameData, playerToken)
  }

  // This part determines whether there is a winner or not
  // This part is only ran after round 5, because it takes at least 5 rounds
  // a player can win.
  if round >= boardSize * 2 - 1 {
    isWinner, winnerToken := checkWinner(gameData)
    if isWinner {
      board.End(gameData, winnerToken)
      playAgain(boardSize)
    } else if !isWinner && round == maxRounds {
      board.End(gameData, "null")
      playAgain(boardSize)
    }
  }
}

// Function that asks the player if he wants to play again
func playAgain(boardSize int) {
  var playAgain string
  fmt.Print("\nWould you like to play again? [yes / no] ")
  fmt.Scan(&playAgain)

  if playAgain == "yes" {
    Start(boardSize)
  } else {
    os.Exit(0)
  }
}

func getData(gameData []string) (int, int) {
  boardSize, _ := strconv.Atoi(gameData[0])
  return boardSize * boardSize, boardSize
}
