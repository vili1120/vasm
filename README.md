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

View [instructions.norg](instructions.norg) or [instructions.txt](instructions.txt) for instruction set

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
