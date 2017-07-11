package controllers

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
	"log"
	//"github.com/get-ion/ion/view"
	//"github.com/get-ion/sessions"
)

type Company struct {
	Name  string
	City  string
	Other string
}

var gcompany = &Company{Name: "glow", City: "Suijui", Other: "invented"}

func setCompanyHandler(ctx context.Context) {
	c := &Company{}
	if err := ctx.ReadJSON(c); err != nil {
		ctx.StatusCode(ion.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	log.Printf("Received: %#v\n", c)
	//c.Other = c.Other + " replied"
	gcompany = c
	ctx.JSON(c)
}

func getCompanyHandler(ctx context.Context) {
	cnm := ctx.URLParam("Name")
	if len(cnm) == 0 {
		ctx.StatusCode(ion.StatusBadRequest)
		ctx.WriteString("Name missed")
		return
	}
	log.Printf("asking for company: %#v\n", cnm)
	ctx.JSON(gcompany)
}

func init() {
	addRouteHandler("GET", "company", getCompanyHandler)
	addRouteHandler("POST", "company", setCompanyHandler)
}
