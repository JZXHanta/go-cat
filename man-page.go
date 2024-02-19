package main

var manPage string = `NAME
		go-cat -- concatenate and print files

SYNOPSIS
		go-cat [-b , -n] [file ...]

DESCRIPTION
		The go-cat utility reads files sequentially, writing them to the standard
		output.  The file operands are processed in command-line order.  If file is
		a single dash ('-'), go-cat reads from the standard input.

		The options are as follows:

			-b      Number the non-blank output lines, starting at 1.

			-n      Number the output lines, starting at 1.
			
EXIT STATUS
		The go-cat utility exits 0 on success, and >0 if an error occurs.
		
EXAMPLES
		The command:

			go-cat file1

		will print the contents of file1 to the standard output.

		The command:

			go-cat file1 file2 > file3

		will sequentially print the contents of file1 and file2 to the file
		file3, truncating file3 if it already exists.  See the manual page for
		your shell (e.g., sh(1)) for more information on redirection.`
