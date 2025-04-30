package lang

import "fmt"

const LETTERSDIGITS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var KEYWORDS = map[string]string{
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

  "IF:EQ": "IF:EQ",
  "IF:NE": "IF:NE",
  "IF:LT": "IF:LT",
  "IF:GT": "IF:GT",
  "IF:LTE": "IF:LTE",
  "IF:GTE": "IF:GTE",

  "PRINTS": "PRINTS",
  "PRINT": "PRINT",

  "ADV": "ADV",
  "DADV": "DADV",
  "END": "END",
}

const (
  NEWLINE = "NEWLINE"
)

func NewToken(type_ string, value string) Token {
  return Token{
    Type: type_,
    Value: value,
  }
}

type Token struct {
  Type string
  Value string
}

func (t Token) String() string {
  if t.Value != "" {
    return fmt.Sprintf("%v: %v", t.Type, t.Value)
  }
  return fmt.Sprintf("%v", t.Type)
}
