package main

import (
	"Mycelium/pkg"
	"log/slog"
	"os"
	"time"
)

var function1 = `
LOAD_CONST 0
LOAD_CONST 1
CALL 1
LOAD_RETURN
`

// Add
var function2 = `
LOAD_LOCAL 0
LOAD_LOCAL 1
ADD
RETURN
`

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	f1 := pkg.Function{
		ID:       0,
		Profile:  pkg.Profile{},
		Inputs:   []pkg.Type{},
		Outputs:  []pkg.Type{},
		Loaded:   true,
		Bytecode: pkg.Parse(function1),
	}

	f2 := pkg.Function{
		ID:       1,
		Profile:  pkg.Profile{},
		Inputs:   []pkg.Type{pkg.TYPE_INT, pkg.TYPE_INT},
		Outputs:  []pkg.Type{pkg.TYPE_INT},
		Loaded:   true,
		Bytecode: pkg.Parse(function2),
	}

	my := pkg.Mycelium{
		ID:           "",
		ConstantPool: []pkg.Value{pkg.IntValue(3), pkg.IntValue(7)},
		Frames:       make(map[int]*pkg.Frame),
		Functions:    make(map[int]pkg.Function),
		FrameCounter: 1,
		Log:          log,
	}

	my.Functions[0] = f1
	my.Functions[1] = f2

	time.Sleep(50 * time.Millisecond)

	// TODO: Need a way to call without returning to an original frame. Both for init and for fire and forget
	my.Call(0, 0)

	// Provides a little space between parsing and running
	// since running happens faster than printing does usually, and prints out of order

	endFrame := my.Frames[1]
	log.Debug("End frame stack value!", "DATA:", endFrame.Stack.Pop().Data)

}
