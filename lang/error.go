package lang

import (
	"fmt"
)

func Error(type_, desc string, line []string, ln int) {
  res := fmt.Sprintf("(Ln: %v)\n", ln+1)
  res += fmt.Sprintf("  ERROR: %v: %v\n", type_, desc)
  res += fmt.Sprintf("    %v\n", line)
  fmt.Print(res)
  return
}

func LexerError(desc string, line []string, ln int) {
  Error("Lexer", desc, line, ln)
}
