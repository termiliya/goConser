package goConser

import "fmt"

const (
	OneAbove  = 1
	OneBehind = 2
)

type ConstOne struct {
	ConstName   string
	ConstValue  interface{}
	Annotate    string
	AnnotatePos int
}

func NewConstOne(name string, value interface{}) ConstOne {
	return ConstOne{ConstName: name, ConstValue: value, Annotate: "", AnnotatePos: OneBehind}
}

func (c ConstOne) SetAnnotate(at string) ConstOne {
	c.Annotate = at
	return c
}

func (c ConstOne) SetAnnotatePos(atp int) ConstOne {
	c.AnnotatePos = atp
	return c
}

func (c ConstOne) WriteOne() string {
	var one string
	var value interface{}
	switch c.ConstValue.(type) {
	case string:
		value = "\"" + c.ConstValue.(string) + "\""
	default:
		value = c.ConstValue
	}
	if c.Annotate == "" {
		one = fmt.Sprintf("%s = %v", c.ConstName, value)
	} else {
		one = fmt.Sprintf("// %s", c.Annotate)
		if c.AnnotatePos == OneAbove {
			one = fmt.Sprintf("%s\n%s = %v", one, c.ConstName, value)
		} else {
			one = fmt.Sprintf("%s = %v %s", c.ConstName, value, one)
		}
	}
	one += "\n"
	return one
}
