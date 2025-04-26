package lang

func Run(text []string, length int) {
  l := NewLexer(text)
  toks := l.MakeInstr()
  
  s := NewStack(length)
  i := NewInterpreter(toks, s)
  i.Interpret()
}
