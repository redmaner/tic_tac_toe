package main

import (
  "fmt"
  "os"
  "strconv"
)

// Function that involves the logic of a round which is played a maximum of nine
// times
func doTurn(gameData []string, round int, playerToken string, playerName string) {
  showBoard(gameData)
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
      showWinner(gameData, winnerToken)
      playAgain()
    } else if !isWinner && round == 9 {
      showWinner(gameData, "null")
      playAgain()
    }
  }
}

// Function that checks if there is a winner
func checkWinner(gameData []string) (bool, string) {

  // bs == boardSize
  bs, _ := strconv.Atoi(gameData[0])

  // bi == boardIndex
  bri := 1

  // Check horizontal win conditions
  for i := 1; i <= bs; i++ {

    if i != 1 {
      bri += bs
    }

    // br == boardRow == slice of strings with boardsize items
    br := []string{}
    for c := 0; c < bs; c++ {
      br = append(br, gameData[bri+c])
    }

    isWinner, winnerToken := checkIfWinCondition(br, bs)
    if isWinner {
      return isWinner, winnerToken
    }
  }

  // Check vertical win conditions
  for i := 1; i <= bs; i++ {

    // bc == boardColumn == slice of strings with boardsize items
    bc := []string{}
    bci := i
    for c := 0; c < bs; c++ {
      if c != 0 {
        bci += bs
      }
      bc = append(bc, gameData[bci])
    }

    isWinner, winnerToken := checkIfWinCondition(bc, bs)
    if isWinner {
      return isWinner, winnerToken
    }
  }

  // Check diagonal from left to right
  bdl := []string{}
  bdli := 1
  for c := 0; c < bs; c++ {
    if c != 0 {
      bdli += bs + 1
    }
    bdl = append(bdl, gameData[bdli])
  }

  isWinner, winnerToken := checkIfWinCondition(bdl, bs)
  if isWinner {
    return isWinner, winnerToken
  }

  // Check diagonal from right to left
  bdr := []string{}
  bdri := bs
  for c := 0; c < bs; c++ {
    if c != 0 {
      bdri += bs - 1
    }
    bdr = append(bdr, gameData[bdri])
  }

  isWinner, winnerToken = checkIfWinCondition(bdr, bs)
  if isWinner {
    return isWinner, winnerToken
  }

  return false, "null"
}

func checkIfWinCondition(cs []string, bs int) (bool, string) {
  var cv string
  for s := 0; s < bs; s++ {
    if cv == "" {
      cv = cs[s]
    } else if cv == cs[s] && s == bs-1 {
      return true, cs[s]
    } else if cv == cs[s] {
      continue
    } else {
      break
    }
  }
  return false, "null"
}

// Function that asks the player if he wants to play again
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
