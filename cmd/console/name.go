package main

import (
	"context"
	"fmt"
	"time"

	"github.com/festum/chronos"
	"github.com/festum/cng"
	"github.com/festum/cng/config"
	"github.com/goextension/log"
	"github.com/spf13/cobra"
)

func cmdName() *cobra.Command {
	path := ""
	born := ""
	last := ""
	gender := 0
	cmd := &cobra.Command{
		Use:   "name",
		Short: "output the name",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("start", time.Now().String())
			config.DefaultJSONPath = path
			cfg := config.LoadConfig()
			fmt.Printf("config loaded: %+v", cfg)
			bornTime, e := time.Parse(chronos.DateFormat, born)
			if e != nil {
				log.Fatalw("parseborn", "error", e)
			}
			f := cng.NewFate(last, bornTime, cng.ConfigOption(cfg), cng.SetGender(cng.Gender(gender)))

			e = f.MakeName(context.Background())
			if e != nil {
				log.Fatalw("makename", "error", e)
			}

			fmt.Println("end", time.Now().String())
		},
	}
	cmd.Flags().StringVarP(&last, "last", "l", "", "set lastname")
	cmd.Flags().StringVarP(&born, "born", "b", time.Now().Format(chronos.DateFormat), "set birth as 2016/01/02 15:04")
	cmd.Flags().StringVarP(&path, "path", "p", ".", "set the input path")
	cmd.Flags().BoolVarP(&gender, "gender", "s", false, "set baby gender")
	return cmd
}
