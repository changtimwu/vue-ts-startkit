package controllers

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func ReplyOK(ctx context.Context) {
	ctx.StatusCode(ion.StatusOK)
	ctx.JSON(map[string]string{"result": "OK"})
}

func ReplyError(ctx context.Context, stcode int, msg string) {
	//ctx.StatusCode( stcode)
	code := ion.StatusOK
	if stcode == ion.StatusBadRequest {
		code = stcode
	}
	ctx.StatusCode(code) //always return OK to avoid browser red
	ctx.JSON(map[string]string{"error": msg})
}
