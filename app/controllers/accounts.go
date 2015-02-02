package controllers

import (
    "github.com/revel/revel"
    "errors"
    "strings"
    "encoding/base64"
    "net/http"
)

type Accounts struct {
	*revel.Controller
}

func (c Accounts) CheckAuth() revel.Result {
    res := checkAuth(c.Controller)
    if res != nil {
        return res
    }
    
    return c.RenderJson("OK")
}

func checkAuth(c *revel.Controller) revel.Result {
    username, password, ok := basicAuth(c.Request.Request)
    if !ok {
        return renderJsonError(c, errors.New("Missing auth info"))
    }
    
    usernameConf, _ := revel.Config.String("auth.username")
	passwordConf, _ := revel.Config.String("auth.password")
	
	if username != usernameConf || password != passwordConf {
	    return renderJsonError(c, errors.New("Incorrect username or password"))
	}
	
	return nil
}

func basicAuth(r *http.Request) (username, password string, ok bool) {
    auth := r.Header.Get("Authorization")
    if auth == "" {
        return
    }
    return parseBasicAuth(auth)
}
func parseBasicAuth(auth string) (username, password string, ok bool) {
    if !strings.HasPrefix(auth, "Basic ") {
        return
    }
    c, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
    if err != nil {
        return
    }
    cs := string(c)
    s := strings.IndexByte(cs, ':')
    if s < 0 {
        return
    }
    return cs[:s], cs[s+1:], true
}