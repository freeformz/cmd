package main

import (
	"fmt"

	"github.com/freeformz/tcmd"
	"github.com/heroku/slog"
)

func main() {
	tcmd.Test()
	fmt.Println(slog.Context{})
}
