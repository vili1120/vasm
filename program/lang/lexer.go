package lang

import "strings"

func Lex(text []string) []Tok {
  toks := []Tok{}

  for i, line := range text {
    cline := strings.Split(line, " ")
    op := cline[0]
    arg := ""
    if len(cline) > 1 {
      arg = cline[1]
    }
    op = strings.ToUpper(op)
    if strings.HasPrefix(op, "#") ||strings.HasPrefix(op, "//") {
      continue
    }
    switch op {
    case "PUSH":
      if arg != "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("No argument for PUSH!", line)
      }
    case "PULL":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for PULL!", line)
      }
    case "READ":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for READ!", line)
      }

    case "DADV":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for DADV!", line)
      }
    case "ADV":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for ADV!", line)
      }

    case "REMOVE":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for REMOVE!", line)
      }

    case "PRINTS":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for PRINTS!", line)
      }
    case "PRINT":
      if arg != "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("No argument for PRINT!", line)
      }

    case "LEN":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for LEN!", line)
      }

    case "ADD":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for ADD!", line)
      }
    case "SUB":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for SUB!", line)
      }

    case "IDX":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for IDX!", line)
      }

    case "LABEL":
      if arg != "" {
        lprogram := []string{}
        i++
        for i < len(text) {
          line := strings.TrimSpace(text[i])
          if strings.ToUpper(line) == "ENDL" {
            break
          }
          lprogram = append(lprogram, line)
          i++
        }

        ltoks := Lex(lprogram)
        toks = append(toks, Tok{op, arg, &Label{arg, ltoks}})
      } else {
        Error("Too many arguments for LABEL!", line)
      }
    case "JUMP":
      if arg != "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("No argument for JUMP!", line)
      }

    case "END":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for END!", line)
      }
    case "ENDL":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for ENDL!", line)
      }
    case "":
      continue
    default:
      Error("Expected a command or END!", line)
    }
  }

  return toks
}

