package controllers

import (
	"fmt"
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

type CtrlRoute struct {
	method, path string
	handler      func(ctx context.Context)
}

var gRoutes = map[string]CtrlRoute{}

func addRouteHandler(method string, path string, hfunc func(ctx context.Context)) {
	gRoutes[method+" "+path] = CtrlRoute{method, path, hfunc}
}

func Setup(app *ion.Application, mountpath string) {
	app.Use(sessionMiddleware)
	for _, route := range gRoutes {
		var fpath = fmt.Sprintf("%s/%s", mountpath, route.path)
		app.Handle(route.method, fpath, route.handler)
		fmt.Printf("register route %s %s\n", route.method, fpath)
	}
}
