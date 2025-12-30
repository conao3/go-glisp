# go-glisp

A Lisp interpreter written in Go. This project implements core Lisp primitives including lambda expressions, list operations, and a REPL for interactive evaluation.

## Features

- S-expression parser with proper handling of dotted pairs
- Core Lisp primitives: `atom`, `eq`, `car`, `cdr`, `cons`, `cond`, `quote`
- Lambda expressions with lexical scoping
- Variable and function bindings via `set` and `fset`
- Interactive REPL with configurable evaluation stages

## Installation

```bash
go install github.com/conao3/go-glisp@latest
```

Or build from source:

```bash
git clone https://github.com/conao3/go-glisp.git
cd go-glisp
go build
```

## Usage

Start the REPL:

```bash
go-glisp
```

Run with a file:

```bash
go-glisp -i script.lisp
```

### Command-line Options

| Flag | Description |
|------|-------------|
| `-i <file>` | Read input from file instead of stdin |
| `-sr` | Stop after reader stage (show parsed AST) |
| `-se` | Stop after evaluator stage (show result) |

## Examples

```lisp
glisp> (cons 1 2)
(1 . 2)

glisp> (car (cons 1 2))
1

glisp> (cdr (cons 1 2))
2

glisp> (set (quote x) 10)
10

glisp> ((lambda (a b) (cons a b)) 1 2)
(1 . 2)
```

### Defining Functions

```lisp
glisp> (fset (quote pair) (lambda (a b) (cons a b)))
(lambda (a b) (cons a b))

glisp> (pair 1 2)
(1 . 2)
```

## Built-in Functions

| Function | Description |
|----------|-------------|
| `atom` | Returns `t` if the argument is an atom, `nil` otherwise |
| `eq` | Returns `t` if two symbols are equal |
| `car` | Returns the first element of a cons cell |
| `cdr` | Returns the second element of a cons cell |
| `cons` | Constructs a new cons cell |
| `cond` | Conditional expression |
| `quote` | Returns its argument unevaluated |
| `lambda` | Creates an anonymous function |
| `set` | Binds a value to a symbol |
| `fset` | Binds a function to a symbol |

## Project Structure

```
go-glisp/
├── main.go           # Entry point and CLI handling
├── reader/           # S-expression parser
├── evaluator/        # Expression evaluator
├── repl/             # Read-Eval-Print loop
└── types/            # Core data types (Int, Symbol, Cons, Environment)
```

## License

MIT
