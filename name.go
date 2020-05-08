package cng

import (
	"strconv"
	"strings"

	"github.com/festum/chronos"
	"github.com/festum/yi"
)

type Name struct {
	FirstName   []*Character //名姓
	LastName    []*Character
	born        *chronos.Calendar
	eightChars  *BaZi
	trigram     *yi.Yi //周易八卦
	zodiac      *Zodiac
	zodiacPoint int
}

func (n Name) String() string {
	var s string
	for _, l := range n.LastName {
		s += l.Ch
	}
	for _, f := range n.FirstName {
		s += f.Ch
	}
	return s
}

func (n Name) Strokes() string {
	var s []string
	for _, l := range n.LastName {
		s = append(s, strconv.Itoa(l.ScienceStroke))
	}

	for _, f := range n.FirstName {
		s = append(s, strconv.Itoa(f.ScienceStroke))
	}
	return strings.Join(s, ",")
}

func (n Name) PinYin() string {
	var s string
	for _, l := range n.LastName {
		s += "[" + strings.Join(l.PinYin, ",") + "]"
	}

	for _, f := range n.FirstName {
		s += "[" + strings.Join(f.PinYin, ",") + "]"
	}
	return s
}

func (n Name) FiveElements() string {
	var s string
	for _, l := range n.LastName {
		s += l.FiveElements
	}
	for _, f := range n.FirstName {
		s += f.FiveElements
	}
	return s
}

func (n Name) XiYongShen() string {
	return n.eightChars.XiYongShen()
}

func createName(impl *fateImpl, f1 *Character, f2 *Character) *Name {
	lastSize := len(impl.lastChar)
	last := make([]*Character, lastSize, lastSize)
	copy(last, impl.lastChar)
	ff1 := *f1
	ff2 := *f2
	first := []*Character{&ff1, &ff2}

	return &Name{
		FirstName: first,
		LastName:  last,
	}
}

func (n *Name) Trigram() *yi.Yi {
	if n.trigram == nil {
		lastSize := len(n.LastName)
		shang := getStroke(n.LastName[0])
		if lastSize > 1 {
			shang += getStroke(n.LastName[1])
		}
		xia := getStroke(n.FirstName[0]) + getStroke(n.FirstName[1])
		n.trigram = yi.NumberQiGua(xia, shang, shang+xia)
	}

	return n.trigram
}

func (n Name) BaZi() string {
	return n.eightChars.String()
}
