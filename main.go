package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

func input(prompt string) string {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print(prompt)
  scanner.Scan()
  return scanner.Text()
}

func main() {
  if len(os.Args) < 2 {
    var src []string
    
    for {
      text := input("> ")
      op := strings.ToUpper(text)
      
      if strings.Split(op, " ")[0] == "LABEL" {
        for {
          text = input("--> ")
          if strings.ToUpper(text) == "END" {
            src = append(src, text)
            break
          }
          src = append(src, text)
        }
      }
      if op == "END" {
        src = append(src, text)
        break
      }

      src = append(src, text)
    }

    lang.Run(src, 16)
  } else {
    var src []string
    f, err := os.Open(os.Args[1])
    if err != nil { panic(err) }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
      src = append(src, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
      panic(err)
    }
    lang.Run(src, 16)
  }
}
