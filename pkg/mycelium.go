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
	arguments []Value,
	returnFrame int,
) {

}
