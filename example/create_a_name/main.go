package main

import (
	"context"

	"github.com/festum/chronos"

	"github.com/festum/cng"
	"github.com/festum/cng/config"
)

func main() {

	//cfg := config.DefaultConfig() 參數如下
	//config.Config{
	//	HardFilter: false,
	//	//輸出最大筆畫數
	//	StrokeMax: 3,
	//	//輸出最小筆畫數
	//	StrokeMin: 18,
	//	//立春修正（出生日期為立春當日時間為已過立春八字需修正）
	//	FixBazi: true,
	//	//三才五格過濾
	//	SupplyFilter: true,
	//	//生肖過濾
	//	ZodiacFilter: true,
	//	//周易八卦過濾
	//	BaguaFilter: true,
	//	//連線DB：
	//	Database: config.Database{
	//		Host:         "localhost",
	//		Port:         "3306",
	//		User:         "root",
	//		Pwd:          "111111",
	//		Name:         "fate",
	//		MaxIdleCon:   0,
	//		MaxOpenCon:   0,
	//		Driver:       "mysql",
	//		File:         "",
	//		Dsn:          "",
	//		ShowSQL:      false,
	//		ShowExecTime: false,
	//	},
	//})
	//出生日期
	born := chronos.New("2020/01/14 02:45")
	//姓氏
	lastName := "張"
	cfg := config.DefaultConfig()
	cfg.BaguaFilter = true
	cfg.ZodiacFilter = true
	cfg.SupplyFilter = true
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

	f := fate.NewFate(lastName, born.Solar().Time(), fate.ConfigOption(cfg))

	e := f.MakeName(context.Background())
	if e != nil {
		panic(e)
	}
}
