package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var (
	cost      int
	noNewline bool
)

func init() {
	cost = 14
	costHelp := fmt.Sprintf("Hashing cost. Default is %d", cost)
	flag.IntVar(&cost, "c", cost, costHelp)
	flag.IntVar(&cost, "cost", cost, costHelp)
	flag.BoolVar(&noNewline, "n", false, "Do not print the trailing newline character.")
	flag.Parse()
}

func main() {
	password, err := ioutil.ReadAll(os.Stdin)
	hash, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	fmt.Print(string(hash))

	if !noNewline {
		fmt.Println()
	}
}
