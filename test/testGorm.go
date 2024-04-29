package main

import (
	"fmt"
	"ginchat/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=127.0.0.1 user=postgres password=7211 dbname=ginchat.db port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "Hanks"
	user.ID = 1
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1)) // find product with integer primary key
	//db.First(user, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(user).Update("PassWord", "5678")
	// Update - update multiple fields
	//db.Model(user).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	//db.Delete(user, 1)
}
