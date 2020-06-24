package config

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"log"
	"time"
)

//配置信息包含：app维度+server维度（服务器通用配置，包括数据源）+db维度+srv维度

//srv需要增加动态性的处理，绑定在公共的服务下，并且进行一段时间的自动检测使用

type App struct {
	RunMode  string
	PageSize int
}

var AppSetting = &App{}

type Server struct {
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type RedisCommon struct {
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisCommonSetting = &RedisCommon{}

var cfg *ini.File

func Init() {
	var err error

	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("app", AppSetting)

	if AppSetting.RunMode == gin.DebugMode {
		cfg, err = ini.Load("config/server/dev.ini")
		if err != nil {
			log.Fatalf("setting.Setup, fail to parse 'conf/dev/dev.ini': %v", err)
		}
	} else {
		cfg, err = ini.Load("config/production/production.ini")
		if err != nil {
			log.Fatalf("setting.Setup, fail to parse 'conf/production/production.ini': %v", err)
		}
	}

	mapTo("server", ServerSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisCommonSetting.IdleTimeout = RedisCommonSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
