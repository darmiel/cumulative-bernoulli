package main

import (
	"io"
)

type Direction uint

const (
	DirLeft Direction = iota + 1
	DirRight
)

const ConsoleWidth = 50

func drawLine(w io.Writer, n, bL, bR int64, dir Direction) {
	// calculate per-n length
	nElemLength := float64(ConsoleWidth) / float64(n)
	consoleWidth := int64(nElemLength * float64(n))
	nValue := n / consoleWidth

	left := false
	right := false

	var res = []byte("")
	var i int64
	for i = 0; i < consoleWidth; i++ {
		if i == 0 {
			res = append(res, byte('['))
			continue
		} else if i+1 == consoleWidth {
			res = append(res, byte(']'))
			continue
		}

		// 100
		// 10 | 25 <-> 50
		// |-----------------------------------------------|

		idx := nValue * (i + 1)
		if idx >= bL && !left {
			if dir == DirRight && res[i-1] == byte('-') {
				res = append(res, byte('>'))
			} else {
				res = append(res, byte('{'))
			}
			left = true
			continue
		}
		if idx >= bR && !right {
			if dir == DirLeft && res[i-1] == byte('-') {
				res = append(res, byte('<'))
			} else {
				res = append(res, byte('}'))
			}
			right = true
			continue
		}

		res = append(res, byte('-'))
	}

	println(w, " ├", string(res))
}
