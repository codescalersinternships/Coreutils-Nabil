# Unix-Based Commands Implemented in Go
Goal is to reimplement several core Unix utilities.


## Operations Implemented

### 1. head

The `head` command in Go prints the first N lines of a file. Implemented with:
```bash
go run main.go -operation=head -directory=input.txt -isFile=true -n=10
```

### 2. tail

The `tail` command in Go prints the last N lines of a file. Implemented with:
```bash
go run main.go -operation=tail -directory=input.txt -isFile=true -n=10
```

### 3. wc

The `wc` command in Go counts lines, words, and bytes in a file and prints the counts.
```bash
go run main.go -operation=wc -directory=input.txt -isFile=true -l=true -w=true -c=true
```

### 4. cat

The `cat` command in Go concatenates files and prints them to standard output. Implemented with:
- `cat -n`: Prints each line with line numbers.
```bash
go run main.go -operation=cat -directory=input.txt
```

### 5. echo

The `echo` command in Go prints its arguments to standard output.
```bash
go run main.go -operation=echo -directory=input.txt -0=true
```

### 6. tree

The `tree` command in Go recursively lists contents of directories in a tree-like format.
```bash
go run main.go -operation=tree -directory=input.txt
```

### 7. env

The `env` command in Go lists the current environment variables and their values.
```bash
go run main.go -operation=env -directory=input.txt
```

### 8. true

The `true` command in Go does nothing except return a successful exit status.
```bash
go run main.go -operation=true
```

### 9. false

The `false` command in Go does nothing except return a non-successful exit status.
```bash
go run main.go -operation=false
```
