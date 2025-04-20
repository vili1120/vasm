package lang

import (
	"fmt"
	"os"
)

func NewInterpreter(instrs []*Instr, stack *Stack) *Interpreter {
  i := &Interpreter{
    Instrs: instrs,
    Pc: -1,
    Stack: stack,
    Labels: map[string]*Label{},
  }
  i.CollectLabels()
  i.Advance()
  return i
}

type Interpreter struct {
  Instrs []*Instr
  CInstr *Instr
  Pc int
  Stack *Stack
  Labels map[string]*Label
}

func (i *Interpreter) CollectLabels() {
  for _, instr := range i.Instrs {
    if instr.Op == "LABEL" && instr.Label != nil {
      i.Labels[instr.Label.Name] = instr.Label
    }
  }
}

func (i *Interpreter) Advance() {
  i.Pc++
  if i.Pc < len(i.Instrs) {
    i.CInstr = i.Instrs[i.Pc]
  }
}

func (i *Interpreter) Interpret() {
  if Debug == true {
    fmt.Println("debug> interpreting")
  }
  for {
    switch i.CInstr.Op {
    case instructions["PUSH"]:
      i.Stack.Push(i.CInstr.Arg)
      i.Advance()
    case instructions["PULL"]:
      i.Stack.Pull()
      i.Advance()
    case instructions["READ"]:
      i.Stack.Read()
      i.Advance()
    case instructions["POP"]:
      i.Stack.Pop()
      i.Advance()

    case instructions["ADD"]:
      i.Stack.Add()
      i.Advance()
    case instructions["SUB"]:
      i.Stack.Sub()
      i.Advance()
    case instructions["MUL"]:
      i.Stack.Mul()
      i.Advance()
    case instructions["DIV"]:
      i.Stack.Div()
      i.Advance()

    case instructions["ADV"]:
      i.Stack.Advance()
      i.Advance()
    case instructions["DADV"]:
      i.Stack.DeAdvance()
      i.Advance()

    case instructions["PRINTS"]:
      i.Stack.Prints()
      i.Advance()
    case instructions["PRINT"]:
      i.Stack.Print(i.CInstr.Arg)
      i.Advance()

    case instructions["LABEL"]:
      i.Advance()
      continue
    case instructions["JUMP"]:
      label, ok := i.Labels[i.CInstr.Arg]
      if !ok {
        InterpreterError("Label doesn't exist", i.CInstr, i.Pc)
        i.Advance()
        break
      }
      i2 := &Interpreter{
        Instrs: label.program,
        Pc: -1,
        Stack: i.Stack,
        Labels: i.Labels,
      }
      i2.Advance()
      i2.Interpret()
      i2.Advance()

    case instructions["END"]:
      os.Exit(0)
    }
  }
}
