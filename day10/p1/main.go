package main

import (
	"fmt"
	"github.com/mbordner/aoc2016/common/file"
	"regexp"
	"strconv"
)

var (
	reInput = regexp.MustCompile(`value\s+(\d+)\s+goes to (bot|output)\s+(\d+)`)
	reBot   = regexp.MustCompile(`bot\s+(\d+)\s+gives low to\s+(bot|output)\s+(\d+)\s+and high to\s+(bot|output)\s+(\d+)`)
)

// 207 to high
func main() {
	lines, _ := file.GetLines("../data.txt")
	for _, line := range lines {
		if reInput.MatchString(line) {
			matches := reInput.FindStringSubmatch(line)
			input := GetChipReceiverByStrValues("input", matches[1])
			input.Configure(line)
		} else if reBot.MatchString(line) {
			matches := reBot.FindStringSubmatch(line)
			bot := GetChipReceiverByStrValues("bot", matches[1])
			bot.Configure(line)
		}
	}

	chips := []int{17, 61}

	for _, b := range receivers[BOT] {
		bot := b.(*botObj)
		if bot.dispersed != nil {
			if bot.dispersed[0].val == chips[0] && bot.dispersed[1].val == chips[1] {
				fmt.Printf("Bot %d dispersed %d and %d\n", bot.ID(), chips[0], chips[1])
			}
		}
	}
}

type ChipReceiverType int
type ChipReceiverID int
type MicroChip struct {
	val int
}

const (
	BOT ChipReceiverType = iota
	OUTPUT
	INPUT
)

type ChipReceiver interface {
	Configure(string)
	Receive(*MicroChip)
	Type() ChipReceiverType
	ID() ChipReceiverID
}

var (
	receivers = make(map[ChipReceiverType]map[ChipReceiverID]ChipReceiver)
)

func GetChipReceiverByStrValues(st, sid string) ChipReceiver {
	var rt ChipReceiverType
	switch st {
	case "bot":
		rt = BOT
	case "output":
		rt = OUTPUT
	case "input":
		rt = INPUT
	}
	id, _ := strconv.ParseInt(sid, 10, 64)
	return GetChipReceiver(rt, ChipReceiverID(id))
}

func GetChipReceiver(t ChipReceiverType, id ChipReceiverID) ChipReceiver {
	if _, e := receivers[t]; !e {
		receivers[t] = make(map[ChipReceiverID]ChipReceiver)
	}
	if r, e := receivers[t][id]; e {
		return r
	}
	var r ChipReceiver
	switch t {
	case BOT:
		r = ChipReceiver(&botObj{id: id, chips: []*MicroChip{}, to: []ChipReceiver{}})
	case OUTPUT:
		r = ChipReceiver(&chipOutput{id: id, chips: []*MicroChip{}})
	case INPUT:
		r = ChipReceiver(&chipInput{id: id})
	}
	receivers[t][id] = r
	return r
}

type chipInput struct {
	id   ChipReceiverID
	to   ChipReceiver
	sent *MicroChip
}

func (in *chipInput) Receive(chip *MicroChip) {
	if in.to != nil {
		in.sent = chip
		in.to.Receive(chip)
	}
}

func (in *chipInput) Type() ChipReceiverType {
	return INPUT
}

func (in *chipInput) ID() ChipReceiverID {
	return in.id
}

func (in *chipInput) Configure(c string) {
	matches := reInput.FindStringSubmatch(c)
	in.to = GetChipReceiverByStrValues(matches[2], matches[3])

	val, _ := strconv.ParseInt(matches[1], 10, 64)
	mc := &MicroChip{val: int(val)}
	in.Receive(mc)
}

type chipOutput struct {
	id    ChipReceiverID
	chips []*MicroChip
}

func (out *chipOutput) Receive(chip *MicroChip) {
	out.chips = append(out.chips, chip)
}

func (out *chipOutput) Type() ChipReceiverType {
	return OUTPUT
}

func (out *chipOutput) ID() ChipReceiverID {
	return out.id
}

func (out *chipOutput) Configure(c string) {

}

type botObj struct {
	id        ChipReceiverID
	chips     []*MicroChip
	to        []ChipReceiver
	dispersed []*MicroChip
}

func (b *botObj) Receive(chip *MicroChip) {
	b.chips = append(b.chips, chip)
	if len(b.chips) == 2 && len(b.to) == 2 {
		b.disperse()
	}
}

func (b *botObj) disperse() {
	b.dispersed = make([]*MicroChip, 2)
	b.dispersed[0], b.dispersed[1] = b.chips[0], b.chips[1]
	if b.dispersed[0].val > b.dispersed[1].val {
		b.dispersed[0], b.dispersed[1] = b.dispersed[1], b.dispersed[0]
	}
	for i := range b.to {
		b.to[i].Receive(b.dispersed[i])
	}
}

func (b *botObj) Type() ChipReceiverType {
	return BOT
}
func (b *botObj) ID() ChipReceiverID {
	return b.id
}
func (b *botObj) Configure(c string) {
	matches := reBot.FindStringSubmatch(c)
	b.to = append(b.to, GetChipReceiverByStrValues(matches[2], matches[3]))
	b.to = append(b.to, GetChipReceiverByStrValues(matches[4], matches[5]))
	if len(b.chips) == 2 {
		b.disperse()
	}
}
