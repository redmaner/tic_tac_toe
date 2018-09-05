package main

import (
  "fmt"
  "strconv"
)

func getPosition(gameData []string, playerToken string) {

  var playerPosition int

  fmt.Print("Please enter desired position [1-9]: ")
  fmt.Scan(&playerPosition)

  if inputValid(playerPosition) {
    if positionValid(gameData, playerPosition) {
      gameData[playerPosition - 1] = playerToken
    } else {
      inputError(gameData, playerToken, "invalidPosition")
    }
  } else {
    inputError(gameData, playerToken, "invalidInput")
  }
}

func inputValid(pp int) bool {
  return pp > 0 && pp <= 9
}

func positionValid(gameData []string, pp int) bool {
  pps := strconv.Itoa(pp)
  return gameData[pp - 1] == pps
}

func inputError(gameData []string, playerToken string, err string) {
  switch err {
  case "invalidInput":
    fmt.Println("Input was invalid. Input must be a number betweeen 1 and 9")
    getPosition(gameData, playerToken)
  case "invalidPosition":
    fmt.Println("Position is already taken")
    getPosition(gameData, playerToken)
  default:
    return
  }
}
