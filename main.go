package main

import (
  "os"
  "strconv"
  "github.com/redmaner/tictacgo/game"
)

func main() {
  args := os.Args
  if len(args) > 1 {
    boardSize, _ := strconv.Atoi(args[1])
    if boardSize > 0 && boardSize <= 100 {
      game.Start(boardSize)
    }
  } else {
    game.Start(3)
  }
}
