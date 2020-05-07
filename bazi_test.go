package cng_test

import (
	"log"
	"testing"

	"github.com/festum/chronos"
	"github.com/festum/cng"
)

func TestPoint(t *testing.T) {
	t1 := chronos.New("2020/01/24 15:30")
	log.Println(t1.Lunar().EightCharacter())

	bz := cng.NewBazi(t1)
	t.Log(bz.XiYong())
	t.Log(bz.XiYongShen())
}
