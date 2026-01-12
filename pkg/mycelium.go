package pkg

import (
	"Mycelium/pkg/utils"
	"log/slog"
)

// TODO: Create constants
// TODO: Create function manifest
// TODO: Start profiling?
// TODO: Create return registers or something? Need to handle async function calls and how that will work in the bytecode

type Mycelium struct {
	ID           string
	ConstantPool []Value
	Frames       map[int]*Frame
	Functions    map[int]Function

	FrameCounter int

	Log *slog.Logger

	// TODO: Function manifest
	// TODO: Configuration -- CPU, Memory, etc.
	// TODO: Current running frames, etc.
}

// TODO: When "return" is called, delete the frame from the map
// TODO: When a function calls another function, take the frame ID of the function _calling_ (and maybe the VM ID) so that we know where to put the return values (if any)

func (m *Mycelium) Call(
	functionId int,
	frameId int,
) {
	m.Log.Debug("Executing function", "ID", functionId, "Frame", frameId)
	// TODO: If frameID is 0, treat it as the VM itself is initiating it and wants the return value
	// 	This will be for the "Main" function as well as any functions that don't have return values

	function, ok := m.Functions[functionId]
	if !ok {
		panic("Attempted to call a function that didn't exist")
	}

	if frameId == 0 && (len(function.Inputs) > 0 || len(function.Outputs) > 0) {
		m.Log.Error("Cannot invoke a function with inputs or outputs without a frame", "Frame ID", frameId, "Input Len", len(function.Inputs), "Output Len", len(function.Outputs))
		panic("Cannot invoke a function with inputs or outputs without a frame")
	}

	newFrame := Frame{
		Id:                 m.FrameCounter,
		FunctionID:         functionId,
		Stack:              utils.Stack[Value]{},
		Local:              make(map[int]Value),
		Return:             make(map[int]ReturnRegister),
		InstructionPointer: 0,
		ReturnCounter:      0,
	}

	m.FrameCounter++
	m.Frames[newFrame.Id] = &newFrame

	if frameId == 0 { // Case where it's fire and forget, just start executing it
		m.Log.Debug("Fire and forget function!")
		newFrame.ParentId = 0
	} else { // Case where we need to pop values off, etc.
		fromFrame, ok := m.Frames[frameId]
		if !ok {
			panic("Invalid frame ID provided to call")
		}

		newFrame.ParentId = frameId

		if fromFrame.Stack.Size() < len(function.Inputs) {
			panic("The stack does not have enough values to handle calling the specified function")
		}

		// Populate the frame inputs from the current frame
		// TODO: Could optimize this for local execution by passing pointers
		for i, j := range function.Inputs {
			value := fromFrame.Stack.Pop()
			if value.Type != j {
				panic("Type mismatch on function call")
			}

			newFrame.Local[i] = value
		}

		if len(function.Outputs) > 0 {
			m.Log.Debug("This is a return function; setting the return register")
			newFrame.ParentReturnRegister = fromFrame.ReturnCounter
			fromFrame.ReturnCounter++
			fromFrame.Stack.Push(IntValue(newFrame.ParentReturnRegister))

			fromFrame.Return[newFrame.ParentReturnRegister] = ReturnRegister{
				Available: false,
				Data:      nil,
			}
		}
	}

	// TODO: This should be async (or possibly async)
	Interpret(m, function.Bytecode, &newFrame, m.ConstantPool)

	// TODO:
	// 	1. Look into the function manifest and find the function
	// 	2. Check how many inputs the function has and the types
	// 	3. Pop all of the inputs off the stack; raise an error if something isn't correct
	// 	4. Get the current frame's ID to know where return goes
	// 	5. IF the function has no outputs, skip the return register stuff. It's fire & forget.
	// 	6. Identify an open return register on the frame and reserve it (Now we have frame & register to return to)
	// 	7. IF the function has outputs, Insert the register value with "available" set to false
	// 	8. IF the function has outputs, push the return register index to the stack
	// 	9. Send it to the "router" -- Router will just run locally for now.
	// 	10. Increment program counter, etc. and continue execution until LOAD_RETURN is called
	// 	11. If LOAD_RETURN is called and the register is not populated, the frame goes into a paused state and execution stops.

	// TODO: (Continuation of step 9 above.
	// 	9.1: Router gets the function info (arguments, where it returns, etc.)
	// 	9.2: Router schedules the function execution and starts profiling if enabled
	// 	9.3: Function is executed
	// 	9.4: Function "RETURN" op is called
	// 	9.5: Identify where this function was called from so where know where to send its return values (if there are any -- otherwise skip)
	// 	9.6: Send it to the "router" similar to a function call
	// 	9.7: Original VM gets the return
	// 	9.8: VM identifies the original function frame
	// 	9.9: VM populates the return register on the function frame
	// 	9.10: If the function is in a paused state due to this register's LOAD_RETURN, resume.
	// 	9.11: Profile aggregation is done if profile is underway.

}

// TODO: Might work better if return is on a frame itself. Frames likely should have the Mycelium instance reachable...
// TODO: This might work well too _from_ a frame. So a frame could automatically inject its ID into the call vs. calling "Call" directly on the Mycelium instance...

func (m *Mycelium) Return(
	frameId int,
) {
	fromFrame, ok := m.Frames[frameId]
	if !ok {
		panic("Invalid frame ID provided to call")
	}

	fromFunc, ok := m.Functions[fromFrame.FunctionID]
	if !ok {
		panic("Invalid function ID provided to a frame")
	}

	m.Log.Debug("Returning from function", "Frame", frameId, "Function ID", fromFrame.FunctionID)

	var n int // "Nil" value of an int
	if fromFrame.ParentId != n && len(fromFunc.Outputs) > 0 {
		parentFrame, ok := m.Frames[fromFrame.ParentId]
		if !ok {
			panic("Invalid frame ID as parent ID")
		}

		retReg := []Value{}

		for _, _ = range fromFunc.Outputs {
			// TODO: Validate length of outputs against the stack
			// TODO: Validate the types of the output for the types on the stack

			retReg = append(retReg, fromFrame.Stack.Pop())
		}

		data := parentFrame.Return[fromFrame.ParentReturnRegister]
		data.Data = retReg
		data.Available = true

		parentFrame.Return[fromFrame.ParentReturnRegister] = data
	}

	// Clears up the memory of that frame
	delete(m.Frames, frameId)
}
