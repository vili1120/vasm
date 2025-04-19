# vasm

**vasm** is a lightweight, assembly-like, stack-based interpreter written in Go. It allows you to define and run simple instructions using a stack-oriented execution model.

---

##   Stack-based Execution

- Stack: an array of cells
- Cell: contains an integer value
- Stack pointer: the index of the current cell

- Instructions operate on the current cell

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
```
For development branch:

```bash
git checkout dev
```

### Build & Run

##### You can run the interpreter directly:

```bash
go run program
```

##### Or build it:

```bash
go build -o vasm
./vasm
```

###### For linux you can use:

```bash
./build.sh
```

---

## 💡 Example Instructions

Below is an example of how a `vasm` program might look (assuming some standard instructions):

```
PUSH 10
PUSH 20
ADD
PULL
END
```

This would push 10 and 20 to the stack, add them, and print the result (`30`).

---

## 📦 Instruction Set

> Note: instructions are not case-sensitive

Some of the current instructions supported (check `main.go` for latest):

- `PUSH <value>` – Push an integer value onto the current cell
- `PULL` - Prints the current cell's value
- `READ` - Reads user input for a value and pushes it to the current cell

- `ADV` - Advances the stack pointer by 1
- `DADV` - Deadvances the stack pointer by 1

- `REMOVE` – Remove the value from the current cell

- `ADD`, `SUB` – Basic arithmetic operations

- `PRINTS` - Prints the whole stack
- `PRINT <idx>` – Prints a cell's value based on the index from the stack

- `LEN` - Prints the length of the stack
- `IDX` - Prints the stack pointer's index

- `END` - Ends the program

> You can easily add more instructions by modifying the instruction parser in `main.go`.

---

## 📂 Project Structure

```
vasm/
|-- README.md                   # Project readme file
|-- build.sh                    # Script to build for linux
|-- program/                    # Directory containing the go module
|   |-- go.mod
|   |-- lang/                   # Directory containing main logic for interpreter
|   |   |-- input.go
|   |   |-- interpreter.go
|   |   |-- lexer.go
|   |   |-- stack.go
|   |   `-- types.go
|   `-- main.go                 # Brings every package together
|-- test.vasm                   # Test file to run after installation
`-- uninstall.sh                # Uninstalls the vasm binary from path
```

---

## 🧱 Future Improvements

- More robust error handling(1/3)
- Labels/functions

---

## 🤝 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you'd like to add or improve.

---

## 📄 License

[MIT](LICENSE)

---

## ✨ Author

Created with ❤️ by [@vdev](https://github.com/vili1120)
