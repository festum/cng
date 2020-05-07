package main

import (
	"github.com/festum/cng"
	"github.com/festum/cng/config"
	"github.com/xormsharp/xorm"
)

func main() {
	var e error

	cfg := config.DefaultConfig()
	db := cng.InitDatabaseWithConfig(*cfg)

	e = db.Sync(cng.Character{})
	if e != nil {
		return
	}

	e = UpdateFix(db.Database().(*xorm.Engine))
	if e != nil {
		panic(e)
	}
	e = CheckLoader(`./cmd/strokefix/dict.json`) //FIXME: use generic path
	if e != nil {
		panic(e)
	}
	e = CheckVerify(db)
	if e != nil {
		panic(e)
	}
}
