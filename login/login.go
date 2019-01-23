package main

import (
	_ "expvar"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"

	_ "github.com/lib/pq"

	"qianuuu.com/lib/client"
	"qianuuu.com/lib/logs"
	"qianuuu.com/kuaigou/login/api"
	"qianuuu.com/kuaigou/login/config"
	"qianuuu.com/kuaigou/login/internal/rest"
)

func startServer() {
	logs.Info("Starting server %s ...", config.Opts.Port)
	r := rest.New()
	api.SetupRoutes(r)
	go r.Run(":" + config.Opts.Port)
}

func handleSignals() {
	// 添加性能分析工具
	pfport := os.Getenv("GO_PPROF")
	if pfport != "" {
		go func() {
			logs.Info("[main] start pprof port %s ...", pfport)
			err := http.ListenAndServe(":"+pfport, nil)
			if err != nil {
				logs.Error("[main] pprof " + err.Error())
			}
		}()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

// Start 入口
func main() {
	// 设置时区
	time.Local = time.FixedZone("CST", 8*3600)

	// 读取配置文件
	if err := config.ParseToml("config.toml"); err != nil {
		logs.Info("配置文件读取失败: %s", err.Error())
		return
	}

	if len(config.Opts.LogPath) > 0 {
		logs.Info("log path: %s", config.Opts.LogPath)
		log, err := logs.New("debug", config.Opts.LogPath)
		if err == nil {
			logs.Export(log)
			defer func() {
				log.Close()
			}()
		} else {
			logs.Error(err.Error())
		}
	}

	client, err := client.NewClient("postgres", config.Opts.ConnString)
	logs.Info("connect db: %s", config.Opts.ConnString)
	if err != nil {
		logs.Error(err.Error())
	}
	defer func() {
		_ = client.Close()
	}()
	startServer()
	handleSignals()
}
