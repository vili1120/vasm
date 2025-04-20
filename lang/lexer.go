package lang

import (
	//"fmt"
	"fmt"
	"strings"
)

func NewLexer(text []string) *Lexer {
  l := &Lexer{
    Lines: text,
    Idx: -1,
  }
  l.Advance()
  return l
}

type Lexer struct {
  Lines []string
  Ln []string
  Op string
  Arg string
  Idx int
}

func (l *Lexer) Advance() {
  l.Idx++
  if l.Idx < len(l.Lines) {
    l.Ln = strings.Split(l.Lines[l.Idx], " ")
    l.Op = strings.ToUpper(l.Ln[0])
    if len(l.Ln) > 1 {
      l.Arg = strings.ToUpper(l.Ln[1])
    } else {
      l.Arg = ""
    }
  }
}

func (l *Lexer) IsArg() bool {
  if l.Arg != "" {
    return true
  }
  return false
}

func (l *Lexer) MakeInstr() []*Instr {
  INSTRS := []*Instr{}

  //fmt.Println("debug> lexing")

  for l.Op != instructions["END"] {
    switch l.Op {
    case instructions["PUSH"]:
      if l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, l.Arg))
        l.Advance()
      }
    case instructions["PULL"], instructions["READ"], instructions["POP"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, ""))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too many arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }

    case instructions["ADV"], instructions["DADV"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, ""))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too many arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }

    case instructions["PRINTS"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, ""))
        l.Advance()
      }
    case instructions["PRINT"]:
      if l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, l.Arg))
        l.Advance()
      }

    case "":
      continue
    }
  }

  INSTRS = append(INSTRS, NewInstr(l.Op, ""))

  return INSTRS
}
