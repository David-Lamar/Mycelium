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
	TYPE_ERROR
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
		case AND:
			ProcessAnd(frame)
			break
		case OR:
			ProcessOr(frame)
			break
		case NOT:
			ProcessNot(frame)
			break
		case CMP:
			ProcessCmp(frame)
		case EQ:
			ProcessEq(frame)
			break
		case LT:
			ProcessLt(frame)
			break
		case GT:
			ProcessGt(frame)
			break
		case JUMP:
			offset := binary.BigEndian.Uint32(statement[1:])
			// TODO: There's gotta be a better way to make a signed int from bytes...
			ProcessJump(frame, int(int32(offset)))
			break
		case JUMP_FALSE:
			offset := binary.BigEndian.Uint32(statement[1:])
			// TODO: There's gotta be a better way to make a signed int from bytes...
			ProcessJumpFalse(frame, int(int32(offset)))
			break
		case JUMP_SUCCESS:
			offset := binary.BigEndian.Uint32(statement[1:])
			// TODO: There's gotta be a better way to make a signed int from bytes...
			ProcessJumpSuccess(frame, int(int32(offset)))
			break
		// TODO: Call
		// TODO: return
		// TODO: make struct
		// TODO: Make array
		case DUP:
			ProcessDup(frame)
			break
		case POP:
			ProcessPop(frame)
			break
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

func ProcessAnd(frame *Frame) {

}

func ProcessOr(frame *Frame) {

}

func ProcessNot(frame *Frame) {

}

func ProcessCmp(frame *Frame) {

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

func ProcessJump(frame *Frame, offset int) {
	frame.InstructionPointer += offset - 1
}

func ProcessJumpFalse(frame *Frame, offset int) {
	popped := frame.Stack.Pop()
	if popped.Type == TYPE_BOOL && popped.Data == false {
		frame.InstructionPointer += offset - 1
	}
}

func ProcessJumpSuccess(frame *Frame, offset int) {
	popped := frame.Stack.Pop()
	if popped.Type != TYPE_ERROR {
		frame.InstructionPointer += offset - 1
	}
}

func ProcessDup(frame *Frame) {
	if frame.Stack.Size() < 1 {
		// TODO: Create an error type and send it up
	}

	frame.Stack.Push(frame.Stack.Top())
}

func ProcessPop(frame *Frame) {
	//if frame.Stack.Size() < 1 {
	//    // TODO: Create an error type and send it up
	//}

	frame.Stack.Pop()
}
