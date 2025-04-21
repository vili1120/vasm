package lang

import (
	//"fmt"
	"fmt"
	//"os"
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
  for {
    l.Idx++
    if l.Idx >= len(l.Lines) {
      l.Op = ""
      l.Arg = ""
      return
    }

    line := strings.TrimSpace(l.Lines[l.Idx])
    if line == "" {
      continue
    }

    l.Ln = strings.Split(line, " ")
    l.Op = strings.ToUpper(l.Ln[0])
    l.Arg = ""
    if len(l.Ln) > 1 {
      l.Arg = strings.ToUpper(l.Ln[1])
    }
    break
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
  
  if Debug == true {
    fmt.Println("debug> lexing")
  }

  for l.Op != instructions["END"] {
    if l.Op == "" {
      l.Advance()
      break
    }

    if strings.HasPrefix(l.Op, "#") {
      l.Advance()
      continue
    }

    if strings.HasPrefix(l.Op, "IF.") {
      l.MakeIf(&INSTRS)
      continue
    }

    switch l.Op {
    case instructions["PUSH"]:
      if l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, l.Arg, nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too few arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }
    case instructions["PULL"], instructions["READ"], instructions["POP"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, "", nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too many arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }

    case instructions["ADD"], instructions["SUB"], instructions["MUL"], instructions["DIV"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, "", nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too many arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }


    case instructions["LABEL"]:
      if l.IsArg() {
        l.MakeLabel(&INSTRS)
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too few arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }
    case instructions["JUMP"]:
      if l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, l.Arg, nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too few arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }

    case instructions["ADV"], instructions["DADV"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, "", nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too many arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }

    case instructions["PRINTS"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, "", nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too many arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }
    case instructions["PRINT"]:
      if l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, l.Arg, nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too few arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }

    case instructions["FI"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, l.Arg, nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too many arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }

    case "":
      continue
    default:
      LexerError("Instruction not defined:", l.Ln, l.Idx)
      l.Advance()
    }
  }

  INSTRS = append(INSTRS, NewInstr(l.Op, l.Arg, nil, nil))

  return INSTRS
}

func (l *Lexer) MakeLabel(instrs *[]*Instr) {
  program := []string{}
  op := l.Op
  name := l.Arg
  l.Advance()

  for l.Idx < len(l.Lines) {
    line := strings.TrimSpace(l.Lines[l.Idx])
    if line == "" {
      l.Advance()
      continue
    }

    if strings.ToUpper(line) == "END" {
        program = append(program, line)
        break
    }
    if line != "" {
      program = append(program, line)
    }

    l.Advance()
  }

  lexer := NewLexer(program)
  linstr := lexer.MakeInstr()
  *instrs = append(*instrs, NewInstr(op, name, &Label{name, linstr}, nil))
}

func (l *Lexer) MakeIf(instrs *[]*Instr) {
  program := []string{}
  op := l.Op
  name := l.Arg
  type_ := ""
  switch l.Op {
  case instructions["IF.EQ"]:
    type_ = "EQ"
  case instructions["IF.NE"]:
    type_ = "NE"
  }
  l.Advance()

  for l.Idx < len(l.Lines) {
    line := strings.TrimSpace(l.Lines[l.Idx])
    if line == "" {
      l.Advance()
      continue
    }

    if strings.ToUpper(line) == "FI" {
        program = append(program, line)
        break
    }
    if line != "" {
      program = append(program, line)
    }

    l.Advance()
  }

  lexer := NewLexer(program)
  linstr := lexer.MakeInstr()

  *instrs = append(*instrs, NewInstr(op, name, nil, &IfInstr{type_: type_, val: name, body: linstr}))
}
