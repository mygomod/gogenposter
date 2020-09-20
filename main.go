package main

import (
	"github.com/mygomod/gogenposter/pkg/mus"
	"github.com/mygomod/gogenposter/pkg/service"
	"github.com/mygomod/muses"
	"github.com/mygomod/muses/pkg/cmd"
	"github.com/mygomod/muses/pkg/oss"
)

func main() {
	app := muses.Container(
		cmd.Register,
		oss.Register,
	)

	app.SetPostRun(mus.Init, func() error {
		return service.Generate()
	})

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
