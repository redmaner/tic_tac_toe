package AI

// First priority of the AI is to extend the game as long as possible
// This means that the other player cannot finish his game by completing a
// winning sequence.
// The playObstructive function will try to obstruct winning sequences of the
// other player.
func playObstructive(gameData []string, ot string, ait string) bool {

  // Determine whether the opponent has already made a move
  if otc, _ := countTokens(gameData, ait, ot); otc == 0 {
    return false
  }

  // boardSize
  bs := getBoardSize(gameData)

  // Obstruction meter
  // The amount of tokens the other player is allowed to have in a sequence
  // before the AI will trigger obstructive play
  var om int // We want this to be an integer at all costs
  om = bs / 2 + 1

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
    if otc >= om && aitc == 0 {
      makePlay(gameData, br, ot, ait)
      return true
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
    if otc >= om && aitc == 0 {
      makePlay(gameData, bc, ot, ait)
      return true
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
    if otc >= om && aitc == 0 {
      makePlay(gameData, bdl, ot, ait)
      return true
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
    if otc >= om && aitc == 0 {
      makePlay(gameData, bdr, ot, ait)
      return true
    }
  }

  return false
}
