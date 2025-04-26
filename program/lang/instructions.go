package lang

import "fmt"

var (
  Debug = false
)

var instructions = map[string]string{
  "PUSH": "PUSH",
  "PULL": "PULL",

  "READ": "READ",
  "POP": "POP",

  "ADD": "ADD",
  "SUB": "SUB",
  "MUL": "MUL",
  "DIV": "DIV",

  "LABEL": "LABEL",
  "JUMP": "JUMP",

  "IF.EQ": "IF.EQ",
  "IF.NE": "IF.NE",
  "IF.LT": "IF.LT",
  "IF.GT": "IF.GT",
  "IF.LTE": "IF.LTE",
  "IF.GTE": "IF.GTE",
  "FI": "FI",

  "RL:INIT": "RL:INIT",
  "RL:FPS": "RL:FPS",
  "RL:FORCLOSE": "RL:FORCLOSE",
  "RL:CLOSE": "RL:CLOSE",
  "RL:SAMPLEDRAW": "RL:SAMPLEDRAW",
  "RL:BEGINDRAWING": "RL:BEGINDRAWING",
  "RL:ENDDRAWING": "RL:ENDDRAWING",
  "RL:CLEAR": "RL:CLEAR",

  "RL:EXEC": "RL:EXEC",

  "ROF": "ROF",

  "ADV": "ADV",
  "DADV": "DADV",

  "PRINTS": "PRINTS",
  "PRINT": "PRINT",

  "END": "END",
}

func NewInstr(op string, arg []string, label *Label, ifinstr *IfInstr, loop []*Instr) *Instr {
  return &Instr{
    Op: op,
    Arg: arg,
    Label: label,
    IfInstr: ifinstr,
    Loop: loop,
  }
}

type Instr struct {
  Op string
  Arg []string
  Label *Label
  IfInstr *IfInstr
  Loop []*Instr
}

func (i *Instr) String() string {
  return fmt.Sprintf("OP: %v, ARG: %v\nLABEL: %v, IFINSTR: %v\nLOOP: %v\n",
    i.Op, i.Arg, i.Label, i.IfInstr, i.Loop,
  )
}

type Label struct {
  Name string
  program []*Instr
}

type IfInstr struct {
  type_ string
  val string
  body []*Instr
}
