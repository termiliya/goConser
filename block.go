package goConser

import "fmt"

const (
	BlockAbove = 1
)

type ConstBlock struct {
	manyConst    []ConstOne
	blockNote    string
	blockNotePos int
}

func NewConstBlock() *ConstBlock {
	return &ConstBlock{blockNotePos: BlockAbove}
}

func (cb *ConstBlock) SetBlockNote(note string) *ConstBlock {
	cb.blockNote = note
	return cb
}

func (cb *ConstBlock) SetBlockNotePos(notePos int) *ConstBlock {
	cb.blockNotePos = notePos
	return cb
}

func (cb *ConstBlock) AddConstOne(one ConstOne) *ConstBlock {
	cb.manyConst = append(cb.manyConst, one)
	return cb
}

func (cb *ConstBlock) WriteBlock() string {
	var block string
	block += "\n"
	if cb.blockNote != "" && cb.blockNotePos == BlockAbove {
		block = fmt.Sprintf("// %s\n", cb.blockNote)
	}
	block += "const (\n"
	for _, one := range cb.manyConst {
		block += one.WriteOne()
	}
	block += ")\n"

	return block
}
