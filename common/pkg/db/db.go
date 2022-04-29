package db

import (
	"fmt"
	"github.com/869413421/micro-service/common/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type BaseModel struct {
	ID        uint64    "gorm:column:id;primaryKey;autoIncrement;not null"
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}

//GetStringID 主键转字符串
func (model BaseModel) GetStringID() string {
	return strconv.Itoa(int(model.ID))
}

// CreatedAtDate 获取模型创建时间
func (model BaseModel) CreatedAtDate() string {
	return model.CreatedAt.Format("2006-01-02 15:04:05")
}

// UpdatedAtDate 获取模型更新时间
func (model BaseModel) UpdatedAtDate() string {
	return model.UpdatedAt.Format("2006-01-02 15:04:05")
}

var gormDb *gorm.DB
var dbConfig *config.Db

// connectDB 链接数据库
func connectDB() (*gorm.DB, error) {
	// 1.获取配置
	serviceConfig := config.LoadConfig()
	dbConfig = serviceConfig.Db

	//2.链接数据库
	gormDb, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Address, dbConfig.Database, dbConfig.Charset,
	)), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	//3.返回数据库链接
	return gormDb, nil
}

func setupDB() {
	//1.获取链接
	conn, err := connectDB()
	if err != nil {
		panic(err)
	}
	conn.Set("gorm:table_options", "ENGINE=InnoDB")
	conn.Set("gorm:table_options", "Charset=utf8")
	sqlDB, err := conn.DB()
	if err != nil {
		panic(fmt.Sprintf("connection to db error %v", err))
	}

	//2.设置最大连接数
	sqlDB.SetMaxOpenConns(dbConfig.MaxConnections)

	//3.设置最大空闲连接数
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdeConnections)

	//4. 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(dbConfig.ConnectionMaxLifeTime * time.Minute)

	//5.设置好连接池，重新赋值
	gormDb = conn
}

// GetDB 开放给外部获得db连接
func GetDB() *gorm.DB {
	//1.如果db为空，初始化链接池
	if gormDb == nil {
		setupDB()
	}

	//2.返回db对象给外部使用
	return gormDb
}
