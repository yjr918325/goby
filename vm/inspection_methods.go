package vm

import (
	"bytes"
	"fmt"
	"strings"
)

func (i *instruction) inspect() string {
	var params []string

	for _, param := range i.Params {
		params = append(params, fmt.Sprint(param))
	}
	return fmt.Sprintf("%s: %s", i.action.name, strings.Join(params, ", "))
}

func (is *instructionSet) inspect() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("<%s>\n", is.label.name))
	for _, i := range is.instructions {
		out.WriteString(i.inspect())
		out.WriteString("\n")
	}

	return out.String()
}

func (cf *callFrame) inspect() string {
	if cf.ep != nil {
		return fmt.Sprintf("Name: %s. is block: %t. ep: %d", cf.instructionSet.label.name, cf.isBlock, len(cf.ep.locals))
	}
	return fmt.Sprintf("Name: %s. is block: %t", cf.instructionSet.label.name, cf.isBlock)
}

func (cfs *callFrameStack) inspect() string {
	var out bytes.Buffer

	for _, cf := range cfs.callFrames {
		if cf != nil {
			out.WriteString(fmt.Sprintln(cf.inspect()))
		}
	}

	return out.String()
}

func (s *stack) inspect() string {
	var out bytes.Buffer
	datas := []string{}

	for i, p := range s.Data {
		if p != nil {
			o := p.Target
			if i == s.VM.sp {
				datas = append(datas, fmt.Sprintf("%s (%T) %d <----", o.Inspect(), o, i))
			} else {
				datas = append(datas, fmt.Sprintf("%s (%T) %d", o.Inspect(), o, i))
			}

		} else {
			if i == s.VM.sp {
				datas = append(datas, "nil <----")
			} else {
				datas = append(datas, "nil")
			}

		}

	}

	out.WriteString("-----------\n")
	out.WriteString(strings.Join(datas, "\n"))
	out.WriteString("\n---------\n")

	return out.String()
}
