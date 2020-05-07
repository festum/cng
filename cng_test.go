package cng_test

import (
	"context"
	"testing"

	"github.com/festum/chronos"
	"github.com/festum/cng"
	"github.com/festum/cng/config"
)

func TestFate_RunMakeName(t *testing.T) {
	born := chronos.New("2020/02/06 15:45").Solar().Time()
	last := "張"
	cfg := config.DefaultConfig()
	cfg.BaguaFilter = true
	cfg.ZodiacFilter = true
	cfg.SupplyFilter = true
	cfg.HardFilter = true
	cfg.StrokeMin = 3
	cfg.StrokeMax = 24
	cfg.Regular = true
	cfg.RunInit = false
	cfg.FileOutput = config.FileOutput{
		OutputMode: config.OutputModeLog,
		Path:       "name.log",
	}
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
	f := cng.NewFate(last, born, cng.ConfigOption(cfg), cng.SetGender(cng.Female))

	//f.SetDB(eng)
	e := f.MakeName(context.Background())
	if e != nil {
		t.Fatal(e)
	}
}
