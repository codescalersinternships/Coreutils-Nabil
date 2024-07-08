# Unix-Based Commands Implemented in Go
Goal is to reimplement several core Unix utilities.


## Operations Implemented

### 1. head

The `head` command in Go prints the first N lines of a file. Implemented with:
```bash
go run ./cmd/head -n=2 input.txt
```

### 2. tail

The `tail` command in Go prints the last N lines of a file. Implemented with:
```bash
go run ./cmd/tail -n=2 input.txt
```

### 3. wc

The `wc` command in Go counts lines, words, and bytes in a file and prints the counts.
```bash
go run ./cmd/wc -l=true -w=true -c=true input.txt
```

### 4. cat

The `cat` command in Go concatenates files and prints them to standard output. Implemented with:
- `cat -n`: Prints each line with line numbers.
```bash
go run ./cmd/cat -n input.txt
```

### 5. echo

The `echo` command in Go prints its arguments to standard output.
```bash
go run ./cmd/echo -n=true hello
```

### 6. tree

The `tree` command in Go recursively lists contents of directories in a tree-like format.
```bash
go run ./cmd/tree -l=2 ./cmd/
```

### 7. env

The `env` command in Go lists the current environment variables and their values.
```bash
go run ./cmd/env
```

### 8. true

The `true` command in Go does nothing except return a successful exit status.
```bash
go run ./cmd/true
```

### 9. false

The `false` command in Go does nothing except return a non-successful exit status.
```bash
go run ./cmd/false
```
