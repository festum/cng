package cng

import "github.com/festum/chronos"

type NaYin struct {
	calendar *chronos.Calendar
}

func NewNaYin(calendar *chronos.Calendar) *NaYin {
	return &NaYin{
		calendar: calendar,
	}
}
