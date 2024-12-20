package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	version = "unset"
	commit  = "unset"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-v" {
		fmt.Printf("%s-%s", version, commit)
		os.Exit(0)
	}

	b := &strings.Builder{}
	_, err := io.Copy(b, os.Stdin)
	if err != nil {
		fmt.Println(err)
	}

	input := b.String()
	input = strings.ReplaceAll(input, " ", "")
	input = strings.TrimSpace(input)

	if !isValidHex(input) {
		log.Fatal("not a valid hex string")
	}

	fmt.Print(strings.Join(getWordsForBytesString(input), " "))
}

func getWordsForBytesString(input string) []string {

	outputStrings := []string{}

	hexBytes, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}

	for i, b := range hexBytes {
		outputStrings = append(outputStrings, wordlist[b][i%2])
	}

	return outputStrings
}

func isValidHex(s string) bool {
	// Regular expression to match a valid hex string
	// ^[0-9a-fA-F]+$ means the string can only contain hex characters
	re := regexp.MustCompile(`^[0-9a-fA-F]+$`)
	return re.MatchString(s)
}
