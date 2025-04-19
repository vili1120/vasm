package lang

import "fmt"

func Error(description string) {
  fmt.Printf("Error: %v\n", description)
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
