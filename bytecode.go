package main

import (
	"errors"
)

type Op byte

const (
	ADD  Op = 0x00
	SUB  Op = 0x08
	DIV  Op = 0x10
	MOD  Op = 0x18
	MULT Op = 0x20

	AND Op = 0x30
	OR  Op = 0x38
	NOT Op = 0x40

	CMP Op = 0x50
	EQ  Op = 0x51
	LT  Op = 0x52
	GT  Op = 0x53

	JUMP         Op = 0x60
	JUMP_FALSE   Op = 0x61
	JUMP_SUCCESS Op = 0x62

	CALL   Op = 0x70
	RETURN Op = 0x78

	MAKE_STRUCT Op = 0xA0
	MAKE_ARRAY  Op = 0xA1

	DUP Op = 0xB0
	POP Op = 0xB1

	LOAD_LOCAL Op = 0xB8
	LOAD_FIELD Op = 0xB9
	LOAD_CONST Op = 0xBA

	STORE_LOCAL Op = 0xC0
	STORE_FIELD Op = 0xC1
	STORE_CONST Op = 0xC2

	MULTI_OP Op = 0xFF
)

func ParseOp(text string) (Op, error) {
	switch text {
	case "ADD":
		return ADD, nil
	case "SUB":
		return SUB, nil
	case "DIV":
		return DIV, nil
	case "MOD":
		return MOD, nil
	case "MULT":
		return MULT, nil
	case "AND":
		return AND, nil
	case "OR":
		return OR, nil
	case "NOT":
		return NOT, nil
	case "CMP":
		return CMP, nil
	case "EQ":
		return EQ, nil
	case "LT":
		return LT, nil
	case "GT":
		return GT, nil
	case "JUMP":
		return JUMP, nil
	case "JUMP_FALSE":
		return JUMP_FALSE, nil
	case "JUMP_SUCCESS":
		return JUMP_SUCCESS, nil
	case "CALL":
		return CALL, nil
	case "RETURN":
		return RETURN, nil
	case "MAKE_STRUCT":
		return MAKE_STRUCT, nil
	case "MAKE_ARRAY":
		return MAKE_ARRAY, nil
	case "DUP":
		return DUP, nil
	case "POP":
		return POP, nil
	case "LOAD_LOCAL":
		return LOAD_LOCAL, nil
	case "LOAD_FIELD":
		return LOAD_FIELD, nil
	case "LOAD_CONST":
		return LOAD_CONST, nil
	case "STORE_LOCAL":
		return STORE_LOCAL, nil
	case "STORE_FIELD":
		return STORE_FIELD, nil
	case "STORE_CONST":
		return STORE_CONST, nil
	case "MULTI_OP":
		return MULTI_OP, nil
	default:
		return MULTI_OP, errors.New("unknown opcode")
	}
}
