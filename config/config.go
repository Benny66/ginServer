package config

import (
	"ginServer/utils/function"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"gopkg.in/ini.v1"
)

var Config *config

func init() {
	Config = NewConfig()
}

func NewConfig() *config {
	c := &config{
		IPAddress:       "0.0.0.0",
		Port:            "8066",
		DBHost:          "127.0.0.1",
		DBPort:          "3306",
		DBUser:          "root",
		DBPass:          "BL@eleven.*#railBroadcast*#",
		DBName:          "audioMatrix",
		WsURL:           "ws://127.0.0.1:6501",
		UdpIpAddr:       function.GetLocalIp(),
		UdpPort:         8806,
		DBType:          "sqlite3",
		DBPath:          "sqlite3/audioMatrix.db",
		Mode:            "release",
		IsUseWs:         true,
		LogExpire:       180, //默认保存6个月
		configPath:      function.GetAbsPath("config.ini"),
		appName:         "Audio Matrix",
		appVersion:      "1.0",
		databaseVersion: "202103101750",
		language:        "zh-cn",
		tokenSecret:     "GUANGZHOU_GD_GENERICBROADCAST_DEPARTMENT_11_HEXL",
		runtimePath:     "runtime/",
		apiLogPath:      "runtime/logs/api/",
		sqlLogPath:      "runtime/logs/sql/",
		systemLogPath:   "runtime/logs/system/",
		wsLogPath:       "runtime/logs/ws/",
		udpLogPath:      "runtime/logs/udp/",
		exportPath:      "public/export/",
		BusinessAddr:    "ws://127.0.0.1:8061",
		LogicAddr:       "ws://127.0.0.1:8062",
	}
	c.loadGlobalConfig()
	return c
}

type config struct {
	IPAddress       string
	Port            string
	DBType          string
	DBPath          string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPass          string
	DBName          string
	WsURL           string
	UdpIpAddr       string
	UdpPort         int
	Mode            string
	IsUseWs         bool
	LogExpire       int //日志保存时间(单位：天)
	configPath      string
	appName         string
	appVersion      string
	databaseVersion string
	language        string
	tokenSecret     string
	runtimePath     string
	apiLogPath      string
	sqlLogPath      string
	wsLogPath       string
	udpLogPath      string
	exportPath      string
	systemLogPath   string
	BusinessAddr    string
	LogicAddr       string
}

/*
* description: 加载config.ini配置
* author: shahao
* created on: 20-3-28 上午11:28
* param param_1:
* param param_2:
* return return_1:
 */
func (c *config) loadGlobalConfig() {
	if !function.IsFileExists(c.configPath) {
		return
	}
	cfg, err := ini.Load(c.configPath)
	if err != nil {
		log.Fatal(err)
	}
	if v := cfg.Section("config").Key("IPAddress").String(); v != "" {
		c.IPAddress = v
	}
	if v := cfg.Section("config").Key("Port").String(); v != "" {
		c.Port = v
	}
	if v := cfg.Section("config").Key("DBHost").String(); v != "" {
		c.DBHost = v
	}
	if v := cfg.Section("config").Key("DBPort").String(); v != "" {
		c.DBPort = v
	}
	if v := cfg.Section("config").Key("DBUser").String(); v != "" {
		c.DBUser = v
	}
	if v := cfg.Section("config").Key("DBPass").String(); v != "" {
		c.DBPass = v
	}
	if v := cfg.Section("config").Key("DBName").String(); v != "" {
		c.DBName = v
	}
	if v := cfg.Section("config").Key("DBType").String(); v != "" {
		c.DBType = v
	}
	if v := cfg.Section("config").Key("DBPath").String(); v != "" {
		c.DBPath = v
	}
	if v := cfg.Section("config").Key("WsAddress").String(); v != "" {
		c.WsURL = "ws://" + v + ":6501"
	}
	if v := cfg.Section("config").Key("Mode").String(); v != "" {
		c.Mode = v
		gin.SetMode(v)
	}
	if v := cfg.Section("config").Key("IsUseWs").String(); v == "false" {
		c.IsUseWs = false
	}
	if v := cfg.Section("config").Key("LogExpire").String(); v != "" {
		day, err := strconv.Atoi(v)
		if err == nil && day > 0 {
			c.LogExpire = day
		}
	}
	if v := cfg.Section("config").Key("UdpIpAddr").String(); v != "" {
		c.UdpIpAddr = v
	}
}
func (c *config) SetUdpIpAddr(ip string) {
	if ip != "" {
		c.UdpIpAddr = ip
	}
	return
}
func (c *config) GetConfigPath() string {
	return c.configPath
}

func (c *config) GetAppName() string {
	return c.appName
}

func (c *config) GetAppVersion() string {
	return c.appVersion
}

func (c *config) GetDatabaseVersion() string {
	return c.databaseVersion
}

func (c *config) GetLanguage() string {
	return c.language
}

func (c *config) GetTokenSecret() string {
	return c.tokenSecret
}

func (c *config) GetRunTimePath() string {
	return c.runtimePath
}

func (c *config) GetApiLogPath() string {
	return c.apiLogPath
}

func (c *config) GetSqlLogPath() string {
	return c.sqlLogPath
}

func (c *config) GetSystemLogPath() string {
	return c.systemLogPath
}

func (c *config) GetDBPath() string {
	return c.DBPath
}

func (c *config) GetWsLogPath() string {
	return c.wsLogPath
}

func (c *config) GetExportPath() string {
	return c.exportPath
}

func (c *config) GetUDPLogPath() string {
	return c.udpLogPath
}
