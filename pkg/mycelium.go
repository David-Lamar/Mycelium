package pkg

// TODO: Create constants
// TODO: Create function manifest
// TODO: Start profiling?
// TODO: Create return registers or something? Need to handle async function calls and how that will work in the bytecode

type Mycelium struct {
	ID           string
	ConstantPool []Value
	Frames       map[int]Frame

	// TODO: Function manifest
	// TODO: Configuration -- CPU, Memory, etc.
	// TODO: Current running frames, etc.
}

// TODO: When "return" is called, delete the frame from the map
// TODO: When a function calls another function, take the frame ID of the function _calling_ (and maybe the VM ID) so that we know where to put the return values (if any)

func (m *Mycelium) Call(
	id string,
	frame Frame,
) {

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
