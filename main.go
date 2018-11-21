package main

import (
	"github.com/LucasGao67/MMSM2Wechat/config"
	"github.com/LucasGao67/MMSM2Wechat/router"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "api config file path")
)

func main() {
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Infof(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + viper.GetString("addr") + "/sd/health")
		if err == nil && resp.StatusCode == http.StatusOK {
			return nil
		}

		log.Info("Wating for the router,retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("can not connect to the router")
}
