package controller

import (
	"log"

	"github.com/CARVIN94/go-example/model/ascoltatori"

	"github.com/CARVIN94/go-example/model/pipe"
	"github.com/CARVIN94/go-restful"
	"github.com/CARVIN94/go-util/reply"
)

func Midd(ctx *restful.Context) {
	pipe := ctx.Pipe.(*pipe.Model)
	log.Print(ctx.Pipe)
	log.Print(pipe.Message)
	data := &reply.Array{
		Rows:  ascoltatori.FindAll(),
		Total: ascoltatori.CountTotal(),
	}
	ctx.ReplyJSON(data)
}

func IndexHandler(ctx *restful.Context) {
	if !ctx.IsUrlencoded() {
		return
	}
	data := &reply.Object{
		Object: "SUCCESS",
	}
	ctx.ReplyJSON(data)
}
