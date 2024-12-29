package main

import (
	"encoding/json"
	"fmt"
	"github.com/mbordner/aoc2016/common"
	"github.com/mbordner/aoc2016/common/file"
	"maps"
	"regexp"
	"slices"
	"sort"
	"strings"
)

var (
	reRTG     = regexp.MustCompile(`((\w*)(?:-compatible)*\s+(microchip|generator))`)
	rtgMap    = make(RTGMap)
	nextRTGId = 'a'
)

type StateStrings []string

type RTGMap map[string]string

type ObjectContainer map[string]bool

type Elevator struct {
	Floor   int             `json:"f"`
	Objects ObjectContainer `json:"o"`
}

type State struct {
	Elevator *Elevator         `json:"e"`
	Floors   []ObjectContainer `json:"f"`
}

func (ss StateStrings) String() string {
	sb := strings.Builder{}
	sb.WriteString("[\n")
	for _, s := range ss {
		sb.WriteString(fmt.Sprintf("  %s,\n", s))
	}
	sb.WriteString("]\n")
	return sb.String()
}

func (r *RTGMap) id(s string) string {
	if i, e := (*r)[s]; e {
		return i
	}
	next := string(nextRTGId)
	(*r)[s] = next
	nextRTGId += 1
	return next
}

func (e *Elevator) Clone() *Elevator {
	o := NewElevator(e.Floor)
	o.Objects = e.Objects.Clone()
	return o
}

func (f *ObjectContainer) ToArray() []string {
	a := make([]string, 0, len(*f))
	for o, b := range *f {
		if b {
			a = append(a, o)
		}
	}
	return a
}

func (f *ObjectContainer) Clone() ObjectContainer {
	return maps.Clone(*f)
}

func (f *ObjectContainer) Has(id string) bool {
	if b, e := (*f)[id]; e {
		return b
	}
	return false
}

func (f *ObjectContainer) Count() int {
	count := 0
	for _, b := range *f {
		if b {
			count++
		}
	}
	return count
}

func (f *ObjectContainer) MarshalJSON() ([]byte, error) {
	objs := make([]string, 0, len(*f))
	for code, on := range *f {
		if on {
			objs = append(objs, code)
		}
	}
	sort.Strings(objs)
	return json.Marshal(objs)
}

func (f *ObjectContainer) UnmarshalJSON(b []byte) error {
	*f = make(ObjectContainer)
	var objs []string
	if err := json.Unmarshal(b, &objs); err != nil {
		return err
	}
	for _, code := range objs {
		(*f)[code] = true
	}
	return nil
}

func (s *State) From(str string) *State {
	_ = json.Unmarshal([]byte(str), s)
	return s
}

func (s *State) Clone() *State {
	var o State
	o.Elevator = s.Elevator.Clone()
	o.Floors = make([]ObjectContainer, len(s.Floors))
	for i, f := range s.Floors {
		o.Floors[i] = f.Clone()
	}
	return &o
}

func (s *State) String() string {
	data, _ := json.Marshal(s)
	return string(data)
}

func (s *State) CondensedStateString() string {
	csfs := make([]ObjectContainer, len(s.Floors))
	for i := range s.Floors {
		csfs[i] = s.Floors[i].Clone()
	}
	for o, b := range s.Elevator.Objects {
		if b {
			csfs[s.Elevator.Floor][o] = true
		}
	}
	type CondensedState struct {
		Elevator int               `json:"e"`
		Floors   []ObjectContainer `json:"f"`
	}
	var cs CondensedState
	cs.Elevator = s.Elevator.Floor
	cs.Floors = csfs
	data, _ := json.Marshal(cs)
	return string(data)
}

func (s *State) NextStates() []string {
	var states []string

	cf := s.Elevator.Floor
	var nfs []int

	if s.Elevator.Floor < len(s.Floors)-1 {
		nfs = append(nfs, cf+1)
	}
	if s.Elevator.Floor > 0 {
		nfs = append(nfs, cf-1)
	}

	onCurrentFloor := s.Floors[cf].Clone()
	for o, b := range s.Elevator.Objects {
		if b {
			onCurrentFloor[o] = true
		}
	}

	currentFloorPairs := common.GetPairSets(onCurrentFloor.ToArray())

	for _, nf := range nfs {
		// try moving single items
		for o, b := range onCurrentFloor {
			if b {
				ns := s.Clone()
				ns.Floors[cf] = onCurrentFloor.Clone()
				ns.Elevator = NewElevator(nf)
				ns.Elevator.Objects[o] = true // move to elevator on next floor
				ns.Floors[cf][o] = false      // remove from current floor
				if ns.IsValid([]int{cf, nf}) {
					states = append(states, ns.String())
				}
			}
		}
		// try moving multiple items
		for _, p := range currentFloorPairs {
			ns := s.Clone()
			ns.Floors[cf] = onCurrentFloor.Clone()
			ns.Elevator = NewElevator(nf)
			for _, po := range p {
				ns.Elevator.Objects[po] = true
				ns.Floors[cf][po] = false
			}
			if ns.IsValid([]int{cf, nf}) {
				states = append(states, ns.String())
			}
		}
	}

	return states
}

// IsValid validates floors, checkOnlyFloors can be nil, and it will validate all floors, otherwise it will just
// validate the floors passed in
func (s *State) IsValid(checkOnlyFloors []int) bool {
	if checkOnlyFloors == nil {
		checkOnlyFloors = make([]int, len(s.Floors))
		for i := range s.Floors {
			checkOnlyFloors[i] = i
		}
	}
	if len(s.Elevator.Objects) < 1 || len(s.Elevator.Objects) > 2 {
		return false // this check isn't necessary, but here only to document rules
	}
	for i := range s.Floors {
		if !slices.Contains(checkOnlyFloors, i) {
			continue
		}
		chips := make(ObjectContainer)
		generators := make(ObjectContainer)
		if s.Elevator.Floor == i {
			for o, b := range s.Elevator.Objects {
				if b {
					if o[1] == 'g' {
						generators[o] = true
					} else {
						chips[o] = true
					}
				}
			}
		}
		for o, b := range s.Floors[i] {
			if b {
				if o[1] == 'g' {
					generators[o] = true
				} else {
					chips[o] = true
				}
			}
		}

		for c := range chips {
			powered := generators.Has(string([]byte{c[0], 'g'}))
			if !powered {
				// if this chip is not powered, it means there isn't matching
				// RTG, but if there are any generators at all, it means this chip will
				// fry so this state can't be valid
				if len(generators) > 0 {
					return false
				}
			}
		}
	}

	return true
}

func (s *State) IsGoal() bool {
	if s.Elevator.Floor != len(s.Floors)-1 {
		return false
	}
	for f := len(s.Floors) - 2; f >= 0; f-- {
		if s.Floors[f].Count() > 0 {
			return false
		}
	}
	return true
}

func NewElevator(f int) *Elevator {
	return &Elevator{Floor: f, Objects: make(ObjectContainer)}
}

func main() {
	initialState := getData("../data.txt")

	for _, n := range []string{"elerium", "dilithium"} {
		code := rtgMap.id(n)
		initialState.Floors[0][code+"m"] = true
		initialState.Floors[0][code+"g"] = true
	}

	initialStateStr := initialState.String()

	queue := make(common.Queue[string], 0, 100)
	visited := make(ObjectContainer)

	prev := make(map[string]string)

	queue.Enqueue(initialStateStr)

	var solution StateStrings

	for !queue.Empty() {
		curStr := *(queue.Dequeue())
		cur := (&State{}).From(curStr)
		if cur.IsGoal() {
			solution = StateStrings{curStr}
			for p := prev[curStr]; p != initialStateStr; p = prev[p] {
				solution = append(StateStrings{p}, solution...)
			}
			break
		} else {
			nextStateStrings := cur.NextStates()
			for _, nss := range nextStateStrings {
				ns := (&State{}).From(nss)
				condensedNextStateString := ns.CondensedStateString()
				if !visited.Has(condensedNextStateString) {
					visited[condensedNextStateString] = true
					prev[nss] = curStr
					queue.Enqueue(nss)
				}
			}
		}
	}

	fmt.Println(append(StateStrings{initialStateStr}, solution...))
	fmt.Println(len(solution))
}

func getData(f string) *State {
	lines, _ := file.GetLines(f)
	state := State{Floors: make([]ObjectContainer, len(lines)), Elevator: NewElevator(0)}
	for i, line := range lines {
		state.Floors[i] = make(ObjectContainer)
		matches := reRTG.FindAllStringSubmatch(line, -1)
		for _, m := range matches {
			code := rtgMap.id(m[2])
			if m[3] == "microchip" {
				code += "m"
			} else if m[3] == "generator" {
				code += "g"
			}
			state.Floors[i][code] = true
		}
	}
	return &state
}
