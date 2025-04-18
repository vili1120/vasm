package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////

func NewStack(length int) *Stack {
  cells := make([]int, length)
  s := &Stack{
    sp: -1,
    cells: cells,
  }
  s.Advance()
  return s
}

type Stack struct {
  sp int
  cells []int
}

func (s *Stack) Advance() {
  s.sp++
  if s.sp >= len(s.cells) {
    s.sp--
  }
}

func (s *Stack) DeAdvance() {
  s.sp--
  if s.sp < 0 {
    s.sp++
  }
}

func (s *Stack) Remove() {
  s.cells[s.sp] = 0
}

func (s *Stack) Push(val int) {
  s.cells[s.sp] = val
  s.Advance()
}

func (s *Stack) Pull() {
  fmt.Println(s.cells[s.sp])
}

func (s *Stack) Read() {
  var i string
  fmt.Scanln(&i)
  val, err := strconv.Atoi(i)
  if err != nil {
    fmt.Println("Invalid input")
    return
  }
  s.Push(val)
}

func (s *Stack) PrintS() {
  fmt.Println(s.cells)
}

func (s *Stack) Print(idx string) {
  val, err := strconv.Atoi(idx)
  if err != nil {
    fmt.Println("Invalid index")
    return
  }
  if !(val > len(s.cells)) || !(val < 0) {
    fmt.Println(s.cells[val])
  }
}

func (s *Stack) Len() {
  fmt.Println(len(s.cells))
}

func (s *Stack) Add() {
  s.DeAdvance()
  if s.sp-1 < 0 {
    panic("Index out of range!")
    return
  } else {
    s.cells[s.sp-1] = s.cells[s.sp-1]+s.cells[s.sp]
    s.Remove()
    s.DeAdvance()
  }
}

func (s *Stack) Sub() {
  s.DeAdvance()
  if s.sp-1 < 0 {
    panic("Index out of range!")
    return
  } else {
    s.cells[s.sp-1] = s.cells[s.sp-1]-s.cells[s.sp]
    s.Remove()
    s.DeAdvance()
  }
}

func (s *Stack) Idx() {
  fmt.Println(s.sp)
}

//////////////////////////////////////////////////////////////////////////////////////

type Tok struct {
  type_ string
  val string
}

//////////////////////////////////////////////////////////////////////////////////////

func lex(text []string) []Tok {
  toks := []Tok{}

  for _, line := range text {
    cline := strings.Split(line, " ")
    op := cline[0]
    arg := ""
    if len(cline) > 1 {
      arg = cline[1]
    }
    if strings.HasPrefix(op, "#") ||strings.HasPrefix(op, "//") {
      continue
    }
    switch op {
    case "PUSH":
      if arg != "" {
        toks = append(toks, Tok{op, arg})
      } else {
        panic("No argument for PUSH!")
      }
    case "PULL":
      if arg == "" {
        toks = append(toks, Tok{op, arg})
      } else {
        panic("Too many arguments for PULL!")
      }
    case "READ":
      if arg == "" {
        toks = append(toks, Tok{op, arg})
      } else {
        panic("Too many arguments for READ!")
      }

    case "DADV":
      if arg == "" {
        toks = append(toks, Tok{op, arg})
      } else {
        panic("Too many arguments for DADV!")
      }
    case "ADV":
      if arg == "" {
        toks = append(toks, Tok{op, arg})
      } else {
        panic("Too many arguments for ADV!")
      }

    case "REMOVE":
      if arg == "" {
        toks = append(toks, Tok{op, arg})
      } else {
        panic("Too many arguments for REMOVE!")
      }

    case "PRINTS":
      if arg == "" {
        toks = append(toks, Tok{op, arg})
      } else {
        panic("Too many arguments for PRINTS!")
      }
    case "PRINT":
      if arg != "" {
        toks = append(toks, Tok{op, arg})
      } else {
        panic("No argument for PRINT!")
      }

    case "LEN":
      if arg == "" {
        toks = append(toks, Tok{op, arg})
      } else {
        panic("Too many arguments for LEN!")
      }

    case "ADD":
      if arg == "" {
        toks = append(toks, Tok{op, arg})
      } else {
        panic("Too many arguments for ADD!")
      }
    case "SUB":
      if arg == "" {
        toks = append(toks, Tok{op, arg})
      } else {
        panic("Too many arguments for SUB!")
      }

    case "IDX":
      if arg == "" {
        toks = append(toks, Tok{op, arg})
      } else {
        panic("Too many arguments for IDX!")
      }

    case "END":
      if arg == "" {
        toks = append(toks, Tok{op, arg})
      } else {
        panic("Too many arguments for END!")
      }
    case "":
      continue
    default:
      panic("Expected a command or END!")
    }
  }

  return toks
}

//////////////////////////////////////////////////////////////////////////////////////

func interpret(toks []Tok, s *Stack) {
  for _, tok := range toks {
    switch tok.type_ {

    case "PUSH":
      val, err := strconv.Atoi(tok.val)
      if err != nil {
        fmt.Printf("Invalid PUSH value: %v\n", tok.val)
        return
      }
      s.Push(val)
    case "PULL":
      s.Pull()
    case "READ":
      s.Read()

    case "DADV":
      s.DeAdvance()
    case "ADV":
      s.Advance()

    case "REMOVE":
      s.Remove()

    case "PRINTS":
      s.PrintS()
    case "PRINT":
      s.Print(tok.val)
    case "LEN":
      s.Len()

    case "ADD":
      s.Add()
    case "SUB":
      s.Sub()

    case "IDX":
      s.Idx()

    case "END":
      return
    }
  }
}

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

func input_full() []string {
	var str []string
	scanner := bufio.NewReader(os.Stdin)

	for {
    fmt.Print("> ")
		line, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			break
		}

		line = strings.TrimSpace(line)
		if line == "END" {
			break
		}

		str = append(str, line)
	}

	return str
}

func input_cmd() []string {
	var str []string
	scanner := bufio.NewReader(os.Stdin)

  fmt.Print("> ")
	line, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err)
		os.Exit(0)
	}

	line = strings.TrimSpace(line)
	if line == "END" {
		os.Exit(0)
	}

	str = append(str, line)

	return str
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
      src = input_full()
      toks := lex(src)
      s := NewStack(length)
      interpret(toks, s)
    } else if mode == "cmd" {
      s := NewStack(length)
      for {
        src = input_cmd()
        toks := lex(src)
        interpret(toks, s)
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
    toks := lex(src)
    s := NewStack(length)
    interpret(toks, s)
  }
}
