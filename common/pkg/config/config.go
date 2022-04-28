package config

import (
	"github.com/869413421/micro-service/common/pkg/types"
	"os"
	"sync"
	"time"
)

var once sync.Once
var config *Configuration

type Configuration struct {
	Db *Db `json:"db"`
}

type Db struct {
	Address               string        `json:"address"`
	Database              string        `json:"database"`
	User                  string        `json:"user"`
	Password              string        `json:"password"`
	Charset               string        `json:"charset"`
	MaxConnections        int           `json:"max_connections"`
	MaxIdeConnections     int           `json:"max_ide_connections"`
	ConnectionMaxLifeTime time.Duration `json:"connection_max_life_time"`
}

// LoadConfig 加载配置文件
func LoadConfig() *Configuration {
	//1.适用sync.one，使配置只加载一次，后续不需要读取直接返回
	once.Do(func() {
		//1.1从环境变量中读取配置信息
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		database := os.Getenv("DB_DATABASE")
		password := os.Getenv("DB_PASSWORD")
		dbMaxConnections, _ := types.StringToInt(os.Getenv("DB_MAX_CONNECTIONS"))
		dbMaxIdeConnections, _ := types.StringToInt(os.Getenv("DB_MAX_IDE_CONNECTIONS"))
		dbConnectionMaxLifeTime, _ := types.StringToInt(os.Getenv("DB_CONNECTIONS_MAX_LIFE_TIME"))

		//1.2初始化配置结构体
		dbConfig := &Db{
			Address:               host,
			Database:              database,
			User:                  user,
			Password:              password,
			Charset:               "utf8",
			MaxConnections:        dbMaxConnections,
			MaxIdeConnections:     dbMaxIdeConnections,
			ConnectionMaxLifeTime: time.Duration(dbConnectionMaxLifeTime) * time.Minute,
		}

		config = &Configuration{Db: dbConfig}
	})

	return config
}
