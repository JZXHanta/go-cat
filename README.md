# go-cat

#### A clone of the `cat` command used in Unix/Linux command lines

---

- With no additional arguments other than file-name (or a space separated list of filenames), prints contents of input file(s) to standard out.

  - example: `go-cat example.txt`

- If ran with the argument '-' and no filename it will print from standard input

  - example: `echo "hello, world" | go-cat -`

- If ran with argument '-n', will print contents of input with line numbers on all lines

  - example: `go-cat -n example.txt`

- If ran with argument '-b', will print contents of input with line numbers on all lines, excluding blank lines

  - example: `go-cat -b example.txt`
