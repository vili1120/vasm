package lang

import "fmt"

func Error(description string, line string) {
  fmt.Printf("Error: %v\n", description)
  fmt.Printf("  %v\n", line)
  return
}

type Tok struct {
  type_ string
  val string
  label *Label
}

type Label struct {
  name string
  program []Tok
}
