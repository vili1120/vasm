# VASM

**VASM** is a minimalist low-level interpreter written in Go. It provides a stack-cell-based memory model and a straightforward instruction set for low-level programming experimentation or educational purposes.

## 🚀 Features

- Stack-pointer-based memory manipulation
- Arithmetic operations between adjacent cells
- Conditional execution
- Label-based control flow with jumping
- Simple Raylib bindings
- Simple I/O capabilities
- Clear and readable syntax

---

## 📦 Installation

### Prerequisites
- [Go](https://golang.org/dl/) 1.16 or higher

### Steps
```bash
git clone https://github.com/vili1120/vasm.git
cd vasm
cd program
go build -o vasm
```

---

## 🛠️ Usage

```bash
./vasm path/to/yourfile.vasm
```

Example:
```bash
./vasm test.vasm
```

---

## 🧾 Instruction Set

View [instructions.norg](instructions/instructions.norg) for instruction set

>> instructions.txt was removed due to being outdated and not being updated alongside the norg version

---

## ✍️ Syntax Notes

- Comments start with `#` and are ignored.
- Labels must end with an `END`.
- Indentation is optional but recommended for clarity.

---

## 📂 Project Structure

```
vasm/
|-- README.md
|-- instructions/               # Different instruction sets
|   |-- instructions.norg
|   `-- instructions.txt
|-- program/                    # The main logic of the interpreter
|   |-- ...
`-- tests/                      # Tests for testing the capabilities of the interpreter
    |-- test.vasm
    |-- test2.vasm
    |-- test3.vasm
    `-- test4.vasm
```

---

## 📄 Example Programs

```vasm
# Adds two numbers and prints the result
PUSH 5
PUSH 3
DADV
ADD
PULL
END
```

```vasm
# This script creates a simple raylib program

# The width of the window
PUSH 800

# The height of the window
PUSH 600

# The target fps
PUSH 60

# The color white for RL:CLEAR
PUSH 1

RL:INIT 0 1
RL:FPS 2

RL:FORCLOSE
    RL:BEGINDRAWING
    
    RL:CLEAR 3

    RL:ENDDRAWING
ROF

RL:EXEC

END
```

---

## 🤝 Contributing

Pull requests are welcome! Please open an issue to discuss improvements or ideas.

---

## 📜 License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for more info.
