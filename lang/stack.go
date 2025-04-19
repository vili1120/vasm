package lang

import (
	"fmt"
	"strconv"
)

func NewStack(length int) *Stack {
  cells := make([]int, length)
  s := &Stack{
    sp: -1,
    cells: cells,
  }
  s.Advance()
  return s
}

type Stack struct {
  sp int
  cells []int
}

func (s *Stack) Advance() {
  s.sp++
  if s.sp >= len(s.cells) {
    Error("Stack overflow!")
    s.sp--
  }
}

func (s *Stack) DeAdvance() {
  s.sp--
  if s.sp < 0 {
    Error("Stack underflow!")
    s.sp++
  }
}

func (s *Stack) Remove() {
  s.cells[s.sp] = 0
}

func (s *Stack) Push(val int) {
  s.cells[s.sp] = val
  s.Advance()
}

func (s *Stack) Pull() {
  fmt.Println(s.cells[s.sp])
}

func (s *Stack) Read() {
  var i string
  fmt.Scanln(&i)
  val, err := strconv.Atoi(i)
  if err != nil {
    Error("Invalid input!")
  }
  s.Push(val)
}

func (s *Stack) PrintS() {
  fmt.Println(s.cells)
}

func (s *Stack) Print(idx string) {
  val, err := strconv.Atoi(idx)
  if err != nil {
    Error("Invalid index!")
  }
  if val < len(s.cells) && val > 0 {
    fmt.Println(s.cells[val])
  } else {
    Error("Index out of range!")
  }
}

func (s *Stack) Len() {
  fmt.Println(len(s.cells))
}

func (s *Stack) Add() {
  s.DeAdvance()
  if s.sp-1 < 0 {
    Error("Index out of range!")
  } else {
    s.cells[s.sp-1] = s.cells[s.sp-1]+s.cells[s.sp]
    s.Remove()
    s.DeAdvance()
  }
}

func (s *Stack) Sub() {
  s.DeAdvance()
  if s.sp-1 < 0 {
    Error("Index out of range!")
  } else {
    s.cells[s.sp-1] = s.cells[s.sp-1]-s.cells[s.sp]
    s.Remove()
    s.DeAdvance()
  }
}

func (s *Stack) Idx() {
  fmt.Println(s.sp)
}
