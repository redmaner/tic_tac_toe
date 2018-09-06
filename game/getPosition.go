package game

import (
  "fmt"
  "strconv"
)

// Function that requests the desired postion on the board of the player
func getPosition(gameData []string, playerToken string) {

  var playerPosition int

  fmt.Print("Please enter desired position [1-9]: ")
  fmt.Scan(&playerPosition)

  if inputValid(playerPosition) {
    if positionValid(gameData, playerPosition) {
      gameData[playerPosition] = playerToken
    } else {
      inputError(gameData, playerToken, "invalidPosition")
    }
  } else {
    inputError(gameData, playerToken, "invalidInput")
  }
}

// Function that validates the user input
// The input of the player must be an integer between 1 and 9
func inputValid(pp int) bool {
  return pp > 0 && pp <= 9
}

// Function that validates whether the desired position is free on the board
// If the postion is already taken, the player has to put in another postion
func positionValid(gameData []string, pp int) bool {
  pps := strconv.Itoa(pp)
  return gameData[pp] == pps
}

// Universal function to call an inputError caused by either an invalid input or
// a desired postion already being taken.
// This function requests the player to put in a desired position once more
// This will loop until the postion that is given by the player is valid
// according to inputValid() and positionValid()
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
