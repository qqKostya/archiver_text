package vlc

import "strings"

type DecodingTree struct {
	Value string
	Zero  *DecodingTree // Left
	One   *DecodingTree // Right
}

func (et encodingTable) DecodingTree() DecodingTree {
	res := DecodingTree{}

	for ch, code := range et {
		res.Add(code, ch)
	}

	return res
}

func (dt *DecodingTree) Decode(str string) string {
	var buf strings.Builder

	curNode := dt

	for _, ch := range str {
		if curNode.Value != "" {
			buf.WriteString(curNode.Value)
			curNode = dt
		}

		switch ch {
		case '0':
			curNode = curNode.Zero
		case '1':
			curNode = curNode.One
		}
	}

	if curNode.Value != "" {
		buf.WriteString(curNode.Value)
		curNode = dt
	}

	return buf.String()
}

func (dt *DecodingTree) Add(code string, value rune) {
	currNode := dt

	for _, ch := range code {
		switch ch {
		case '0':
			if currNode.Zero == nil {
				currNode.Zero = &DecodingTree{}
			}
			currNode = currNode.Zero
		case '1':
			if currNode.One == nil {
				currNode.One = &DecodingTree{}
			}
			currNode = currNode.One
		}
	}

	currNode.Value = string(value)
}
