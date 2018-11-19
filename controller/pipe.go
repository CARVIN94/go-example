package controller

import (
	"github.com/CARVIN94/go-example/model/pipe"
	"github.com/CARVIN94/go-restful"
)

// Pipe test
func Pipe(ctx *restful.Context) {
	defer ctx.Close()
	data := &pipe.Model{
		State:   1111,
		Message: "is in pipe",
	}
	ctx.Pipe = data
}
