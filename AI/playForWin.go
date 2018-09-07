package AI

import (
  "strconv"
)

func playForWin(gameData []string, ot string, ait string) bool {

  // boardSize
  bs := getBoardSize(gameData)

  var seq []string
  var seqTokens int

  // bi == boardIndex
  bri := 1

  // Check horizontal win conditions
  for i := 1; i <= bs; i++ {

    if i != 1 {
      bri += bs
    }

    // br == boardRow == slice of strings with boardsize items
    br := make([]string, 0, bs)
    for c := 0; c < bs; c++ {
      br = append(br, gameData[bri+c])
    }

    // Opponent token count
    otc, aitc := countTokens(br, ot, ait)
    if otc == 0 && aitc >= seqTokens {
      seq = br
      seqTokens = aitc
    }
  }

  // Check vertical win conditions
  for i := 1; i <= bs; i++ {

    // bc == boardColumn == slice of strings with boardsize items
    bc := make([]string, 0, bs)
    bci := i
    for c := 0; c < bs; c++ {
      if c != 0 {
        bci += bs
      }
      bc = append(bc, gameData[bci])
    }

    // Opponent token count
    otc, aitc := countTokens(bc, ot, ait)
    if otc == 0 && aitc >= seqTokens {
      seq = bc
      seqTokens = aitc
    }
  }

  // Check diagonal from left to right
  {
    bdl := make([]string, 0, bs)
    bdli := 1
    for c := 0; c < bs; c++ {
      if c != 0 {
        bdli += bs + 1
      }
      bdl = append(bdl, gameData[bdli])
    }

    // Opponent token count
    otc, aitc := countTokens(bdl, ot, ait)
    if otc == 0 && aitc >= seqTokens {
      seq = bdl
      seqTokens = aitc
    }
  }

  // Check diagonal from right to left
  {
    bdr := make([]string, 0, bs)
    bdri := bs
    for c := 0; c < bs; c++ {
      if c != 0 {
        bdri += bs - 1
      }
      bdr = append(bdr, gameData[bdri])
    }

    // Opponent token count
    otc, aitc := countTokens(bdr, ot, ait)
    if otc == 0 && aitc >= seqTokens {
      seq = bdr
      seqTokens = aitc
    }
  }

  if seq != nil {
    makePlayForWin(gameData, seq, ot, ait)
    return true
  }

  return false
}

func makePlayForWin(gameData []string, ss []string, ot string, ait string) {
  for _, value := range ss {
    if value == ot || value == ait {
      continue
    } else {
      index, _ := strconv.Atoi(value)
      gameData[index] = ait
      break
    }
  }
}
