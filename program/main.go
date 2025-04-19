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
  var length int
  for !(length >= 8) {
    fmt.Print("Enter Stack Length(min.: 8)> ")
    fmt.Scanln(&length)
  }

  clear()

  var src []string
  if len(os.Args) < 2 {
    var mode string
    fmt.Print("Set REPL mode(full/cmd)> ")
    fmt.Scanln(&mode)

    clear()

    if mode == "full" {
      fmt.Println("To end the program type END")
      src = lang.InputFull()
      toks := lang.Lex(src)
      labels := lang.CollectLabels(toks)
      s := lang.NewStack(length)
      lang.Interpret(toks, s, labels)
    } else if mode == "cmd" {
      fmt.Println("To end the program type END")
      s := lang.NewStack(length)
      var slabels []string
      for {
        src = lang.InputCmd(&slabels)
        if src == nil {
          continue
        }
        toks := lang.Lex(src)
        labels := lang.CollectLabels(toks)
        lang.Interpret(toks, s, labels)
      }
    }
  } else {
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
    toks := lang.Lex(src)
    labels := lang.CollectLabels(toks)
    s := lang.NewStack(length)
    lang.Interpret(toks, s, labels)
  }
}
