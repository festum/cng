package main

import (
	"context"

	"github.com/festum/chronos"
	"github.com/festum/cng"
	"github.com/festum/cng/config"
)

func main() {
	born := chronos.New("2020/07/01 18:00")
	lastName := "秦"
	cfg := config.DefaultConfig()
	cfg.TrigramFilterEnabled = true
	cfg.ZodiacFilterEnabled = true
	cfg.SupplyFilterEnabled = true
	cfg.HardFilter = true
	cfg.StrokeMin = 3
	cfg.StrokeMax = 24
	cfg.Database = config.Database{
		Host:         "localhost",
		Port:         "3306",
		User:         "root",
		Pwd:          "111111",
		Name:         "fate",
		MaxIdleCon:   0,
		MaxOpenCon:   0,
		Driver:       "mysql",
		File:         "",
		Dsn:          "",
		ShowSQL:      false,
		ShowExecTime: false,
	}
	cfg.FileOutput = config.FileOutput{
		OutputMode: config.OutputModeLog,
		Path:       "name.log",
	}

	f := cng.NewFate(lastName, born.Solar().Time(), cng.ConfigOption(cfg))

	e := f.MakeName(context.Background())
	if e != nil {
		panic(e)
	}
}
