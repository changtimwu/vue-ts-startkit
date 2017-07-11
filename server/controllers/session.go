package controllers

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
	"github.com/get-ion/sessions"
	"log"
	"strings"
)

var (
	cookieNameForSessionID = "intrising-session"
	sessmgr                = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

func alwaysAllowed(s string) bool {
	prefixes := [...]string{"/static/"}
	if s == "/index.html" || s == "/api/login" {
		return true
	}
	for _, pre := range prefixes {
		if strings.HasPrefix(s, pre) {
			return true
		}
	}
	return false
}

func sessionMiddleware(ctx context.Context) {
	s := sessmgr.Start(ctx)
	req := ctx.Request()
	userid := ""
	hasAccess := false
	if alwaysAllowed(req.URL.Path) {
		hasAccess = true
	} else {
		userid = s.GetString("userid")
		log.Println("userid=", userid)
		if userid != "" {
			hasAccess = true
		}
	}

	log.Printf("%v: %v by %v hasAccess=%v\n", req.Method, req.URL.Path, userid, hasAccess)
	if !hasAccess {
		/* API access without token should returns error */
		if strings.HasPrefix(req.URL.Path, "/api") {
			ReplyError(ctx, ion.StatusBadRequest, "No token found! access denied")
		} else {
			/* page access without token should redirect */
			ctx.Redirect("/login")
		}
		return
	}
	ctx.Next()
}
