package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"vasm/lang"
)

//////////////////////////////////////////////////////////////////////////////////////

func clear() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "darwin", "linux":
		cmd = exec.Command("clear")
	default:
		fmt.Println("Unsupported operating system")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error clearing the screen:", err)
	}
}

func main() {
  var src []string
  
  for {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("> ")
    scanner.Scan()
    src = append(src, scanner.Text())

    if scanner.Text() == "end" {
      break
    }
  }

  lang.Run(src, 16)
}
