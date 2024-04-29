package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func utils() {

}

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Config APP:", viper.Get("app"))
	fmt.Println("Config postgreSQL:", viper.Get("postgreSQL"))
}

var DB *gorm.DB

func InitSQL(dns string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully")
}
