package main

import (
	"github.com/plagiari-sm/api-gin-couchbase-boilerplate/app"
	psmconfig "github.com/plagiari-sm/psm-config"
)

func init() {
	psmconfig.NewConfig()
}
func main() {
	app := new(app.APP)
	app.Initialize()
	app.Serve()
}
