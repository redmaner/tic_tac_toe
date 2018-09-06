package AI

import (
  "strconv"
)

// Function for AI to make a play
func Play(gameData []string, ot string, ait string) {

  if playedObstructive := playObstructive(gameData, ot, ait); playedObstructive {
    return
  } else if playedForWin := playForWin(gameData, ot, ait); playedForWin {
    return
  } else {
    playRandom(gameData, ot, ait)
  }
}

func getBoardSize(gameData []string) int {
  boardSize, _ := strconv.Atoi(gameData[0])
  return boardSize
}

func countTokens(ss []string, ot string, ait string) (int, int) {
  oCounter := 0
  aiCounter := 0
  for i := 0; i < len(ss); i++ {
    if ss[i] == ot {
      oCounter++
    } else if ss[i] == ait {
      aiCounter++
    }
  }
  return oCounter, aiCounter
}
