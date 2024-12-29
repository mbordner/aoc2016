package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common/file"
	"regexp"
	"strconv"
	"strings"
)

var (
	reDigits = regexp.MustCompile(`^(-?\d+)$`)
)

type Registers map[string]int
type Computer struct {
	regs    Registers
	ptr     int
	program []string
}

func (c *Computer) LoadProgram(program []string) {
	c.regs = make(Registers)
	for _, r := range []string{"a", "b", "c", "d"} {
		c.regs[r] = 0
	}
	c.ptr = 0
	c.program = program
}

func (c *Computer) getVal(s string) int {
	if reDigits.MatchString(s) {
		val, _ := strconv.ParseInt(s, 10, 64)
		return int(val)
	}
	return c.getRegVal(s)
}

func (c *Computer) getRegVal(s string) int {
	if v, e := c.regs[s]; e {
		return v
	}
	panic("invalid register " + s)
}

func (c *Computer) setRegVal(r string, v int) {
	if _, e := c.regs[r]; !e {
		panic("invalid register " + r)
	}
	c.regs[r] = v
}

func (c *Computer) Run() Registers {
	for c.ptr < len(c.program) {
		tokens := strings.Fields(c.program[c.ptr])

		c.ptr++

		switch tokens[0] {
		case "cpy":
			c.setRegVal(tokens[2], c.getVal(tokens[1]))
		case "inc":
			c.setRegVal(tokens[1], c.getRegVal(tokens[1])+1)
		case "dec":
			c.setRegVal(tokens[1], c.getRegVal(tokens[1])-1)
		case "jnz":
			val := c.getVal(tokens[1])
			if val != 0 {
				c.ptr += c.getVal(tokens[2]) - 1
			}
		}

	}
	return c.regs
}

func main() {
	program, _ := file.GetLines("../data.txt")
	computer := &Computer{}
	computer.LoadProgram(program)
	computer.regs["c"] = 1
	regs := computer.Run()
	fmt.Println(regs["a"])
}
