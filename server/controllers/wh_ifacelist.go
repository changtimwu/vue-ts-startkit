package controllers

import (
	"github.com/get-ion/ion/context"
)

func iflistHandler(ctx context.Context) {

}

func init() {
	addRouteHandler("GET", "ifacelist", iflistHandler)
}
