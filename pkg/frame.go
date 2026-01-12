package pkg

import "Mycelium/pkg/utils"

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
	ParentId             int
	ParentReturnRegister int
	Id                   int
	FunctionID           int
	Stack                utils.Stack[Value]
	Local                map[int]Value
	Return               map[int]ReturnRegister
	ReturnCounter        int
	InstructionPointer   int

	// TODO: A "state" for if it's paused or running
	// TODO: A mutex or _something_ for the Return registers.
	// 	Could get into bad problems where the return register and state are updated on separate threads leading to values not actually being present
	// TODO: Needs an ID
}

type ReturnRegister struct {
	Available bool
	Data      []Value
}
