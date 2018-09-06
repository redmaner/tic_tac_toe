package AI


// The AI and the opponent can no longer win. Game will be a draw
// AI plays the first available spot on the board
func playRandom(gameData []string, ot string, ait string) {

  bs := getBoardSize(gameData)
  cells := bs * bs + 1

  for i := 1; i < cells; i++ {
    if gameData[i] != ot && gameData[i] != ait {
      gameData[i] = ait
      break
    } else {
      continue
    }
  }
}
