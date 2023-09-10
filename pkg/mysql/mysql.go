package mysql

import (
	"chat_socket/config"
	"fmt"
	"time"

	"github.com/charlie-bit/utils/db/common"
	"github.com/charlie-bit/utils/db/gmysql"
	"gorm.io/gorm"
)

func NewMysqlGormDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Cfg.Mysql.Username, config.Cfg.Mysql.Password, config.Cfg.Mysql.Address[0], config.Cfg.Mysql.Database,
	)
	client, err := gmysql.NewMysqlClient(
		&common.Config{
			DSN:             dsn,
			Debug:           true,
			MaxIdleConns:    config.Cfg.Mysql.MaxIdleConn,
			MaxOpenConns:    config.Cfg.Mysql.MaxOpenConn,
			ConnMaxLifetime: time.Second * time.Duration(config.Cfg.Mysql.MaxLifeTime),
		},
	)
	return client.DB, err
}
