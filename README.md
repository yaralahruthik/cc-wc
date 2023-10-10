# Write Your Own wc Tool

This challenge was to build my own version of the Unix command line tool wc!

[Challenge Link](https://codingchallenges.fyi/challenges/challenge-wc/)

## Description

This cli tool is a part replica of `wc` command in linux. It counts the number of lines, words and characters in a file.

## Usage

```bash
$ go run main.go <filename>
```

Also, supports reading from standard input if no file is provided to read.

```bash
$ cat abcd | go run main.go
```

## Flags

- `-l` : Prints the number of lines in the file.
- `-w` : Prints the number of words in the file.
- `-c` : Prints the number of bytes in the file.
- `-m` : Prints the number of characters in a file. If the current locale does not support multibyte characters this will match the -c option

If no flag is provided, number of lines, words, and bytes are printed.
