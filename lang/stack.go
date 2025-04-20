package lang

import (
	"fmt"
	"log"
	"strconv"
)

func NewStack(length int) *Stack {
  s := &Stack{
    Sp: 0,
    Cells: make([]int, length),
  }
  return s
}

type Stack struct {
  Sp int
  Cells []int
}

func (s *Stack) Advance() {
  s.Sp++
  if s.Sp > (len(s.Cells)-1) {
    s.Sp--
  }
}

func (s *Stack) DeAdvance() {
  s.Sp--
  if s.Sp < 0 {
    s.Sp = 0
  }
}

func (s *Stack) Push(arg string) {
  val, err := strconv.Atoi(arg)
  if err != nil {
    log.Fatal(err)
    return
  }
  s.Cells[s.Sp] = val
  s.Advance()
}

func (s *Stack) Pull() {
  fmt.Println(s.Cells[s.Sp])
}

func (s *Stack) Read() {
  var i string
  fmt.Scanln(&i)

  s.Push(i)
}

func (s *Stack) Pop() {
  s.Cells[s.Sp] = 0
}

func (s *Stack) Prints() {
  fmt.Println(s.Cells)
}

func (s *Stack) Print(idx string) {
  val, err := strconv.Atoi(idx)
  if err != nil {
    log.Fatal(err)
    return
  }
  if val > len(s.Cells)-1 || val < 0 {
    fmt.Println("Index out of range")
    return
  }
  fmt.Println(s.Cells[val])
}
