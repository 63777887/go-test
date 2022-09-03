package main

import (
	_ "jwk/test/pkg/driver"
	"jwk/test/pkg/web"
	"jwk/test/utils"
)

var (
	Log = utils.Logger
)

func main() {
	Log.Debug("TestServer Start...")
	defer func() {
		Log.Debug("TestServer End...")
	}()
	web.WebServer()
}
