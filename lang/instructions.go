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

  "ADV": "ADV",
  "DADV": "DADV",

  "PRINTS": "PRINTS",
  "PRINT": "PRINT",

  "END": "END",
}

func NewInstr(op string, arg string, label *Label) *Instr {
  if arg == "" {
    arg = ""
  }
  return &Instr{
    Op: op,
    Arg: arg,
    Label: label,
  }
}

type Instr struct {
  Op string
  Arg string
  Label *Label
}

type Label struct {
  Name string
  program []*Instr
}
