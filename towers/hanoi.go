package towers

import (
	"github.com/golang-collections/collections/stack"
	"log"
	"math"
)

type HanoiTower interface {
	SolveRecursive()
	SolveIterative()
}

func NewHanoiTower(disksNumber int, silent bool) HanoiTower {
	return &hanoiTower{
		silent:      silent,
		disksNumber: disksNumber,
		srcStack:    *stack.New(),
		auxStack:    *stack.New(),
		dstStack:    *stack.New(),
	}
}

type hanoiTower struct {
	silent      bool
	disksNumber int
	srcStack    stack.Stack
	auxStack    stack.Stack
	dstStack    stack.Stack
}

func (h *hanoiTower) SolveRecursive() {
	h.solveRecursive(h.disksNumber, "A", "B", "C")
}

func (h *hanoiTower) solveRecursive(diskNumber int, from, aux, to string) {
	if diskNumber == 1 {
		h.logMove(diskNumber, from, to)
		return
	}

	h.solveRecursive(diskNumber-1, from, to, aux)

	h.logMove(diskNumber, from, to)

	h.solveRecursive(diskNumber-1, aux, from, to)
}

func (h *hanoiTower) SolveIterative() {
	h.solveIterative(h.disksNumber, "A", "B", "C")
}

func (h *hanoiTower) solveIterative(diskNumber int, src, aux, dst string) {
	if diskNumber%2 == 0 {
		temp := dst
		dst = aux
		aux = temp
	}

	numMoves := int(math.Pow(2, float64(diskNumber))) - 1

	for i := h.disksNumber; i >= 1; i-- {
		h.srcStack.Push(i)
	}

	for i := 1; i <= numMoves; i++ {
		if i%3 == 1 {
			h.moveDisksBetweenTwoStacks(h.srcStack, h.dstStack, src, dst)
		} else if i%3 == 2 {
			h.moveDisksBetweenTwoStacks(h.srcStack, h.auxStack, src, aux)
		} else if i%3 == 0 {
			h.moveDisksBetweenTwoStacks(h.auxStack, h.dstStack, aux, dst)
		}
	}
}

func (h *hanoiTower) moveDisksBetweenTwoStacks(src, dst stack.Stack, labelSrc, labelDst string) {
	p1Top := src.Pop()
	p2Top := dst.Pop()

	if p1Top == nil {
		src.Push(p2Top)
		h.logMove(p2Top, labelDst, labelSrc)
		return
	}

	if p2Top == nil {
		dst.Push(p1Top)
		h.logMove(p1Top, labelSrc, labelDst)
		return
	}

	if p1Top.(int) > p2Top.(int) {
		src.Push(p1Top)
		src.Push(p2Top)
		h.logMove(p2Top.(int), labelDst, labelSrc)
		return
	}

	// p1Top.(int) < p2Top.(int)
	dst.Push(p2Top)
	dst.Push(p1Top)
	h.logMove(p1Top.(int), labelSrc, labelDst)
}

func (h *hanoiTower) logMove(diskNumber interface{}, from, to string) {
	if h.silent {
		return
	}

	log.Printf("Move disk %i from %s to %s", diskNumber, from, to)
}
