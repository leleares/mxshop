package db

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"mxshop/app/pkg/code"
	"mxshop/app/pkg/options"
	v1 "mxshop/app/user/srv/data/v1"
	errors2 "mxshop/pkg/errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbFactory v1.DataFactory
	once      sync.Once
)

type mysqlFactory struct {
	db *gorm.DB
}

func (mf *mysqlFactory) Users() v1.UserStore {
	return NewUsers(mf.db)
}

func (mf *mysqlFactory) Begin() *gorm.DB {
	return mf.db.Begin()
}

var _ v1.DataFactory = (*mysqlFactory)(nil)

func GetDBFactoryOr(mysqlOpts *options.MySQLOptions) (v1.DataFactory, error) {
	if mysqlOpts == nil && dbFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store factory")
	}

	var err error
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlOpts.Username,
			mysqlOpts.Password,
			mysqlOpts.Host,
			mysqlOpts.Port,
			mysqlOpts.Database)

		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.LogLevel(mysqlOpts.LogLevel),
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		)

		db, openErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if openErr != nil {
			err = openErr
			return
		}

		sqlDB, dbErr := db.DB()
		if dbErr != nil {
			err = dbErr
			return
		}

		sqlDB.SetMaxOpenConns(mysqlOpts.MaxOpenConnections)
		sqlDB.SetMaxIdleConns(mysqlOpts.MaxIdleConnections)
		sqlDB.SetConnMaxLifetime(mysqlOpts.MaxConnectionLifetime)

		dbFactory = &mysqlFactory{db: db}
	})

	if dbFactory == nil || err != nil {
		return nil, errors2.WithCode(code.ErrConnectDB, "failed to get mysql store factory")
	}
	return dbFactory, nil
}
