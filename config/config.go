package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"myGo/logger"
)

// 服务映射
type ServerConf struct {
	Port string
	Ip   string
}

// mysql配置映射
type MysqlConf struct {
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	Charset  string
	Timeout  string
}

// Redis映射
type RedisConf struct {
	DB           int
	Addr         string
	Port         string
	Password     string
	PoolSize     int
	MaxRetries   int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
}

type IpConf struct {
	IpWhite string
}

type JwtAutoConf struct {
	SingKey string
	Timeout int64
	Header  string
}

var SerConf ServerConf
var MyConf MysqlConf
var RedConf RedisConf
var IConf IpConf
var JwtConf JwtAutoConf

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
		logger.Error("初始化配置文件失败！")
		return
	}
	//服务映射
	SerConf.Port = cfg.Section("server").Key("port").String()
	SerConf.Ip = cfg.Section("server").Key("ip").String()

	//Mysql映射
	MyConf.Dbname = cfg.Section("mysql").Key("dbname").String()
	MyConf.Username = cfg.Section("mysql").Key("username").String()
	MyConf.Password = cfg.Section("mysql").Key("password").String()
	MyConf.Host = cfg.Section("mysql").Key("host").String()
	MyConf.Port = cfg.Section("mysql").Key("port").String()
	MyConf.Charset = cfg.Section("mysql").Key("charset").String()
	MyConf.Timeout = cfg.Section("mysql").Key("timeout").String()

	//redis映射
	RedConf.DB, _ = cfg.Section("redis").Key("db").Int()
	RedConf.Addr = cfg.Section("redis").Key("addr").String()
	RedConf.Port = cfg.Section("redis").Key("port").String()
	RedConf.Password = cfg.Section("redis").Key("password").String()
	RedConf.PoolSize, _ = cfg.Section("redis").Key("pool-size").Int()
	RedConf.MaxRetries, _ = cfg.Section("redis").Key("max-retries").Int()
	RedConf.DialTimeout, _ = cfg.Section("redis").Key("dial-timeout").Int()
	RedConf.ReadTimeout, _ = cfg.Section("redis").Key("read-timeout").Int()
	RedConf.WriteTimeout, _ = cfg.Section("redis").Key("write-timeout").Int()

	//Ip白名单
	IConf.IpWhite = cfg.Section("ip").Key("ip-white").String()

	//Jwt映射
	JwtConf.SingKey = cfg.Section("jwt").Key("sing-key").String()
	JwtConf.Timeout, _ = cfg.Section("jwt").Key("timeout").Int64()
	JwtConf.Header = cfg.Section("jwt").Key("header").String()
	logger.Info("初始化配置文件完成......")
}
