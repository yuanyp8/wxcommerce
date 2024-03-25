package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/yuanyp8/wxcommerce/shared/constants"
)

/**
 * @Author: yuanyp8
 * @Author: geniusgavin@163.com
 * @Date: 2024/3/21 23:20
 * @Desc:
 */

/*
配置读取顺序如下：
	1. 首先从 ENV中提取Consul的连接信息,如果获取不到连接信息, 从如下地址进行获取
	2. 其次从Flags中获取ServerConfig发filepath, 如果不提供filepath, 从本地默认地址读取
最后读取的重要性越高, 命令行参数如果捕获到则可以覆盖前面的配置
*/

var (
	validate     = validator.New()
	serverConfig *ServerConfig
)

type ServerConfig struct {
	Name        string         `mapstructure:"name" json:"name"`
	Host        string         `mapstructure:"host" json:"host"`
	PasetoInfo  *PasetoConfig  `mapstructure:"paseto" json:"paseto"`
	MysqlInfo   *MysqlConfig   `mapstructure:"mysql" json:"mysql"`
	OtelInfo    *OtelConfig    `mapstructure:"otel" json:"otel"`
	WXInfo      *WXConfig      `mapstructure:"wx_config" json:"wx_config"`
	BlobSrvInfo *BlobSrvConfig `mapstructure:"blob_srv" json:"blob_srv"`
}

func newServerConfig() *ServerConfig {
	return &ServerConfig{
		Name:        "pmsmodel",
		Host:        "",
		PasetoInfo:  newPasetoConfig(),
		MysqlInfo:   newMysqlConfig(),
		OtelInfo:    newOtelConfig(),
		WXInfo:      newWXConfig(),
		BlobSrvInfo: newBlobSrvConfig(),
	}
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Database string `mapstructure:"db" json:"db" validate:"required" env:"MYSQL_DATABASE"`
	User     string `mapstructure:"user" json:"user" env:"MYSQL_USER"`
	Password string `mapstructure:"password" json:"password" validate:"required" env:"MYSQL_PASSWORD"`
	Salt     string `mapstructure:"salt" json:"salt"` // 盐值，用于加密等操作
}

func newMysqlConfig() *MysqlConfig {
	return &MysqlConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "wx_commerce",
		User:     "root",
	}
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

func newOtelConfig() *OtelConfig {
	return &OtelConfig{
		EndPoint: "http://127.0.0.1:16686",
	}
}

type WXConfig struct {
	AppId     string `mapstructure:"app_id" json:"app_id"`
	AppSecret string `mapstructure:"app_secret" json:"app_secret"`
}

func newWXConfig() *WXConfig {
	return &WXConfig{
		AppId:     "wx",
		AppSecret: "wx",
	}
}

type PasetoConfig struct {
	SecretKey string `mapstructure:"secret_key" json:"secret_key"`
	Implicit  string `mapstructure:"implicit" json:"implicit"`
}

func newPasetoConfig() *PasetoConfig {
	return &PasetoConfig{
		SecretKey: "pmsmodel",
		Implicit:  "pmsmodel",
	}
}

// BlobSrvConfig OSS存储服务
type BlobSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

func newBlobSrvConfig() *BlobSrvConfig {
	return &BlobSrvConfig{
		Name: "oss",
	}
}

func loadServerConfig(filepath ...string) error {
	var (
		path   string
		config = newServerConfig()
	)

	if len(filepath) == 0 || filepath[0] == "" {
		path = constants.PMSConfigPath
	} else {
		path = filepath[0]
	}

	vip := viper.New()
	vip.SetConfigFile(path)

	if err := vip.ReadInConfig(); err != nil {
		return fmt.Errorf("fatal error loading config file: %w", err)
	}

	if err := vip.Unmarshal(config); err != nil {
		return fmt.Errorf("fatal error unmarshal config: %w", err)
	}

	if err := env.Parse(config); err != nil {
		return fmt.Errorf("fatal error parsing env config: %w", err)
	}
	if err := validate.Struct(config); err != nil {
		return err
	}
	return nil
}

func C() *ServerConfig {
	return serverConfig
}
