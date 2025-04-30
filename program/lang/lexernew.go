package lang

func NewLexerNew(text string) *LexerNew {
  ln := &LexerNew{
    Text: text,
    Idx: -1,
  }
  ln.Advance()
  return ln
}

type LexerNew struct {
  Text string
  Idx int
  CurrentChar rune
}

func (l *LexerNew) Advance() {
  l.Idx++
  if l.Idx < len(l.Text) {
    l.CurrentChar = rune(l.Text[l.Idx])
  }
}

func (l *LexerNew) MakeTokens() []Token {
  toks := []Token{}

  for {
    switch l.CurrentChar {
    case ';', '\n':
      toks = append(toks, NewToken(NEWLINE, ""))
      l.Advance()
    }
  }
}
