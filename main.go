package main

import (
	"net/http"
	"time"
	"skinstore/common"
	"skinstore/web/router"
	"skinstore/web"
)

func main(){
	common.NewLog().Info("start server ...")
	r := &router.Router{}
	r.RegHandlers(web.InitRoute())
	svr :=http.Server{
		Addr:":8082",
		ReadTimeout:time.Second*5,
		WriteTimeout:time.Second*5,
		Handler:r,
	}
	svr.ListenAndServe()
}

