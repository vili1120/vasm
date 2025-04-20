package lang

import (
	//"fmt"
	"os"
)

func NewInterpreter(instrs []*Instr, stack *Stack) *Interpreter {
  i := &Interpreter{
    Instrs: instrs,
    Pc: -1,
    Stack: stack,
  }
  i.Advance()
  return i
}

type Interpreter struct {
  Instrs []*Instr
  CInstr *Instr
  Pc int
  Stack *Stack
}

func (i *Interpreter) Advance() {
  i.Pc++
  if i.Pc < len(i.Instrs) {
    i.CInstr = i.Instrs[i.Pc]
  }
}

func (i *Interpreter) Interpret() {
  //fmt.Println("debug> interpreting")
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

    case instructions["END"]:
      os.Exit(0)
    }
  }
}
