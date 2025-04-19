package lang

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(prompt string) string {
	scanner := bufio.NewReader(os.Stdin)
  fmt.Print(prompt)
	line, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err)
	}
  return line
}

func InputFull() []string {
	var str []string

	for {
    line := Input("> ")
		line = strings.ToUpper(strings.TrimSpace(line))
    if strings.Split(line, " ")[0] == "LABEL" {
      for {
        line = Input("--> ")
        line = strings.ToUpper(strings.TrimSpace(line))
        if line == "ENDL" {
          str = append(str, line)
          break
        }
        str = append(str, line)
	    }
    }
		if line == "END" {
			break
		}

		str = append(str, line)
	}

	return str
}

func InputCmd() []string {
	var str []string
	
  line := Input("> ")

  line = strings.ToUpper(strings.TrimSpace(line))
  if strings.Split(line, " ")[0] == "LABEL" {
    for {
      line := Input("--> ")
      line = strings.ToUpper(strings.TrimSpace(line))
      if line == "ENDL" {
        str = append(str, line)
        break
      }
      str = append(str, line)
	  }
  }
  if line == "END" {
    os.Exit(0)
  }

	str = append(str, line)

	return str
}
