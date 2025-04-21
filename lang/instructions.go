package lang

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
  "FI": "FI",

  "ADV": "ADV",
  "DADV": "DADV",

  "PRINTS": "PRINTS",
  "PRINT": "PRINT",

  "END": "END",
}

func NewInstr(op string, arg string, label *Label, ifinstr *IfInstr) *Instr {
  if arg == "" {
    arg = ""
  }
  return &Instr{
    Op: op,
    Arg: arg,
    Label: label,
    IfInstr: ifinstr,
  }
}

type Instr struct {
  Op string
  Arg string
  Label *Label
  IfInstr *IfInstr
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
