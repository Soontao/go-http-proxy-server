package app

import (
	"net/http"

	"github.com/elazarl/goproxy"
)

// CreateApp but not run
func CreateApp(param *WebAppParam) *WebApplication {
	app := &WebApplication{
		param: param,
	}
	return app
}

type WebApplication struct {
	param *WebAppParam
}

func (app *WebApplication) Run(addr string) error {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	return http.ListenAndServe(addr, proxy)
}
