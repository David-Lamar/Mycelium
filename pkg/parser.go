package pkg

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

func Parse(toParse string) [][]byte {
	var bytecode [][]byte

	commands := strings.Split(toParse, "\n")

	for _, line := range commands {
		if len(line) == 0 {
			continue
		}

		bytecode = append(bytecode, parseLine(line))
	}

	println("")
	println("---------- Parsing done! ----------")
	println("")

	return bytecode
}

func parseLine(toParse string) []byte {
	bytes := strings.Split(toParse, " ")
	var ret []byte

	for _, a := range bytes {
		if len(a) == 0 {
			continue
		}

		if a[0] == 'x' {
			// TODO: Handle binary specified numbers
			panic("cannot currently handle binary values")
		}

		value, err := strconv.Atoi(a)

		if err != nil {
			fmt.Printf("Handling op string in bytecode: %s\n", a)
			op, err := ParseOp(a)

			if err != nil {
				panic(err)
			}

			ret = append(ret, byte(op))
			continue
		}

		fmt.Printf("Handling non-binary number in bytecode: %d\n", value)

		bs := make([]byte, 4)
		binary.BigEndian.PutUint32(bs, uint32(value))
		ret = append(ret, bs...)
	}

	return ret
}
