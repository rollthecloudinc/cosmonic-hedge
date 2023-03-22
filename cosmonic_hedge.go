package main

import (
	"log"
	"github.com/wasmcloud/actor-tinygo"
	"github.com/wasmcloud/interfaces/httpserver/tinygo"
	shapeshift "github.com/rollthecloudinc/cosmonic-hedge/shapeshift"
)

func main() {
	me := CosmonicHedge{}
	actor.RegisterHandlers(httpserver.HttpServerHandler(&me))
}

type CosmonicHedge struct{}

func (e *CosmonicHedge) HandleRequest(ctx *actor.Context, req httpserver.HttpRequest) (*httpserver.HttpResponse, error) {
	actionContext := &shapeshift.ActionContext{}
	log.Print(actionContext.CloudName)
	r := httpserver.HttpResponse{
		StatusCode: 200,
		Header:     make(httpserver.HeaderMap, 0),
		Body:       []byte("hello"),
	}
	return &r, nil
}