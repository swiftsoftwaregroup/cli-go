/*
Package main implements a CLI tool for greeting users in different languages.

cli-go is a command-line interface application that allows users to generate
greetings in various languages. It reads a name from a file and can produce
greetings in English, Spanish, or Bulgarian.

Usage:

The main command of the application is "greet", which takes a file path as an
argument and an optional language flag:

	cli-go greet [file] [--language language]

Where [file] is the path to a text file containing a name, and [--language] is
an optional flag to specify the greeting language (default is English).

Supported languages:
  - English (en)
  - Spanish (es)
  - Bulgarian (bg)

Example usage:

	echo "John" > name.txt
	cli-go greet name.txt
	cli-go greet --language es name.txt
	cli-go greet -l bg name.txt

The application uses the Cobra library for CLI structure and argument parsing.
It demonstrates basic file I/O, string manipulation, and modular CLI design in Go.
*/
package main
