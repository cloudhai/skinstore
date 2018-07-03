package main

import (
	"net/http"
	"time"
	"skinstore/web/router"
	"skinstore/web"
	"skinstore/common/logger"
	"skinstore/mqtt"
	"fmt"
	mqtt2 "github.com/eclipse/paho.mqtt.golang"
)

var Mc *mqtt.MqttClient

func main(){

	r := &router.Router{}
	r.RegHandlers(web.InitRoute())
	svr :=http.Server{
		Addr:":8082",
		ReadTimeout:time.Second*5,
		WriteTimeout:time.Second*5,
		Handler:r,
	}
	mc := mqtt.NewMqttClient("admin","cloudhai","test")
	mc.Subscribe(func(client mqtt2.Client, message mqtt2.Message) {
		fmt.Println("get msg form mqtt:"+string(message.Payload()))
	})
	logger.NewLog().Info("start server ...")
	svr.ListenAndServe()
}

