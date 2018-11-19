package controller

import (
	"log"

	"github.com/CARVIN94/go-example/model/ascoltatori"

	"github.com/CARVIN94/go-example/model/pipe"
	"github.com/CARVIN94/go-reply"
	"github.com/CARVIN94/go-restful"
)

// Midd test
func Midd(ctx *restful.Context) {
	defer ctx.Close()
	pipe := ctx.Pipe.(*pipe.Model)
	log.Print(ctx.Pipe)
	log.Print(pipe.Message)
	data := &reply.Array{
		Rows:  ascoltatori.FindAll(),
		Total: ascoltatori.CountTotal(),
	}
	ctx.ReplyJSON(data)
}

// IndexHandler test
func IndexHandler(ctx *restful.Context) {
	defer ctx.Close()
	data := &reply.Object{
		Object: "SUCCESS",
	}
	ctx.ReplyJSON(data)
}
