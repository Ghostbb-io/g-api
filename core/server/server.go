package server

import (
	"fmt"
	"github.com/Ghostbb-io/g-api/core/router"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
	"github.com/Ghostbb-io/g-api/pkg/global"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Run() {
	router.Init()
	port := fmt.Sprintf(":%d", global.GB_CONFIG.System.Port)
	server := &http.Server{
		Addr:           port,
		Handler:        ginx.GetEngine(),
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 保證文本順序輸出
	time.Sleep(10 * time.Microsecond)
	global.GB_LOG.Info("Ghostbb server run success on ", zap.String("address", port))
	https := "https"
	if !global.GB_CONFIG.System.Https {
		https = "http"
	}
	fmt.Printf(`
	HRMS api 
	Default swagger url：%s://127.0.0.1%s/api/swagger/index.html
	`, https, port)

	fmt.Println("Server startup completed!!!")
	if global.GB_CONFIG.System.Https {
		global.GB_LOG.Error(server.ListenAndServeTLS("./SSL/certificate.crt", "./SSL/private.key").Error())
	} else {
		global.GB_LOG.Error(server.ListenAndServe().Error())
	}
}
