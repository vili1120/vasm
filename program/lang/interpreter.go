package lang

import (
	"fmt"
	"os"
	"os/exec"
	//"strconv"
	"strings"
	//rl "github.com/gen2brain/raylib-go/raylib"
)

func NewInterpreter(instrs []*Instr, stack *Stack) *Interpreter {
  rlcode := `
package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
  `
  i := &Interpreter{
    Instrs: instrs,
    Pc: -1,
    Stack: stack,
    Labels: map[string]*Label{},
    RLcode: &rlcode,
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
  RLcode *string
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
  for i.Pc < len(i.Instrs) {
    if Debug {
      //fmt.Printf("Pc: %d, Instruction: %v\n", i.Pc, i.CInstr)
    }

    if strings.HasPrefix(i.CInstr.Op, "RL:") {
      fmt.Println("Processing instruction:", i.CInstr.Op)
      i.RL()
      continue
    }

    if strings.HasPrefix(i.CInstr.Op, "IF") {
      if i.Stack.If(i.CInstr.IfInstr.val, i.CInstr.IfInstr.type_) {
        i2 := &Interpreter{
          Instrs: i.CInstr.IfInstr.body,
          Pc: -1,
          Stack: i.Stack,
          Labels: i.Labels,
        }
        i2.Advance()
        i2.Interpret()
        i2.Advance()
      }
      for {
        i.Advance()
        if i.Pc >= len(i.Instrs) || i.CInstr.Op == instructions["FI"] {
          break
        }
      }

      i.Advance()
    }

    switch i.CInstr.Op {
    case instructions["PUSH"]:
      i.Stack.Push(i.CInstr.Arg[0])
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
      i.Stack.Print(i.CInstr.Arg[0])
      i.Advance()

    case instructions["LABEL"]:
      i.Advance()
      continue
    case instructions["JUMP"]:
      label, ok := i.Labels[i.CInstr.Arg[0]]
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

    case instructions["FI"]:
      return

    case instructions["ROF"]:
      return

    case instructions["END"]:
      os.Exit(0)
    }
  }
}

func (i *Interpreter) RL() {
  switch i.CInstr.Op {
  case instructions["RL:INIT"]:
    fmt.Println(i.Stack.Get(i.CInstr.Arg[0]))
    fmt.Println(i.Stack.Get(i.CInstr.Arg[1]))
    width := i.Stack.Get(i.CInstr.Arg[0])
    height := i.Stack.Get(i.CInstr.Arg[1])
    i.Stack.Prints()
    *i.RLcode += fmt.Sprintf(`
rl.InitWindow(int32(%v), int32(%v), "")
defer rl.CloseWindow()
`, width, height)
    i.Advance()
    return
  case instructions["RL:FPS"]:
    fps := i.Stack.Get(i.CInstr.Arg[0])
    *i.RLcode += fmt.Sprintf("\nrl.SetTargetFPS(int32(%v))", int32(fps))
    i.Advance()
    return
  case instructions["RL:SAMPLEDRAW"]:
		*i.RLcode += `
    rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
    `
    i.Advance()
    return
  case instructions["RL:FORCLOSE"]:
    *i.RLcode += "\nfor !rl.WindowShouldClose() {"
    i2 := &Interpreter{
      Instrs: i.CInstr.Loop,
      Pc: -1,
      Stack: i.Stack,
      Labels: i.Labels,
      RLcode: i.RLcode,
    }
    i2.Advance()
    i2.Interpret()
    i2.Advance()
    *i.RLcode += "\n}"

    i.Advance()
    return
  case instructions["RL:CLOSE"]:
    *i.RLcode += "\nrl.CloseWindow()"
    i.Advance()
    return
  case instructions["RL:BEGINDRAWING"]:
    *i.RLcode += "\nrl.BeginDrawing()"
    i.Advance()
    return
  case instructions["RL:ENDDRAWING"]:
    *i.RLcode += "\nrl.EndDrawing()"
    i.Advance()
    return
  case instructions["RL:CLEAR"]:
    switch strings.ToUpper(i.CInstr.Arg[0]) {
    case "WHITE":
      *i.RLcode += "\nrl.ClearBackground(rl.White)"
      i.Advance()
      return
    }
  case instructions["RL:EXEC"]:
    *i.RLcode += "\n}"
    fmt.Println("Generating Raylib Go code...")

    err := os.WriteFile("out.go", []byte(*i.RLcode), 0644)
    if err != nil {
      fmt.Println("Error writing Go file:", err)
      os.Exit(1)
    }

    fmt.Println("Running generated Go file...")
    cmd := exec.Command("go", "run", "out.go")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin
    err = cmd.Run()
    if err != nil {
      fmt.Println("Error executing generated Go code:", err)
      os.Exit(1)
    }
    os.Remove("out.go")

    // ✅ Keep advancing
    i.Advance()
    return

  }
}
