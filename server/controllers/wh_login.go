package controllers

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
	"log"
)

type UserAuth struct {
	Id       string
	Password string
}

var AllUsers = []UserAuth{{"admin", "admin"}}

func checkAccount(au *UserAuth) bool {
	log.Println("check account  ", au)
	for _, u := range AllUsers {
		if u == *au {
			return true
		}
	}
	return false
}

func loginPostHandler(ctx context.Context) {
	u := &UserAuth{}
	if err := ctx.ReadJSON(u); err != nil {
		ReplyError(ctx, ion.StatusBadRequest, err.Error())
		return
	}
	if !checkAccount(u) {
		ReplyError(ctx, ion.StatusUnauthorized, "authentication fail")
		return
	}
	sess := sessmgr.Start(ctx)
	sess.Set("userid", u.Id)
	ctx.JSON(map[string]string{"userid": u.Id})
}

func logoutPostHandler(ctx context.Context) {
	sess := sessmgr.Start(ctx)
	sess.Clear()
	ReplyOK(ctx)
}

func chpassPostHandler(ctx context.Context) {

}

func init() {

	//sessmgr = sessions.New(sessions.Config{Cookie: "myappsessionid"})

	addRouteHandler("POST", "login", loginPostHandler)
	addRouteHandler("POST", "logout", logoutPostHandler)
	addRouteHandler("POST", "chpass", chpassPostHandler)
}
