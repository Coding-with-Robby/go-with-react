package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB")), &gorm.Config{})

	if err != nil {
		fmt.Println("Can't connect to db")
	}

}
