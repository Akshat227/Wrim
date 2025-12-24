# Wrim

Wrim is a small, educational lexer implementation written in Go. It tokenizes a tiny, Monkey-inspired language with support for identifiers, integers, basic operators, delimiters, and a couple of keywords (`let`, `fn`). The project is organized for clarity and maintainability and is a good starting point for learning how interpreters and compilers begin with lexical analysis.

**Repository structure**

- `lexer/` — lexer implementation and unit tests.
- `token/` — token definitions and keyword lookup.
- `go.mod` — Go module file for dependency management.

## Features

- Tokenizes identifiers, integers, and punctuation: `=`, `+`, `,`, `;`, `(`, `)`, `{`, `}`.
- Recognizes keywords `let` and `fn`.
- Lightweight and easy to extend (add more tokens, multi-char operators, strings, comments, etc.).
- Optional debug printing of tokens (toggle with `SetDebug(true)`).

## Quickstart

Prerequisites:

- Go 1.18+ installed and on your `PATH`.

Clone the repo and run the tests:

```bash
git clone <your-repo-url>
cd Wrim
go test ./...
```

To run the lexer manually in code, see a minimal example:

```go
package main

import (
    "Wrim/lexer"
    "fmt"
)

func main() {
    input := `let five = 5;`
    l := lexer.New(input)
    // Enable debug printing if you want to see tokens printed to stdout
    l.SetDebug(true)

    for {
        tok := l.NextToken()
        fmt.Printf("%+v\n", tok)
        if tok.Type == "EOF" {
            break
        }
    }
}
```

## API notes

- `lexer.New(input string) *Lexer` — create a new lexer instance.
- `(*Lexer).NextToken() token.Token` — return the next token from input.
- `(*Lexer).SetDebug(bool)` — toggle token printing for debugging and development.

## Testing & Development

Run unit tests with:

```bash
go test ./...
```

Formatting and vetting:

```bash
gofmt -w .
go vet ./...
```

## Extending the lexer

- To add multi-character operators (e.g. `==`, `!=`) implement a lookahead in `NextToken`.
- To support strings, add a `readString` helper and a `STRING` token.
- To support comments, extend `skipWhitespace` or add a `skipComment` helper invoked when `/` or `#` is seen.

## Contributing

Contributions are welcome. Please open issues for bugs or feature requests and submit pull requests with focused, well-documented changes.

## License

This repository is provided under the MIT License — include a `LICENSE` file if you want an explicit license.

---

If you'd like, I can also:

- Add a `NewWithDebug(input string, debug bool)` constructor.
- Add a small example program under `cmd/` demonstrating the lexer in action.

Pick one and I will implement it.
