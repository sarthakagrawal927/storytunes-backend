package postgres

import (
	"github.com/storyTunes/backend.git/cmd/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB
var err error

func InitPostgresDB() {
	PostgresDB.AutoMigrate(&User{})
}

func ConnectPostgresDatabase() *gorm.DB {
	// postgres://sarthak:4vyfLZJfTCbg2uBR4Clky4ETsEOlFRW0@dpg-cfdu894gqg45rnvbps10-a.singapore-postgres.render.com/main_2azm
	connStr := "host=dpg-cfdu894gqg45rnvbps10-a.singapore-postgres.render.com user=sarthak password=4vyfLZJfTCbg2uBR4Clky4ETsEOlFRW0 dbname=main_2azm port=5432 "
	PostgresDB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	utils.HandleError(err)
	return PostgresDB
}

func ClosePostgresDatabase() {
	dbInstance, err := PostgresDB.DB()
	utils.HandleError(err)
	dbInstance.Close()
}
