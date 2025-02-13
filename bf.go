package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var cells []byte

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: bf <verb> <filename.bf> OR bf run \"code\"")
		os.Exit(1)
	}
	verb := args[0]
	input := args[1]
	var code []byte
	if info, err := os.Stat(input); err == nil && !info.IsDir() && filepath.Ext(input) == ".b" {
		data, err := os.ReadFile(input)
		if err != nil {
			panic(err)
		}
		code = data
	} else {
		code = []byte(input)
	}
	switch verb {
	case "run":
		run(code)
	case "live":

	default:
		fmt.Println("Invalid verb:", verb)
		os.Exit(1)
	}
}

func filterCode(code []byte) []byte {
	filtered := make([]byte, 0, len(code))
	for _, a := range code {
		switch a {
		case '+', '-', '>', '<', '[', ']', '.', ',':
			filtered = append(filtered, a)
		}
	}
	return filtered
}

func run(code []byte) {
	code = filterCode(code)
	cells = make([]byte, 3000000)
	// ap:	Array Pointer
	//		Position in the array
	// ip:	Instruction Pointer
	//		Current instruction
	ap := 0
	ip := 0
	jt := buildJumpTable(code)
	for ; ip < len(code); ip++ {
		switch code[ip] {
		case '+':
			cells[ap]++
		case '-':
			cells[ap]--
		case '<':
			ap--
			if ap < 0 {
				ap = len(cells) - 1
			}
		case '>':
			ap++
			if ap > len(cells)-1 {
				ap = 0
			}
		case '.':
			fmt.Printf("%c", cells[ap])
		case ',':
			fmt.Scanf("%c", &cells[ap])
		case '[':
			if cells[ap] == 0 {
				ip = jt[ip] - 1
			}
		case ']':
			if cells[ap] != 0 {
				ip = jt[ip] - 1
			}
		}
	}
}

func buildJumpTable(code []byte) map[int]int {
	stack := []int{}
	jt := make(map[int]int)
	for i, c := range code {
		switch c {
		case '[':
			stack = append(stack, i)
		case ']':
			if len(stack) == 0 {
				fmt.Printf("Syntax error: unmatched ']' at position %d\n", i)
				os.Exit(1)
			}
			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			jt[open] = i
			jt[i] = open
		}
	}
	if len(stack) != 0 {
		fmt.Printf("Syntax error: unmatched '[' at position %d\n", stack[0])
		os.Exit(1)
	}
	return jt
}
