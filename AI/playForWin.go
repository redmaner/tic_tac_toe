package AI

import (
  "math/rand"
  "strconv"
  "time"
)

func playForWin(gameData []string, ot string, ait string) bool {

  // boardSize
  bs := getBoardSize(gameData)

  // Max possible sequences
  ms := bs + bs + 2

  // Empty multi dimensional slice, to store available sequences as a slice of strings
  as := make([][]string, 0, ms)

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
    if otc == 0 && aitc > seqTokens {
      seq = br
      seqTokens = aitc
    } else if otc == 0 && aitc == 0 {
      as = append(as, br)
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
    if otc == 0 && aitc > seqTokens {
      seq = bc
      seqTokens = aitc
    } else if otc == 0 && aitc == 0 {
      as = append(as, bc)
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
    if otc == 0 && aitc > seqTokens {
      seq = bdl
      seqTokens = aitc
    } else if otc == 0 && aitc == 0 {
      as = append(as, bdl)
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
    if otc == 0 && aitc > seqTokens {
      seq = bdr
      seqTokens = aitc
    } else if otc == 0 && aitc == 0 {
      as = append(as, bdr)
    }
  }

  // If a sequence has already AI tokens, play the sequence with the most AI seq
  // This is the play for a winning sequence
  if seq != nil {
    makePlayForWin(gameData, seq, ot, ait)
    return true
  }

  // If there is no sequence yet with an AI token, and there a sequences where
  // the opponent hasn't played yet, play an availble sequence.
  // If there are more availble sequences, play a random sequence.
  // This will make the play of the AI more unpredictable
  if len(as) != 0 {
    seed := rand.NewSource(time.Now().UnixNano())
    rs := rand.New(seed)

    rsv := rs.Intn(len(as))
    makePlayForWin(gameData, as[rsv], ot, ait)
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
