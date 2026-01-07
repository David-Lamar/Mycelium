package pkg

import (
	"Mycelium/pkg/utils"
	"encoding/binary"
)

type Type int

const (
	TYPE_INT Type = iota
	TYPE_BYTE
	TYPE_BOOL
	TYPE_FLOAT
	TYPE_STRUCT
	TYPE_OBJECT
)

type Frame struct {
	Stack              utils.Stack[Value]
	Local              []Value
	InstructionPointer int
}

type Value struct {
	Type Type
	Data interface{}
}

func Interpret(
	byteCode [][]byte,
	frame *Frame,
	constants []Value,
) {
	for frame.InstructionPointer < len(byteCode) {
		statement := byteCode[frame.InstructionPointer]

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
			// TODO: AND
			// TODO: OR
			// TODO: NOT
			// TODO: CMP
		case EQ:
			ProcessEq(frame)
			break
		case LT:
			ProcessLt(frame)
			break
		case GT:
			ProcessGt(frame)
			break
			// TODO: Jump
		// TODO: Jump false
		// TODO: Jump success
		// TODO: Call
		// TODO: return
		// TODO: make struct
		// TODO: Make array
		// TODO: dup
		// TODO: pop
		// TODO: load local
		// TODO: load field
		case LOAD_CONST:
			index := binary.BigEndian.Uint32(statement[1:])
			ProcessLoadConst(frame, constants, int(index))
			break
		}

		// TODO: Instructions that jump may need to take this into account or continue to skip this.
		frame.InstructionPointer++
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

func ProcessEq(
	frame *Frame,
) {
	if frame.Stack.Size() < 2 {
		// TODO: Create an error type and send it up
	}

	param2 := frame.Stack.Pop()
	param1 := frame.Stack.Pop()

	if param1.Type != param2.Type {
		frame.Stack.Push(Value{
			Type: TYPE_BOOL,
			Data: false,
		})
		return
	}

	frame.Stack.Push(Value{
		Type: TYPE_BOOL,
		Data: param1.Data == param2.Data,
	})
}

func ProcessGt(
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
				Type: TYPE_BOOL,
				Data: pv1 > pv2,
			})
		}

		// TODO: Return error if type of param2 not supported

	}

	// TODO: Return error if type of param1 not supported

}

func ProcessLt(
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
				Type: TYPE_BOOL,
				Data: pv1 < pv2,
			})
		}

		// TODO: Return error if type of param2 not supported

	}

	// TODO: Return error if type of param1 not supported

}
