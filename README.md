# Golang Fast Start

A fast-paced, practical walkthrough of Go's core concepts. Not a beginner hand-holding tutorial — this project is built to get you productive in Go quickly by covering real patterns and deeper mechanics through runnable code examples.

### Resources

- [A Tour of Go](https://go.dev/tour/list)
- [Learn GO Fast: Full Tutorial](https://www.youtube.com/watch?v=8uiZC0l4Ajw)
- [GeeksforGeeks - Go Programming Language](https://www.geeksforgeeks.org/go-programming-language-introduction/)

## Modules

### [01 - Variables, Types & Constants](01-variables/main.go)

Covers all declaration styles (explicit, short `:=`, multiple), type casting between numeric types, and why Go is strict about it. Explores how `string` length reports bytes not characters, and how `rune` (alias for `int32`) properly represents Unicode code points. Demonstrates default zero values for every built-in type and constant immutability.

### [02 - Functions & Error Handling](02-functions/main.go)

Go functions return multiple values — this module shows why that matters by implementing the idiomatic error handling pattern. Instead of exceptions, Go returns an `error` as the last value and the caller decides how to handle it. Covers `errors.New`, nil checks, and how this pattern shapes the way Go code is structured.

### [03 - Arrays, Slices, Maps & Loops](03-arrays-slices-maps-loops/main.go)

Explains the difference between fixed-size arrays (value types, contiguous memory) and slices (dynamic, backed by arrays with length and capacity). Demonstrates how slice capacity doubles on growth and why pre-allocating with `make` matters — includes a benchmark showing ~5x performance gain on 1M elements. Covers maps with key existence checks, deletion, and unordered iteration. All loop variants: classic `for`, `range`, while-style, and map iteration.

### [04 - Strings, Runes & Bytes](04-strings-runes-bytes/main.go)

Digs into how Go handles text under the hood. Strings are immutable byte slices encoded in UTF-8, meaning indexing gives you bytes, not characters. Shows how multi-byte characters like `é` break naive indexing and why converting to `[]rune` solves it. Compares string concatenation via `+` (allocates a new string every time) vs `strings.Builder` (efficient buffer-based approach).

### [05 - Structs & Interfaces](05-structs-interfaces/main.go)

Defines custom types with structs, including nested structs and field access. Attaches behavior through method receivers. Introduces interfaces as implicit contracts — any type that implements the required methods satisfies the interface, no `implements` keyword needed. Demonstrates polymorphism by passing different struct types to a single function through a shared interface.

### [06 - Pointers](06-pointers/main.go)

Covers pointer mechanics: `&` to get an address, `*` to dereference, `new()` for heap allocation. Shows the real impact of pass-by-value (arrays are copied) vs pass-by-reference (pointers modify the original). Explains nil pointer dereferencing risks and why slices behave like references even without explicit pointers — they share the underlying array.

## Quick Start

```bash
# Run any module
go run 01-variables/main.go
go run 02-functions/main.go
# ...and so on
```

## Requirements

- Go 1.25+
