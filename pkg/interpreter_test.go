package pkg

import (
	"Mycelium/pkg/utils"
	"fmt"
	"testing"
)

// TODO: Should probably test edge cases like int overflows and stuff
// TODO: Should test the different variants of add, sub, etc. with future data types

func TestLoadConst(t *testing.T) {
	var program = "LOAD_CONST 0"

	frame := baseTest(program, []Value{intValue(77)})

	if frame.Stack.Pop().Data != 77 {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	// Converts to: 1 + 2
	var program = base2Pop("ADD")

	frame := baseTest(program, []Value{intValue(1), intValue(2)})

	if frame.Stack.Pop().Data != 3 {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	// Converts to: 3 - 1
	var program = base2Pop("SUB")

	frame := baseTest(program, []Value{intValue(3), intValue(1)})

	if frame.Stack.Pop().Data != 2 {
		t.Fail()
	}
}

func TestMod(t *testing.T) {
	// Converts to: 100 % 10
	var program = base2Pop("MOD")

	frame := baseTest(program, []Value{intValue(100), intValue(10)})

	if frame.Stack.Pop().Data != 0 {
		t.Fail()
	}
}

func TestMult(t *testing.T) {
	// Converts to: 100 % 10
	var program = base2Pop("MULT")

	frame := baseTest(program, []Value{intValue(3), intValue(7)})

	if frame.Stack.Pop().Data != 21 {
		t.Fail()
	}
}

// ---------------------- Helpers --------------------------

func base2Pop(op string) string {
	return fmt.Sprintf(`
LOAD_CONST 0
LOAD_CONST 1
%s
`, op)
}

func baseTest(
	program string,
	constants []Value,
) Frame {
	bytecode := Parse(program)

	frame := createFrame()

	Interpret(bytecode, &frame, constants)

	return frame
}

func createFrame() Frame {
	return Frame{
		Stack:              utils.Stack[Value]{},
		Local:              make([]Value, 0),
		InstructionPointer: 0,
	}
}

func intValue(value int) Value {
	return Value{
		Type: TYPE_INT,
		Data: value,
	}
}
