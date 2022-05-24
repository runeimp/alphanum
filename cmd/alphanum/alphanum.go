package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/runeimp/alphanum"
)

const (
	AppLabel   = AppName + " v" + AppVersion
	AppName    = "AlphaColumn"
	AppVersion = "0.1.0"
	CLIName    = "alphanum"
	helpInfo   = `
%s [OPTIONS] LETTERS_OR_INDEXES

%s accepts one or more column letters or indexes and converts between the two

OPTIONS
  -h, --help     Display this help info
  -v, --version  Display the app version

EXAMPLE
$ %s ABC 703
ABC = 731
703 = AAA

`
)

func main() {

	cmd := "help"

	if len(os.Args) > 1 {
		cmd = strings.ToLower(os.Args[1])
	}

	switch cmd {
	case "-h", "-hlp", "-help", "--help", "hlp", "help":
		fmt.Println(AppLabel)
		fmt.Printf(helpInfo, CLIName, AppName, CLIName)
		os.Exit(0)
	case "-v", "-ver", "-version", "--version", "ver", "version":
		fmt.Println(AppLabel)
		os.Exit(0)
	default:
		// fmt.Fprintf(os.Stderr, "Unknown command or option %q\n", cmd)
		result := alphanum.ParseSlice(os.Args[1:])
		for i, arg := range os.Args[1:] {
			fmt.Printf("%s = %s\n", strings.ToUpper(arg), result[i])
		}
		// n := alphanum.ColumnLettersToNumbers(os.Args[1])
		// fmt.Printf("main() | n: %d\n", n)
	}
}
