package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rushi444/gin-graphql/graph/model"
)

type dbConfig struct {
	host     string
	port     string
	user     string
	dbname   string
	password string
}

var config = dbConfig{"db", "5432", "user", "app", "pass"}

func getDatabaseURL() string {

	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		config.host, config.port, config.user, config.dbname, config.password)
}

// GetDatabase : Init DB
func GetDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(getDatabaseURL()), &gorm.Config{})
	if err != nil {
		log.Println("Unable to connect to DB", err)
		postgres, _ := db.DB()
		postgres.Close()
	}
	RunMigrations(db)
	return db
}

// RunMigrations : keeps schema in sync with db
func RunMigrations(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.Set("gorm:table_options", "collation_connection=utf8_general_ci")
	db.AutoMigrate(&model.User{}, &model.Todo{})
}
