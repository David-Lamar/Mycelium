package pkg

type Function struct {
	ID      string
	Profile Profile
	// TODO: For the types, we'll likely want to make sure that custom types (like structs) are accounted for as unique
	Inputs  []Type
	Outputs []Type

	Loaded   bool
	Bytecode [][]byte

	// TODO: The location of the function on disk if not loaded; _or_ a way to find it.

}

type Profile struct {
}

// TODO: Methods for profiling?
// TODO: Methods for loading from disk if not present in the VM right now?

// TODO: LRU Cache for functions
// TODO: Only evict functions who don't have active frames
// TODO: Use profiling data to influence which functions should stay in mem
// TODO: Maybe use call graph to determine which functions _may_ get called by future paths
// TODO: Lazy load as/if needed
