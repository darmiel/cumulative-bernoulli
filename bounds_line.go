package main

import (
	"log"
)

type Direction uint

const (
	DirLeft Direction = iota + 1
	DirRight
)

const ConsoleWidth = 50

func drawLine(n, bL, bR int64, dir Direction) {
	// calculate per-n length
	nElemLength := float64(ConsoleWidth) / float64(n)
	consoleWidth := int64(nElemLength * float64(n))
	nValue := n / consoleWidth

	left := false
	right := false

	var res []byte
	var i int64
	for i = 0; i < consoleWidth; i++ {
		if i == 0 {
			res = append(res, []byte("[")[0])
			continue
		} else if i+1 == consoleWidth {
			res = append(res, []byte("]")[0])
			continue
		}

		// 100
		// 10 | 25 <-> 50
		// |-----------------------------------------------|

		idx := nValue * (i + 1)
		if idx >= bL && !left {
			res = append(res, []byte("{")[0])
			left = true
			continue
		}
		if idx >= bR && !right {
			res = append(res, []byte("}")[0])
			right = true
			continue
		}

		res = append(res, []byte("-")[0])
	}

	log.Println(" ├", string(res))
}
