package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
	"github.com/get-ion/ion/view"
	//"github.com/get-ion/sessions"
	"./controllers"
)

// $ go get -u github.com/jteeuwen/go-bindata/...
// $ go-bindata ./public/...
// $ go build
// $ ./embedded-single-page-application

var page = struct {
	Title string
}{"Welcome"}

func newApp() *ion.Application {
	app := ion.New()
	app.RegisterView(view.HTML("./dist", ".html").Binary(Asset, AssetNames))

	app.Get("/", func(ctx context.Context) {
		ctx.ViewData("Page", page)
		ctx.View("index.html")
	})

	assetHandler := app.StaticEmbeddedHandler("./dist", Asset, AssetNames)
	app.SPA(assetHandler)

	return app
}

/* TODO: how to turn on html5 history route rewrite */
/* TODO: check token in header */
func main() {
	app := newApp()
	controllers.Setup(app, "/api")
	app.Run(ion.Addr(":8081"))
}

// Note that app.Use/UseGlobal/Done will be executed
// only to the registered routes like our index (app.Get("/", ..)).
// The file server is clean, but you can still add middleware to that by wrapping its "assetHandler".
//
// With this method, unlike StaticWeb("/" , "./public") which is not working by-design anymore,
// all custom http errors and all routes are working fine with a file server that is registered
// to the root path of the server.
