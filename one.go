package goConser

import "fmt"

const (
	OneAbove  = 1
	OneBehind = 2
)

type ConstOne struct {
	constName   string
	constValue  interface{}
	annotate    string
	annotatePos int
}

func NewConstOne(name string, value interface{}) ConstOne {
	return ConstOne{constName: name, constValue: value, annotate: "", annotatePos: OneBehind}
}

func (c ConstOne) SetAnnotate(at string) ConstOne {
	c.annotate = at
	return c
}

func (c ConstOne) SetAnnotatePos(atp int) ConstOne {
	c.annotatePos = atp
	return c
}

func (c ConstOne) WriteOne() string {
	var one string
	if c.annotate == "" {
		one = fmt.Sprintf("%s = %v", c.constName, c.constValue)
	} else {
		one = fmt.Sprintf("// %s", c.annotate)
		if c.annotatePos == OneAbove {
			one = fmt.Sprintf("%s\n%s = %v", one, c.constName, c.constValue)
		} else {
			one = fmt.Sprintf("%s = %v %s", c.constName, c.constValue, one)
		}
	}
	one += "\n"
	return one
}
