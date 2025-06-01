# GoChain 🦾

A minimal, educational blockchain implementation written in Go.  
Easily add transactions, explore the chain, and learn the basics of blockchain technology—all from your terminal.

---

## 🚀 Features

- **Simple CLI** — Add transactions, list the full blockchain as JSON, and more.
- **Persistent Storage** — Blockchain data is saved to a file (`GoChain.json`).
- **Genesis Block** — Automatically created on first run.
- **Custom Transactions** — Specify sender, receiver, and value.

---

## 🛠️ Usage

### 1. Clone & Build

```sh
git clone git@github.com:jean0t/gochain.git
cd gochain
go build -o gochain # or acces the binary from /dist
```

### 2. Add a Transaction

```sh
./gochain add --from alice --to bob --value 42.5
```

### 3. List the Blockchain

```sh
./gochain list
```
This prints the whole blockchain as formatted JSON.

---

## ⚙️ Commands

| Command                 | Description                                 |
|-------------------------|---------------------------------------------|
| `add`                   | Add a new transaction                       |
| `list`                  | List all blocks in the chain (as JSON)      |

#### Flags for `add`:
- `--from`    — Sender address (string, required)
- `--to`      — Receiver address (string, required)
- `--value`   — Value to transfer (float, required)

---

## 🧩 Example Output

```json
[
  {
    "parent": "000...000",
    "timestamp": 1685555555555,
    "data": {
      "from": "jean0t",
      "to": "jean0t",
      "value": 30000
    },
    "signature": "b1f8a7..."
  },
  {
    "parent": "b1f8a7...",
    "timestamp": 1685555556666,
    "data": {
      "from": "alice",
      "to": "bob",
      "value": 42.5
    },
    "signature": "a2d3e4..."
  }
]
```

---

## 📦 Project Structure

```
gochain/
├── cmd/           # CLI commands (Cobra)
├── internal/      # Blockchain logic
├── main.go
├── LICENSE
└── README.md
```

---

## 🧑‍💻 For Developers

- Written in idiomatic Go.
- Easily extensible—add new CLI commands, consensus logic, or networking!

---

## 💡 Why?

This project is for learning and experimentation.  
It’s not intended for use as a production-ready blockchain.

---

## 📜 License

MIT License

---

> Made with ❤️ by jean0t
