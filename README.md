# vasm

**vasm** is a lightweight, assembly-like, stack-based interpreter written in Go. It allows you to define and run simple instructions using a stack-oriented execution model.

> 🧠 Inspired by virtual machines and low-level computation, vasm is a great starting point for experimenting with custom interpreters or learning stack-based execution principles.

---

## 🚀 Features

- Stack-based execution
- Simple, assembly-style instruction set
- Easy to extend with new operations
- Written entirely in Go for speed and portability

---

## 🛠️ Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18+ recommended)

### Clone the Repository

```bash
git clone https://github.com/vili1120/vasm.git
cd vasm
git checkout dev
```

### Build & Run

You can run the interpreter directly:

```bash
go run main.go
```

Or build it:

```bash
go build -o vasm
./vasm
```

---

## 💡 Example Instructions

Below is an example of how a `vasm` program might look (assuming some standard instructions):

```
PUSH 10
PUSH 20
ADD
PRINT
```

This would push 10 and 20 to the stack, add them, and print the result (`30`).

---

## 📦 Instruction Set

Some of the current instructions supported (check `main.go` for latest):

- `PUSH <value>` – Push an integer value onto the stack
- `POP` – Pop the top value from the stack
- `ADD`, `SUB`, `MUL`, `DIV` – Basic arithmetic operations
- `PRINT` – Prints the top of the stack

> You can easily add more instructions by modifying the instruction parser in `main.go`.

---

## 🧪 Testing

> Currently, there are no automated tests — but contributions are welcome!

You can manually test by editing the `code` slice in `main.go` with your own instruction sets.

---

## 📂 Project Structure

```
vasm/
├── go.mod         # Go module definition
└── main.go        # Core interpreter logic
```

---

## 🧱 Future Improvements

- REPL (interactive shell)
- File-based program loading
- More robust error handling
- Instruction macros or custom functions
- Testing framework

---

## 🤝 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you'd like to add or improve.

---

## 📄 License

[MIT](LICENSE)

---

## ✨ Author

Created with ❤️ by [@vili1120](https://github.com/vili1120)
