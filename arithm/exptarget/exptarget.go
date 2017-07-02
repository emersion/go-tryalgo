// Package exptarget provides a function to build an arithmetic expression to
// reach a target.
package exptarget

import (
	"strconv"
)

// Exp is either an Op or an Int.
type Exp interface{
	String() string
}

// Op is an operation.
type Op struct {
	Type OpType
	Left, Right Exp
}

func (op *Op) String() string {
	return "(" + op.Left.String() + op.Type.String() + op.Right.String() + ")"
}

// OpType is an operation type.
type OpType int

const (
	OpAdd OpType = iota
	OpSub
	OpMul
	OpDiv
)

func (t OpType) String() string {
	switch t {
	case OpAdd:
		return "+"
	case OpSub:
		return "-"
	case OpMul:
		return "*"
	case OpDiv:
		return "/"
	}
	panic("invalid operation")
}

// Int is an integer.
type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

// ExpTarget returns all numbers that can be obtained by adding, substracting,
// multiplying or divising numbers. It cannot handle more than 32 numbers.
func ExpTarget(numbers []int) map[int]Exp {
	if len(numbers) > 32 {
		panic("doesn't support more than 32 numbers")
	}

	cache := make(map[uint32]map[int]Exp)
	for i, n := range numbers {
		cache[uint32(1) << uint(i)] = map[int]Exp{n: Int(n)}
	}

	wholeSet := (uint32(1) << uint(len(numbers))) - 1
	for set := uint32(1); set <= wholeSet; set++ {
		if _, ok := cache[set]; ok {
			continue
		}
		cache[set] = make(map[int]Exp)

		for left := uint32(1); left < set; left++ {
			// We want left in set
			if left & set != left {
				continue
			}
			right := set - left

			for l, lexp := range cache[left] {
				for r, rexp := range cache[right] {
					cache[set][l + r] = &Op{OpAdd, lexp, rexp}
					if n := l - r; n > 0 {
						cache[set][n] = &Op{OpSub, lexp, rexp}
					}
					cache[set][l * r] = &Op{OpMul, lexp, rexp}
					if l % r == 0 {
						cache[set][l / r] = &Op{OpDiv, lexp, rexp}
					}
				}
			}
		}
	}

	return cache[wholeSet]
}
