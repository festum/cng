package cng_test

import (
	"log"
	"testing"

	"github.com/festum/cng"
)

func TestWuGe_WaiGe(t *testing.T) {
	l1, l2, f1, f2 := 1, 1, 1, 1
	for i := 0; i < 80000; i++ {
		if f2 >= cng.WuGeMax {
			f1++
			f2 = 1
		}
		if f1 >= cng.WuGeMax {
			l2++
			f1 = 1
		}
		if l2 >= cng.WuGeMax {
			l1++
			l2 = 1
		}
		wg := cng.CalcWuGe(l1, l2, f1, f2)
		sum := l1 + l2 + f1 + f2
		if wg.ZongGe() != sum {
			log.Println(wg.ZongGe() == sum, l1, l2, f1, f2, wg.ZongGe())
		}
		f2++
	}
}
