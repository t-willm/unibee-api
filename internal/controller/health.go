package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"unibee/utility"
)

func HealthCheck(r *ghttp.Request) {
	r.Response.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	r.Response.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT,DELETE,OPTIONS,PATCH")
	r.Response.Header().Add("Access-Control-Allow-Origin", "*")
	r.Response.WriteHeader(http.StatusOK)
	return
}

func Version(r *ghttp.Request) {
	r.Response.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	r.Response.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT,DELETE,OPTIONS,PATCH")
	r.Response.Header().Add("Access-Control-Allow-Origin", "*")
	r.Response.WriteHeader(http.StatusOK)
	r.Response.Write(utility.ReadBuildVersionInfo(r.Context()))
	return
}
