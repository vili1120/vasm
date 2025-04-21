# VASM

**VASM** is a minimalist low-level interpreter written in Go. It provides a stack-cell-based memory model and a straightforward instruction set for low-level programming experimentation or educational purposes.

## 🚀 Features

- Stack-pointer-based memory manipulation
- Arithmetic operations between adjacent cells
- Conditional execution
- Label-based control flow with jumping
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

### 🧮 Simple Cell Manipulation
| Instruction       | Description                                             |
|------------------|---------------------------------------------------------|
| `PUSH <value>`   | Adds a value to the current cell and advances           |
| `PULL`           | Prints the value of the current cell                    |
| `READ`           | Reads a value from user input and pushes it             |
| `POP`            | Resets the current cell to zero                         |

### ➕ Arithmetic Operations
| Instruction       | Description                                             |
|------------------|---------------------------------------------------------|
| `ADD`            | Adds the current cell to the previous one               |
| `SUB`            | Subtracts the current cell from the previous one        |
| `MUL`            | Multiplies the current and previous cell values         |
| `DIV`            | Divides the previous cell by the current one            |

### 🔁 Labels & Control Flow
| Instruction         | Description                                           |
|--------------------|-------------------------------------------------------|
| `LABEL <name>`     | Defines a label block (ends with `END`)              |
| `JUMP <label>`     | Jumps to a label (skips all following instructions)  |

### 🔎 Conditionals
| Instruction       | Description                                             |
|------------------|---------------------------------------------------------|
| `IF.EQ <value>`  | Executes inner block if current cell equals value       |
| `IF.NE <value>`  | Executes inner block if current cell does not equal value |

### 📤 Output
| Instruction       | Description                                             |
|------------------|---------------------------------------------------------|
| `PRINTS`         | Prints the full stack                                   |
| `PRINT`          | Prints a cell based on index                            |

### 📍 Pointer Movement & Program Control
| Instruction       | Description                                             |
|------------------|---------------------------------------------------------|
| `ADV`            | Advances the stack pointer                              |
| `DADV`           | Moves the pointer backward                              |
| `END`            | Ends execution                                           |

---

## ✍️ Syntax Notes

- Comments start with `#` and are ignored.
- Labels must end with an `END`.
- Indentation is optional but recommended for clarity.

---

## 📂 Project Structure

```
vasm/
├── lang/              # Core logic for parsing and executing
│   └── ...
├── test.vasm          # Sample program
├── test2.vasm         # Another sample
├── instructions.txt   # Detailed instruction set
├── main.go            # Entry point of the interpreter
└── go.mod             # Go module file
```

---

## 📄 Example Program

```vasm
# Adds two numbers and prints the result
PUSH 5
PUSH 3
DADV
ADD
PULL
```

---

## 🤝 Contributing

Pull requests are welcome! Please open an issue to discuss improvements or ideas.

---

## 📜 License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for more info.
