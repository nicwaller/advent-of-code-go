package intcode

import (
	"advent-of-code/lib/util"
	"advent-of-code/lib/vm"
	"fmt"
	"os"
	"time"
)

type IntcodeVM struct {
	vm     *vm.CounterMachine[int]
	input  <-chan int
	output chan<- int
}

//var ivm = NewVM(nil, nil)

//goland:noinspection GoUnusedExportedFunction
func NewVM(input <-chan int, output chan<- int) *IntcodeVM {
	vma := vm.NewCounterMachine[int]()
	ivm := IntcodeVM{
		vm:     &vma,
		input:  input,
		output: output,
	}

	// ADD
	vma.Op(1, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		tape[argv[2]] = tape[argv[0]] + tape[argv[1]]
	})
	vma.Op(101, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		tape[argv[2]] = argv[0] + tape[argv[1]]
	})
	vma.Op(1001, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		tape[argv[2]] = tape[argv[0]] + argv[1]
	})
	vma.Op(1101, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		tape[argv[2]] = argv[0] + argv[1]
	})

	// MULT
	vma.Op(2, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		tape[argv[2]] = tape[argv[0]] * tape[argv[1]]
	})
	vma.Op(102, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		tape[argv[2]] = argv[0] * tape[argv[1]]
	})
	vma.Op(1002, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		tape[argv[2]] = tape[argv[0]] * argv[1]
	})
	vma.Op(1102, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		tape[argv[2]] = argv[0] * argv[1]
	})

	// INPUT
	vma.Op(3, 1, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		if ivm.input == nil {
			panic("VM requested input, but input channel is not defined")
		}
		timeout := 250 * time.Millisecond
		select {
		case i := <-ivm.input:
			tape[argv[0]] = i
		case <-time.After(timeout):
			fmt.Printf("No input available after waiting %v\n", timeout)
			panic("VM cannot read input")
		}
	})

	// OUTPUT
	vma.Op(4, 1, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		if ivm.output == nil {
			panic("VM requested output, but output channel is not defined")
		}
		outVal := tape[argv[0]]
		select {
		case ivm.output <- outVal:
			// okay
		case <-time.After(100 * time.Millisecond):
			panic("VM output was blocked for at least 10 ms")
		}
	})
	vma.Op(104, 1, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		if output == nil {
			panic("VM requested output, but output channel is not defined")
		}
		outVal := argv[0]
		select {
		case ivm.output <- outVal:
			//fmt.Println(outVal)
		case <-time.After(100 * time.Millisecond):
			panic("VM output was blocked for at least 10 ms")
		}
	})

	// JUMP-IF-TRUE
	vma.Op(5, 2, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		if tape[argv[0]] != 0 {
			cm.Jump(tape[argv[1]])
		}
	})
	vma.Op(105, 2, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		if argv[0] != 0 {
			cm.Jump(tape[argv[1]])
		}
	})
	vma.Op(1005, 2, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		if tape[argv[0]] != 0 {
			cm.Jump(argv[1])
		}
	})
	vma.Op(1105, 2, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		if argv[0] != 0 {
			cm.Jump(argv[1])
		}
	})

	// JUMP-IF-FALSE
	vma.Op(6, 2, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		if tape[argv[0]] == 0 {
			cm.Jump(tape[argv[1]])
		}
	})
	vma.Op(106, 2, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		if argv[0] == 0 {
			cm.Jump(tape[argv[1]])
		}
	})
	vma.Op(1006, 2, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		if tape[argv[0]] == 0 {
			cm.Jump(argv[1])
		}
	})
	vma.Op(1106, 2, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		if argv[0] == 0 {
			cm.Jump(argv[1])
		}
	})

	// LESS-THAN
	vma.Op(7, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		var ret int
		if tape[argv[0]] < tape[argv[1]] {
			ret = 1
		} else {
			ret = 0
		}
		tape[argv[2]] = ret
	})
	vma.Op(107, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		var ret int
		if argv[0] < tape[argv[1]] {
			ret = 1
		} else {
			ret = 0
		}
		tape[argv[2]] = ret
	})
	vma.Op(1007, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		var ret int
		if tape[argv[0]] < argv[1] {
			ret = 1
		} else {
			ret = 0
		}
		tape[argv[2]] = ret
	})
	vma.Op(1107, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		var ret int
		if argv[0] < argv[1] {
			ret = 1
		} else {
			ret = 0
		}
		tape[argv[2]] = ret
	})

	// EQUAL-TO
	vma.Op(8, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		var ret int
		if tape[argv[0]] == tape[argv[1]] {
			ret = 1
		} else {
			ret = 0
		}
		tape[argv[2]] = ret
	})
	vma.Op(108, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		var ret int
		if argv[0] == tape[argv[1]] {
			ret = 1
		} else {
			ret = 0
		}
		tape[argv[2]] = ret
	})
	vma.Op(1008, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		var ret int
		if tape[argv[0]] == argv[1] {
			ret = 1
		} else {
			ret = 0
		}
		tape[argv[2]] = ret
	})
	vma.Op(1108, 3, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		var ret int
		if argv[0] == argv[1] {
			ret = 1
		} else {
			ret = 0
		}
		tape[argv[2]] = ret
	})

	// HALT
	vma.Op(99, 0, func(argv []int, tape []int, cm *vm.CounterMachine[int]) {
		cm.Halt()
	})
	return &ivm
}

//goland:noinspection GoUnusedExportedFunction
func (ivm *IntcodeVM) Exec(program []int) int {
	prog := util.Copy(program)
	ivm.vm.Exec(prog)
	return prog[0]
}

//goland:noinspection GoUnusedExportedFunction
func (ivm *IntcodeVM) ExecAsync(program []int, done chan<- bool) int {
	prog := util.Copy(program)
	go func() {
		ivm.vm.Exec(prog)
		done <- true
	}()
	return prog[0]
}

//goland:noinspection GoUnusedExportedFunction
func ExecIO(program []int, input []int) []int {
	inputChan := make(chan int, len(input))
	for _, v := range input {
		inputChan <- v
	}
	outputChan := make(chan int)
	doneChan := make(chan bool)
	//ivm.output = outputChan
	//ivm.input = inputChan
	ivm := NewVM(inputChan, outputChan)

	prog := util.Copy(program)
	go func() {
		ivm.vm.Exec(prog)
		doneChan <- true
	}()

	var results = make([]int, 0)
CollectOutput:
	for {
		select {
		case x := <-outputChan:
			results = append(results, x)
		case <-doneChan:
			break CollectOutput
		case <-time.After(3000 * time.Millisecond):
			fmt.Println("The VM has been running for 3000 ms, and that's probably bad")
			os.Exit(1)
		}
	}
	return results
}

//goland:noinspection GoUnusedExportedFunction
func Exec(program []int) int {
	prog := util.Copy(program)
	ivm := NewVM(nil, nil)
	ivm.vm.Exec(prog)
	return prog[0]
}

//goland:noinspection GoUnusedExportedFunction
func ExecArgs(program []int, args []int) int {
	prog := util.Copy(program)
	copy(prog[1:], args)
	ivm := NewVM(nil, nil)
	ivm.vm.Exec(prog)
	return prog[0]
}
