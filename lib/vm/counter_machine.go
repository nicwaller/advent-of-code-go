package vm

import (
	"advent-of-code/lib/util"
	"fmt"
)

type CounterMachine[T comparable] struct {
	opcodes map[T]Operation[T]
	halt    bool
	pc      int
}

type Operation[T comparable] struct {
	argc int
	run  func(argv []T, tape []T, vm *CounterMachine[T])
}

//goland:noinspection GoUnusedExportedFunction
func NewCounterMachine[T comparable]() CounterMachine[T] {
	return CounterMachine[T]{
		opcodes: make(map[T]Operation[T]),
	}
}

func (cm *CounterMachine[T]) Op(opcode T, argCount int, run func(argv []T, tape []T, vm *CounterMachine[T])) {
	cm.opcodes[opcode] = Operation[T]{
		argc: argCount,
		run:  run,
	}
}

func (cm *CounterMachine[T]) Halt() {
	cm.halt = true
}

func (cm *CounterMachine[T]) Jump(toPC int) {
	cm.pc = toPC
}

func (cm *CounterMachine[T]) Exec(program []T) {
	exec(cm, program)
}

func exec[T comparable](cm *CounterMachine[T], program []T) {
	// do NOT make a copy
	// the caller can do that, if they want
	// or the tape can be mutated in-place
	// pc = Process Counter
	//pc := 0
	args := make([]T, 3)
	var opcode T
	var ins Operation[T]
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Printf("VM execution failed!\n PC=%d\n Opcode=%v\n Args=%v\n", cm.pc, opcode, args[:ins.argc])
			panic(cm.pc)
		}
	}()
	for cm.halt = false; !cm.halt; {
		var found bool
		opcode = program[cm.pc]
		//fmt.Println(opcode)
		ins, found = cm.opcodes[opcode]
		if !found {
			fmt.Printf("Invalid Opcode: %v\n  Program Counter: %d\n", opcode, cm.pc)
			fmt.Println(program[util.IntMax(0, cm.pc-4):util.IntMin(len(program)-1, cm.pc+4)])
			panic(cm.pc)
		}
		copy(args, program[cm.pc+1:cm.pc+ins.argc+1])
		cm.pc += 1 + ins.argc
		ins.run(args[:ins.argc], program, cm)
	}
}
