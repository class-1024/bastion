package main

import (
	"context"
	_ "bastion/docs"
	"bastion/pkg/setup"
	"bastion/router"
	"bastion/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title bastion server
// @version 1.0
// @description personal projects
// @termsOfService https://github.com/nanzm

// @contact.name nan
// @contact.url https://nancode.cn
// @contact.email msg@nancode.cn

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	setup.Boot()
	utils.InitValidate()

	HttpServerRun()
}

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	gin.SetMode(viper.GetString("base.debug_mode"))

	r := router.Init()

	httpAddr := viper.GetString("http.addr")
	HttpSrvHandler = &http.Server{
		Addr:           httpAddr,
		Handler:        r,
		ReadTimeout:    time.Duration(viper.GetInt("http.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(viper.GetInt("http.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(viper.GetInt("http.max_header_bytes")),
	}
	go func() {
		log.Printf("[INFO] 服务运行在: http://127.0.0.1%s\n", httpAddr)
		if err := HttpSrvHandler.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERROR] 服务运行错误:%s err: %v\n", httpAddr, err)
		}
	}()

	// 使应用挂起 监听信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 关闭
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf("[ERROR] 优雅关闭服务错误:%v\n", err)
		return
	}
	log.Printf("[INFO] 优雅关闭成功 服务停止了 \n")
}
