package main

import (
	"ginchat/router"
	"ginchat/utils"
	"log"

	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig()
	dns, ok := viper.Get("postgreSQL.DNS").(string)
	if !ok {
		log.Fatalf("Invalid type for database DSN")
	}
	utils.InitSQL(dns)

	r := router.Router()
	r.Run(":8882") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
