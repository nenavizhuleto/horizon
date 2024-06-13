package protocol

import "strings"

type Typable interface {
	Type() MessageType
}

// Join implements strings.Join for MessageType
func Join(types ...MessageType) (result MessageType) {
	switch len(types) {
	case 0:
		return
	case 1:
		return types[0]
	}

	var n int
	var maxInt = int(^uint(0) >> 1)

	if len(MessageTypeSeparator) > 0 {
		if len(MessageTypeSeparator) >= maxInt/(len(types)-1) {
			panic("strings: Join output length overflow")
		}
		n += len(MessageTypeSeparator) * (len(types) - 1)
	}
	for _, t := range types {
		if len(t) > maxInt-n {
			panic("strings: Join output length overflow")
		}
		n += len(t)
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(string(types[0]))
	for _, t := range types[1:] {
		b.WriteString(MessageTypeSeparator)
		b.WriteString(string(t))
	}

	return MessageType(b.String())
}
