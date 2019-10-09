package main

import "fmt"

type progressBar struct {
	total, current, len, prev int
}

func (p *progressBar) Update() {
	p.current++
}

func (p *progressBar) Draw() {
	display := ""
	percent := float64(p.current) / float64(p.total)
	complete := int(percent * float64(p.len))
	curr := 0
	for x := 0; x < p.len; x++ {
		if x < complete {
			display += "+"
			curr++
			continue
		}
		display += "-"
	}
	if curr == p.prev && p.current < p.total {
		return
	}

	p.prev = curr
	fmt.Printf("\rProgress:[%s](%.0f%%)", display, percent*100)
	if p.current >= p.total {
		fmt.Printf("\n")
	}
}
