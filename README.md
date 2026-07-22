# Programming-Language

A custom, tree-walking interactive interpreter, built on Go and ANTLR.

## Running the language

There are two ways to run the program: interactive shell (REPL) and file script execution.

### Interactive REPL mode

Launch the interpreter with no arguments to start a terminal session with a presistent command line history (use up/down arrows to navigate) and dynamic multi-line indentation ('...'). If you want to exit, you can type 'exit' or CTRL+c. If you want to just stop a multi-line block, you can CTRl+c.

```sh
  go run main.go
```

```sh
  go run main.go <input_filename>
```

Testing can be done from the testing directory running

```sh
  cd testing
  go test
```
