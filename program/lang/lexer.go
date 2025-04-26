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
  Args []string
  Idx int
}

func (l *Lexer) Advance() {
  for {
    l.Idx++
    if l.Idx >= len(l.Lines) {
      l.Op = ""
      l.Args =  []string{}
      return
    }

    line := strings.TrimSpace(l.Lines[l.Idx])
    if line == "" {
      continue
    }

    l.Ln = strings.Split(line, " ")
    l.Op = strings.ToUpper(l.Ln[0])
    l.Args = []string{}
    if len(l.Ln) > 1 {
      for _, arg := range l.Ln[1:] {
        l.Args = append(l.Args, strings.ToUpper(arg))
      }
    }
    break
  }
}


func (l *Lexer) IsArg() bool {
  if len(l.Args) >= 1 {
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

    if strings.HasPrefix(l.Op, "RL:") {
      l.MakeRL(&INSTRS)
      continue
    }

    switch l.Op {
    case instructions["PUSH"]:
      if l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, l.Args, nil, nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too few arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }
    case instructions["PULL"], instructions["READ"], instructions["POP"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, []string{}, nil, nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too many arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }

    case instructions["ADD"], instructions["SUB"], instructions["MUL"], instructions["DIV"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, []string{}, nil, nil, nil))
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
        INSTRS = append(INSTRS, NewInstr(l.Op, l.Args, nil, nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too few arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }

    case instructions["ADV"], instructions["DADV"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, []string{}, nil, nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too many arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }

    case instructions["PRINTS"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, []string{}, nil, nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too many arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }
    case instructions["PRINT"]:
      if l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, l.Args, nil, nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too few arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }

    case instructions["FI"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, []string{}, nil, nil, nil))
        l.Advance()
      } else {
        LexerError(fmt.Sprintf("Too many arguments for %v", l.Op), l.Ln, l.Idx)
        l.Advance()
      }

    case instructions["ROF"]:
      if !l.IsArg() {
        INSTRS = append(INSTRS, NewInstr(l.Op, []string{}, nil, nil, nil))
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

  INSTRS = append(INSTRS, NewInstr(l.Op, []string{}, nil, nil, nil))

  return INSTRS
}

func (l *Lexer) MakeLabel(instrs *[]*Instr) {
  program := []string{}
  op := l.Op
  name := l.Args[0]
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
  *instrs = append(*instrs, NewInstr(op, []string{name}, &Label{name, linstr}, nil, nil))
}

func (l *Lexer) MakeIf(instrs *[]*Instr) {
  program := []string{}
  op := l.Op
  name := l.Args[0]
  type_ := ""
  switch l.Op {
  case instructions["IF.EQ"]:
    type_ = "EQ"
  case instructions["IF.NE"]:
    type_ = "NE"
  case instructions["IF.LT"]:
    type_ = "LT"
  case instructions["IF.GT"]:
    type_ = "GT"
  case instructions["IF.LTE"]:
    type_ = "LTE"
  case instructions["IF.GTE"]:
    type_ = "GTE"
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

  *instrs = append(*instrs, NewInstr(op, []string{name}, nil, &IfInstr{type_: type_, val: name, body: linstr}, nil))
}

func (l *Lexer) MakeRL(instrs *[]*Instr) {
  switch l.Op {
  case instructions["RL:INIT"]:
    if len(l.Args) == 2 {
      *instrs = append(*instrs, NewInstr(l.Op, l.Args, nil, nil, nil))
      l.Advance()
    } else if len(l.Args) < 2 {
      LexerError("Too few arguments for RL:INIT", l.Ln, l.Idx)
      l.Advance()
    } else {
      LexerError("Too many arguments for RL:INIT", l.Ln, l.Idx)
      l.Advance()
    }
  case instructions["RL:FORCLOSE"]:
    if !l.IsArg() {
      program := []string{}
      op := l.Op
      args := l.Args
      l.Advance()

      for l.Idx < len(l.Lines) {
        line := strings.TrimSpace(l.Lines[l.Idx])
        if line == "" {
          l.Advance()
          continue
        }

        if strings.ToUpper(line) == instructions["ROF"] {
          program = append(program, line)
          l.Advance()
          break
        }
        if line != "" {
          program = append(program, line)
        }

        l.Advance()
      }

      lexer := NewLexer(program)
      loop := lexer.MakeInstr()
      *instrs = append(*instrs, NewInstr(op, args, nil, nil, loop))
    } else {
      LexerError("Too many arguments for RL:FORCLOSE", l.Ln, l.Idx)
      l.Advance()
      return
    }
  case instructions["RL:FPS"]:
    if l.IsArg() {
      *instrs = append(*instrs, NewInstr(l.Op, l.Args, nil, nil, nil))
    } else {
      LexerError("Too many or too few arguments for RL:FPS", l.Ln, l.Idx)
    }
    l.Advance()
  case instructions["RL:SAMPLEDRAW"]:
    if !l.IsArg() {
      *instrs = append(*instrs, NewInstr(l.Op, l.Args, nil, nil, nil))
    } else {
      LexerError("Too many arguments for RL:SAMPLEDRAW", l.Ln, l.Idx)
    }
    l.Advance()
  case instructions["RL:BEGINDRAWING"]:
    if !l.IsArg() {
      *instrs = append(*instrs, NewInstr(l.Op, l.Args, nil, nil, nil))
    } else {
      LexerError("Too many arguments for RL:BEGINDRAWING", l.Ln, l.Idx)
    }
    l.Advance()
  case instructions["RL:ENDDRAWING"]:
    if !l.IsArg() {
      *instrs = append(*instrs, NewInstr(l.Op, l.Args, nil, nil, nil))
    } else {
      LexerError("Too many arguments for RL:ENDDRAWING", l.Ln, l.Idx)
    }
    l.Advance()
  case instructions["RL:CLEAR"]:
    if l.IsArg() {
      *instrs = append(*instrs, NewInstr(l.Op, l.Args, nil, nil, nil))
    } else {
      LexerError("Too many or too few arguments for RL:CLEAR", l.Ln, l.Idx)
    }
    l.Advance()
  case instructions["RL:CLOSE"]:
    if !l.IsArg() {
      *instrs = append(*instrs, NewInstr(l.Op, l.Args, nil, nil, nil))
    } else {
      LexerError("Too many arguments for RL:CLOSE", l.Ln, l.Idx)
    }
    l.Advance()
  case instructions["RL:EXEC"]:
    if !l.IsArg() {
      *instrs = append(*instrs, NewInstr(l.Op, l.Args, nil, nil, nil))
    } else {
      LexerError("Too many arguments for RL:EXEC", l.Ln, l.Idx)
    }
    l.Advance()
  }
}
