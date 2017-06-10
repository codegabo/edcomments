package configuration

import (
	"encoding/json"
	"log"
	"os"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

func getConfiguration() configuration {
	var c configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// esta linea toma la información del archivo y la mapea a la estructura
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// función que devuelve la conexión a la BD
// GetConnection obtiene una conexión a la bd
func GetConnection() *gorm.DB {
	c := getConfiguration()
	// dsn: Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
