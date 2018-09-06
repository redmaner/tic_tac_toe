package board

import (
  "fmt"
  "os"
  "os/exec"
  "runtime"
)

// clearScreen is a cross platform function to clear the screen
func clearScreen() {

  switch runtime.GOOS {
  case "linux":
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
  case "windows":
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
  default:
    for s := 1; s <= 5; s++ {
      fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
    }
  }
}
