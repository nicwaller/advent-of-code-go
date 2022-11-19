package vm

type CounterMachine[T comparable] struct {
	tape    []T
	opcodes map[T]Operation[T]
}

type Operation[T comparable] struct {
	argc int
	run  func(argv []T, tape []T, halt func())
}

//goland:noinspection GoUnusedExportedFunction
func NewCounterMachine[T comparable]() CounterMachine[T] {
	return CounterMachine[T]{
		opcodes: make(map[T]Operation[T]),
	}
}

func (cm *CounterMachine[T]) Op(opcode T, argCount int, run func(registers []T, tape []T, halt func())) {
	cm.opcodes[opcode] = Operation[T]{
		argc: argCount,
		run:  run,
	}
}

func (cm *CounterMachine[T]) Exec(program []T) {
	// do NOT make a copy
	// the caller can do that, if they want
	// or the tape can be mutated in-place

	// pc = Process Counter
	pc := 0
	args := make([]T, 3)
	halt := false
	haltFn := func() { halt = true }
	for !halt {
		ins, found := cm.opcodes[program[pc]]
		if !found {
			panic(pc)
		}
		copy(args, program[pc+1:pc+ins.argc+1]) // TODO: verify correctness
		ins.run(args, program, haltFn)
		pc += 1 + ins.argc
	}
}

//goland:noinspection GoUnusedFunction
func exec(progOrig []int) int {
	prog := make([]int, len(progOrig))
	copy(prog, progOrig)

	ptr := 0
	for {
		opcode := prog[ptr]
		if opcode == 99 {
			break
		}
		op1 := prog[prog[ptr+1]]
		op2 := prog[prog[ptr+2]]
		register := prog[ptr+3]
		switch opcode {
		case 1: // add
			prog[register] = op1 + op2
		case 2: // mult
			prog[register] = op1 * op2
		default:
			panic(opcode)
		}
		ptr += 4
	}
	return prog[0]
}
