package main

import (
	"DataCompliance/pkg/setting"
	"DataCompliance/router"
	"fmt"
	"log"
	"net/http"
)

//初始化
func init() {
	setting.Setup()
}

func main() {

	router := router.StartRouter()
	fmt.Printf("%v", setting.ServerSetting)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("[info] start http server listening %d", setting.ServerSetting.HttpPort)
	s.ListenAndServe()
}
