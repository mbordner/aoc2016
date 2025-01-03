package main

import (
	"errors"
	"fmt"
	"github.com/mbordner/aoc2016/common/file"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	reDigits = regexp.MustCompile(`^(-?\d+)$`)
)

const (
	cpy = "cpy"
	tgl = "tgl"
	inc = "inc"
	dec = "dec"
	jnz = "jnz"
)

type Registers map[string]int
type Computer struct {
	regs    Registers
	ptr     int
	program [][]string
	instr   []string
}

func (c *Computer) LoadProgram(program []string) {
	c.regs = make(Registers)
	for _, r := range []string{"a", "b", "c", "d"} {
		c.regs[r] = 0
	}
	c.ptr = 0
	c.program = make([][]string, len(program))
	c.instr = make([]string, len(program))
	for p, stmt := range program {
		c.program[p] = strings.Fields(stmt)
		c.instr[p] = c.program[p][0]
	}
}

func (c *Computer) getVal(s string) (int, error) {
	if reDigits.MatchString(s) {
		val, _ := strconv.ParseInt(s, 10, 64)
		return int(val), nil
	}
	return c.getRegVal(s)
}

func (c *Computer) getRegVal(s string) (int, error) {
	if v, e := c.regs[s]; e {
		return v, nil
	}
	return 0, errors.New("no such register")
}

func (c *Computer) setRegVal(r string, v int) error {
	if _, e := c.regs[r]; !e {
		return errors.New("no such register")
	}
	c.regs[r] = v
	return nil
}

func (c *Computer) toggleInstr(instr string) string {
	switch instr {
	case cpy:
		instr = jnz
	case jnz:
		instr = cpy
	case inc:
		instr = dec
	case dec:
		instr = inc
	case tgl:
		instr = inc
	}
	return instr
}

func (c *Computer) Run() Registers {
	for c.ptr < len(c.program) {
		tokens := c.program[c.ptr]
		instr := c.instr[c.ptr]

		fmt.Printf("%d: %s %s [%v]\n", c.ptr, instr, strings.Join(tokens, " "), c.regs)

		c.ptr++

		switch instr {
		case tgl:
			if ptr, err := c.getRegVal(tokens[1]); err == nil {
				ptr += c.ptr - 1
				if ptr >= 0 && ptr < len(c.program) {
					c.instr[ptr] = c.toggleInstr(c.instr[ptr])
				}
			}
		case cpy:
			if val, err := c.getVal(tokens[1]); err == nil {
				_ = c.setRegVal(tokens[2], val)
			}
		case inc:
			if val, err := c.getRegVal(tokens[1]); err == nil {
				_ = c.setRegVal(tokens[1], val+1)
			}
		case dec:
			if val, err := c.getRegVal(tokens[1]); err == nil {
				_ = c.setRegVal(tokens[1], val-1)
			}
		case jnz:
			if val, err := c.getVal(tokens[1]); err == nil {
				if val != 0 {
					if val, err = c.getVal(tokens[2]); err == nil {
						c.ptr += val - 1
					}
				}
			}
		}
	}
	return c.regs
}

func main() {
	program, _ := file.GetLines("../data.txt")
	computer := &Computer{}
	computer.LoadProgram(program)
	if len(os.Args) > 1 {
		args := os.Args[1:]
		for i, reg := range []string{"a", "b", "c", "d"} {
			if i < len(args) {
				val, _ := strconv.ParseInt(args[i], 10, 64)
				computer.regs[reg] = int(val)
			}
		}
	}
	regs := computer.Run()
	fmt.Println(regs)
}
