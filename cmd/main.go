package main

import (
	"Mycelium/pkg"
	"Mycelium/pkg/utils"
	"fmt"
	"time"
)

var program = `
LOAD_CONST 0
LOAD_CONST 1
ADD

LOAD_CONST 1
LOAD_CONST 1
SUB

LOAD_CONST 1
LOAD_CONST 1
MOD

LOAD_CONST 1
LOAD_CONST 1
MULT`

// Goal!
var countTo10 = `
LOAD_CONST 0
STORE_LOCAL 0

LOAD_LOCAL 0
LOAD_CONST 1
LT
JUMP_FALSE 6
LOAD_LOCAL 0
LOAD_CONST 2
ADD
STORE_LOCAL 0
JUMP -6
RETURN`

func main() {
	bytecode := pkg.Parse(program)

	frame := pkg.Frame{
		Stack:              utils.Stack[pkg.Value]{},
		Local:              make(map[int]pkg.Value),
		InstructionPointer: 0,
	}
	constants := []pkg.Value{
		{
			Type: pkg.TYPE_INT,
			Data: 1,
		},
		{
			Type: pkg.TYPE_INT,
			Data: 2,
		},
	}

	// Provides a little space between parsing and running
	// since running happens faster than printing does usually, and prints out of order
	time.Sleep(50 * time.Millisecond)

	pkg.Interpret(bytecode, &frame, constants)

	for {
		if frame.Stack.Size() == 0 {
			break
		}

		fmt.Printf("%d\n", frame.Stack.Pop().Data)
	}
}
