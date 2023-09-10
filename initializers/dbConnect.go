package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	var err error;
	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"), os.Getenv("PGPORT"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!=nil{
		panic("Error while connecting to Database")
	}
}

