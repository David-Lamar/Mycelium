package main

import (
	"Mycelium/utils"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

type Type int

const (
	TYPE_INT Type = iota
	TYPE_BYTE
	TYPE_FLOAT
	TYPE_STRUCT
	TYPE_OBJECT
)

type Frame struct {
	Stack utils.Stack[Value]
}

type Value struct {
	Type Type
	Data interface{}
}

var program = `LOAD_CONST 0x00 0x00 0x00 0x00
LOAD_CONST 0x00 0x00 0x00 0x01
ADD
LOAD_CONST 0x00 0x00 0x00 0x01
LOAD_CONST 0x00 0x00 0x00 0x01
SUB
LOAD_CONST 0x00 0x00 0x00 0x01
LOAD_CONST 0x00 0x00 0x00 0x01
MOD
LOAD_CONST 0x00 0x00 0x00 0x01
LOAD_CONST 0x00 0x00 0x00 0x01
MULT`

func main() {
	bytecode := Parse(program)
	frame := Frame{
		Stack: utils.Stack[Value]{},
	}
	constants := []Value{
		{
			Type: TYPE_INT,
			Data: 1,
		},
		{
			Type: TYPE_INT,
			Data: 2,
		},
	}

	Interpret(bytecode, &frame, constants)

	for true {
		if frame.Stack.Size() == 0 {
			break
		}

		fmt.Printf("%d\n", frame.Stack.Pop().Data)
	}
}

func Parse(toParse string) [][]byte {
	var bytecode [][]byte

	commands := strings.Split(toParse, "\n")

	for _, line := range commands {
		bytecode = append(bytecode, ParseLine(line))
	}

	return bytecode
}

func ParseLine(toParse string) []byte {
	bytes := strings.Split(toParse, " ")
	var ret []byte

	for _, a := range bytes {
		b, err := strconv.ParseUint(a, 0, 8)
		if err != nil {
			op, err := ParseOp(a)

			if err != nil {
				panic(err)
			}

			ret = append(ret, byte(op))
			continue
		}
		ret = append(ret, byte(b))
	}

	return ret
}

func Interpret(
	byteCode [][]byte,
	frame *Frame,
	constants []Value,
) {

	for _, statement := range byteCode {
		switch Op(statement[0]) {
		case ADD:
			ProcessAdd(frame)
			break
		case SUB:
			ProcessSub(frame)
			break
		case DIV:
			ProcessDiv(frame)
			break
		case MOD:
			ProcessMod(frame)
			break
		case MULT:
			ProcessMult(frame)
			break
		case LOAD_CONST:
			index := binary.BigEndian.Uint32(statement[1:])
			ProcessLoadConst(frame, constants, int(index))
		}
	}
}

func ProcessLoadConst(
	frame *Frame,
	constants []Value,
	index int,
) {
	// TODO: may need to do error handling if the index is outside of the range of constants
	frame.Stack.Push(constants[index])
}

func ProcessAdd(
	frame *Frame,
) {
	if frame.Stack.Size() < 2 {
		// TODO: Create an error type and send it up
	}

	param2 := frame.Stack.Pop()
	param1 := frame.Stack.Pop()

	if param1.Type == TYPE_INT {

		pv1 := param1.Data.(int)

		if param2.Type == TYPE_INT {
			pv2 := param2.Data.(int)

			frame.Stack.Push(Value{
				Type: TYPE_INT,
				Data: pv1 + pv2,
			})
		}

		// TODO: Return error if type of param2 not supported

	}

	// TODO: Return error if type of param1 not supported
}

func ProcessSub(
	frame *Frame,
) {
	if frame.Stack.Size() < 2 {
		// TODO: Create an error type and send it up
	}

	param2 := frame.Stack.Pop()
	param1 := frame.Stack.Pop()

	if param1.Type == TYPE_INT {

		pv1 := param1.Data.(int)

		if param2.Type == TYPE_INT {
			pv2 := param2.Data.(int)

			frame.Stack.Push(Value{
				Type: TYPE_INT,
				Data: pv1 - pv2,
			})
		}

		// TODO: Return error if type of param2 not supported

	}

	// TODO: Return error if type of param1 not supported
}

func ProcessDiv(
	frame *Frame,
) {
	if frame.Stack.Size() < 2 {
		// TODO: Create an error type and send it up
	}

	param2 := frame.Stack.Pop()
	param1 := frame.Stack.Pop()

	if param1.Type == TYPE_INT {

		pv1 := param1.Data.(int)

		if param2.Type == TYPE_INT {
			pv2 := param2.Data.(int)

			frame.Stack.Push(Value{
				Type: TYPE_INT,
				Data: pv1 / pv2,
			})
		}

		// TODO: Return error if type of param2 not supported

	}

	// TODO: Return error if type of param1 not supported
}

func ProcessMod(
	frame *Frame,
) {
	if frame.Stack.Size() < 2 {
		// TODO: Create an error type and send it up
	}

	param2 := frame.Stack.Pop()
	param1 := frame.Stack.Pop()

	if param1.Type == TYPE_INT {

		pv1 := param1.Data.(int)

		if param2.Type == TYPE_INT {
			pv2 := param2.Data.(int)

			frame.Stack.Push(Value{
				Type: TYPE_INT,
				Data: pv1 % pv2,
			})
		}

		// TODO: Return error if type of param2 not supported

	}

	// TODO: Return error if type of param1 not supported
}

func ProcessMult(
	frame *Frame,
) {
	if frame.Stack.Size() < 2 {
		// TODO: Create an error type and send it up
	}

	param2 := frame.Stack.Pop()
	param1 := frame.Stack.Pop()

	if param1.Type == TYPE_INT {

		pv1 := param1.Data.(int)

		if param2.Type == TYPE_INT {
			pv2 := param2.Data.(int)

			frame.Stack.Push(Value{
				Type: TYPE_INT,
				Data: pv1 * pv2,
			})
		}

		// TODO: Return error if type of param2 not supported

	}

	// TODO: Return error if type of param1 not supported
}
