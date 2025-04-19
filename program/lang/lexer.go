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
        Error("No argument for PUSH!")
      }
    case "PULL":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for PULL!")
      }
    case "READ":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for READ!")
      }

    case "DADV":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for DADV!")
      }
    case "ADV":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for ADV!")
      }

    case "REMOVE":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for REMOVE!")
      }

    case "PRINTS":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for PRINTS!")
      }
    case "PRINT":
      if arg != "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("No argument for PRINT!")
      }

    case "LEN":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for LEN!")
      }

    case "ADD":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for ADD!")
      }
    case "SUB":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for SUB!")
      }

    case "IDX":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for IDX!")
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
        Error("Too many arguments for LABEL!")
      }
    case "JUMP":
      if arg != "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("No argument for JUMP!")
      }

    case "END":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for END!")
      }
    case "ENDL":
      if arg == "" {
        toks = append(toks, Tok{op, arg, nil})
      } else {
        Error("Too many arguments for ENDL!")
      }
    case "":
      continue
    default:
      Error("Expected a command or END!")
    }
  }

  return toks
}

