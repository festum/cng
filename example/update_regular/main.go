package main

import (
	"github.com/festum/cng"
	"github.com/festum/cng/config"
	"github.com/festum/cng/regular"
)

func main() {
	c := config.LoadConfig()
	db := fate.InitDatabaseWithConfig(*c)
	r := regular.New(db)
	r.Run()
}
