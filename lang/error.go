package lang

import (
	"fmt"
)

func Error(type_, desc string, line []string, ln int) {
  res := fmt.Sprintf("(Ln: %v)\n", ln+1)
  res += fmt.Sprintf("ERROR: %v: %v\n", type_, desc)
  res += fmt.Sprintf("  %v\n", line)
  fmt.Print(res)
  return
}

func LexerError(desc string, line []string, ln int) {
  Error("Lexer", desc, line, ln)
}

func InterpreterError(desc string, line *Instr, ln int) {
  Error("Interpreter", desc, []string{line.Op, line.Arg}, ln)
}

func StackError(desc string) {
  res := fmt.Sprintf("STACK ERROR: %v", desc)
  fmt.Print(res)
  return
}
