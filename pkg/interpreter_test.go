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
	var program = base2Pop("MULT")

	frame := baseTest(program, []Value{intValue(3), intValue(7)})

	if frame.Stack.Pop().Data != 21 {
		t.Fail()
	}
}

func TestEq_false(t *testing.T) {
	var program = base2Pop("EQ")

	frame := baseTest(program, []Value{intValue(3), intValue(7)})

	if frame.Stack.Pop().Data != false {
		t.Fail()
	}
}

func TestEq_true(t *testing.T) {
	var program = base2Pop("EQ")

	frame := baseTest(program, []Value{intValue(3), intValue(3)})

	if frame.Stack.Pop().Data != true {
		t.Fail()
	}
}

func TestLt_true(t *testing.T) {
	var program = base2Pop("LT")

	frame := baseTest(program, []Value{intValue(3), intValue(7)})

	if frame.Stack.Pop().Data != true {
		t.Fail()
	}
}

func TestLt_false(t *testing.T) {
	var program = base2Pop("LT")

	frame := baseTest(program, []Value{intValue(7), intValue(3)})

	if frame.Stack.Pop().Data != false {
		t.Fail()
	}
}

func TestGt_true(t *testing.T) {
	var program = base2Pop("GT")

	frame := baseTest(program, []Value{intValue(7), intValue(3)})

	if frame.Stack.Pop().Data != true {
		t.Fail()
	}
}

func TestGt_false(t *testing.T) {
	var program = base2Pop("GT")

	frame := baseTest(program, []Value{intValue(3), intValue(7)})

	if frame.Stack.Pop().Data != false {
		t.Fail()
	}
}

func TestJump_positive(t *testing.T) {
	var program = `
LOAD_CONST 0
LOAD_CONST 0
JUMP 3
LOAD_CONST 1
LOAD_CONST 1
ADD
`

	frame := baseTest(program, []Value{intValue(3), intValue(7)})

	data := frame.Stack.Pop().Data
	fmt.Printf("%d\n", data)

	if data != 6 {
		t.Fail()
	}
}

func TestJump_negative(t *testing.T) {
	var program = `
JUMP 2
JUMP 4
LOAD_CONST 0
LOAD_CONST 0
JUMP -3
POP
POP
LOAD_CONST 1
LOAD_CONST 1
ADD
`

	frame := baseTest(program, []Value{intValue(3), intValue(7)})

	data := frame.Stack.Pop().Data

	if data != 14 {
		t.Fail()
	}
}

func TestJumpFalse_positive(t *testing.T) {
	var program = `
LOAD_CONST 0
LOAD_CONST 0
LOAD_CONST 2
JUMP_FALSE 3
LOAD_CONST 1
LOAD_CONST 1
ADD
`

	frame := baseTest(program, []Value{intValue(3), intValue(7), boolValue(false)})

	data := frame.Stack.Pop().Data
	fmt.Printf("%d\n", data)

	if data != 6 {
		t.Fail()
	}
}

func TestJumpFalse_negative(t *testing.T) {
	var program = `
JUMP 2
JUMP 5
LOAD_CONST 0
LOAD_CONST 0
LOAD_CONST 2
JUMP_FALSE -4
POP
POP
LOAD_CONST 1
LOAD_CONST 1
ADD
`

	frame := baseTest(program, []Value{intValue(3), intValue(7), boolValue(false)})

	data := frame.Stack.Pop().Data

	if data != 14 {
		t.Fail()
	}
}

func TestJumpFalse_skip(t *testing.T) {
	var program = `
LOAD_CONST 0
LOAD_CONST 0
LOAD_CONST 2
JUMP_FALSE 3
LOAD_CONST 1
LOAD_CONST 1
ADD
`

	frame := baseTest(program, []Value{intValue(3), intValue(7), boolValue(true)})

	data := frame.Stack.Pop().Data
	fmt.Printf("%d\n", data)

	if data != 14 {
		t.Fail()
	}
}

func TestDup(t *testing.T) {
	var program = `
LOAD_CONST 0
DUP
`

	frame := baseTest(program, []Value{intValue(3), intValue(7)})

	if frame.Stack.Size() != 2 {
		t.Fail()
	}

	if frame.Stack.Pop().Data != 3 || frame.Stack.Pop().Data != 3 {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	var program = `
LOAD_CONST 0
POP
`

	frame := baseTest(program, []Value{intValue(3), intValue(7)})

	if frame.Stack.Size() != 0 {
		t.Fail()
	}
}

func TestStoreLocal(t *testing.T) {
	var program = `
LOAD_CONST 0
STORE_LOCAL 0
`

	frame := baseTest(program, []Value{intValue(3)})

	if frame.Local[0].Data != 3 {
		t.Fail()
	}
}

func TestLoadLocal(t *testing.T) {
	var program = `
LOAD_CONST 0
STORE_LOCAL 0
LOAD_LOCAL 0
`

	frame := baseTest(program, []Value{intValue(3)})

	if frame.Stack.Pop().Data != 3 || frame.Stack.Size() != 0 {
		t.Fail()
	}
}

func TestCountTo10(t *testing.T) {
	var program = `
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
JUMP -8
LOAD_LOCAL 0
RETURN`

	frame := baseTest(program, []Value{intValue(0), intValue(10), intValue(1)})

	if frame.Stack.Pop().Data != 10 {
		t.Fail()
	}
}

func TestStoreConst(t *testing.T) {
	var program = `
STORE_CONST 0 0
`

	frame := baseTest(program, []Value{intValue(3)})

	if frame.Local[0].Data != 3 {
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
		Local:              make(map[int]Value),
		InstructionPointer: 0,
	}
}

func intValue(value int) Value {
	return Value{
		Type: TYPE_INT,
		Data: value,
	}
}

func boolValue(value bool) Value {
	return Value{
		Type: TYPE_BOOL,
		Data: value,
	}
}
