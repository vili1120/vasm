package lang

import (
	"fmt"
	"os"
	"strconv"
)

func CollectLabels(toks []Tok) map[string]*Label {
  labels := make(map[string]*Label)
  for _, tok := range toks {
    if tok.type_ == "LABEL" && tok.label != nil {
      labels[tok.val] = tok.label
    }
  }
  return labels
}

func Interpret(toks []Tok, s *Stack, labels map[string]*Label) {
  for _, tok := range toks {
    switch tok.type_ {

    case "PUSH":
      val, err := strconv.Atoi(tok.val)
      if err != nil {
        Error(fmt.Sprintf("Invalid PUSH value: %v\n", tok.val), fmt.Sprintf("%v %v", tok.type_, tok.val))
        return
      }
      s.Push(val)
    case "PULL":
      s.Pull()
    case "READ":
      s.Read()

    case "DADV":
      s.DeAdvance()
    case "ADV":
      s.Advance()

    case "REMOVE":
      s.Remove()

    case "PRINTS":
      s.PrintS()
    case "PRINT":
      s.Print(tok.val)
    case "LEN":
      s.Len()

    case "ADD":
      s.Add()
    case "SUB":
      s.Sub()

    case "IDX":
      s.Idx()

    case "LABEL":
      continue

    case "JUMP":
      label, ok := labels[tok.val]
      if !ok {
        Error(fmt.Sprintf("Label '%v' doesn't exist!", tok.val), fmt.Sprintf("%v %v", tok.type_, tok.val))
        return
      }
      Interpret(label.program, s, labels)
      return

    case "END":
      os.Exit(0)
    case "ENDL":
      os.Exit(0)
    }
  }
}
