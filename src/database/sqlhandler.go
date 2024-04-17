package database

import (
	"fmt"

	"github.com/denismathan/goAuth/src/configurations"
	"github.com/denismathan/goAuth/src/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

var sqlHandler *SqlHandler

func NewSqlHandler(cfg configurations.Database) SqlHandler {
	// dsn := "root:password@tcp(127.0.0.1:3306)/go_sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(builddsn(cfg)), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler = new(SqlHandler)
	sqlHandler.db = db
	db.AutoMigrate(&entities.Todo{}, &entities.User{}, &entities.Token{})
	return *sqlHandler
}

func builddsn(cfg configurations.Database) string {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)
	return dsn
}
