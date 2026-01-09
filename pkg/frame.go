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
	Stack              utils.Stack[Value]
	Local              map[int]Value
	Return             map[int]ReturnRegister
	InstructionPointer int
}

type ReturnRegister struct {
	Available bool
	Data      []Value
}
