package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

const PROMPT = ">>> "

const HEADER = `
          ╭────────────────────────────────────────────╮
          │      Give a number(N), and looping         │ 
          │   count that you want to waste time on..   │ 
          ╰────────────────────────────────────────────╯
`

const OPERATION_TEMPLATE = `
  %v try ... 
  ╭─────▶
  │ %v                                        
  ╰─────▶
`

// Start, executes scanner to read user-prompt-inputs
func Start(in io.Reader, out io.Writer) {
	color.New(color.FgMagenta).Add(color.Bold).Println(HEADER)

	scanner := bufio.NewScanner(in)
	for {
		color.New(color.FgWhite).Print(PROMPT)
		if scanned := scanner.Scan(); !scanned {
			return
		}

		n, c := parseInput(scanner.Text())
		if n == 0 && c == 0 {
			return
		}

		// Generate a logger functionality.
		logger := func(operation string, tryIndex int) {
			parsed := fmt.Sprintf(OPERATION_TEMPLATE, tryIndex, operation)
			fmt.Println(parsed)
		}

		executeLoop(n, c, logger)
	}
}

// parseInput parsed user-prompt-input to Number and Loop-count integers.
func parseInput(input string) (n, c int) {
	inputs := strings.Split(input, " ")
	if len(inputs) == 0 {
		return 0, 0
	}

	num, _ := strconv.Atoi(inputs[0])

	count := -1
	if len(inputs) > 1 {
		count, _ = strconv.Atoi(inputs[1])
	}

	return num, count
}
