package main

import (
	checker "asciiart/checker"
	print "asciiart/print"
	split "asciiart/split"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if checker.IsValidInput(args) {
		colorName := strings.TrimPrefix(args[0], "--color=")
		result := asciiArt(colorName, args[1], args)
		fmt.Print(result)
		fmt.Print(string("\033[0m"))
	}
}

func asciiArt(colorName string, input string, args []string) string {
	var result string
	// =============== READ INPUT FILE & HANDLE ERROR ===============
	src, err := os.ReadFile("./banners/standard.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// =============== SLICE BANNER AND INPUT ===============
	src = src[:1][1 : len(src)-1]
	bannerLines := split.SplitBySep(src, "\n")
	inputSlice := strings.Split(input, "\\n")
	result = (print.PrintFinalResult(inputSlice, bannerLines, args))
	return result
}
