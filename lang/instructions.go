package lang

var instructions = map[string]string{
  "PUSH": "PUSH",
  "PULL": "PULL",

  "READ": "READ",
  "POP": "POP",

  "LABEL": "LABEL",
  "JUMP": "JUMP",

  "ADV": "ADV",
  "DADV": "DADV",

  "PRINTS": "PRINTS",
  "PRINT": "PRINT",

  "END": "END",
}

func NewInstr(op string, arg string) *Instr {
  if arg == "" {
    arg = ""
  }
  return &Instr{
    Op: op,
    Arg: arg,
  }
}

type Instr struct {
  Op string
  Arg string
}
