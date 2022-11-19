package intcode

import (
	"advent-of-code/lib/util"
	"advent-of-code/lib/vm"
)

var ivm = NewVM()

//goland:noinspection GoUnusedExportedFunction
func NewVM() vm.CounterMachine[int] {
	vma := vm.NewCounterMachine[int]()
	// ADD
	vma.Op(1, 3, func(argv []int, tape []int, halt func()) {
		tape[argv[2]] = tape[argv[0]] + tape[argv[1]]
	})
	// MULT
	vma.Op(2, 3, func(argv []int, tape []int, halt func()) {
		tape[argv[2]] = tape[argv[0]] * tape[argv[1]]
	})
	// HALT
	vma.Op(99, 3, func(argv []int, tape []int, halt func()) {
		halt()
	})
	return vma
}

//goland:noinspection GoUnusedExportedFunction
func Exec(program []int) int {
	prog := util.Copy(program)
	ivm.Exec(prog)
	return prog[0]
}

//goland:noinspection GoUnusedExportedFunction
func ExecArgs(program []int, args []int) int {
	prog := util.Copy(program)
	copy(prog[1:], args)
	ivm.Exec(prog)
	return prog[0]
}
