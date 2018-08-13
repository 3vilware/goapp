package configuration

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"encoding/json"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql" // _ permite usar funciones del paquete sin pone mysql.Func()
)

type Configuration struct{
	Server string
	Port string 
	User string 
	Password string
	Database string 
}

func getConfiguration() Configuration{
	var c Configuration
	file, err := os.Open("./config.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	err = json.NewDecoder(file).Decode(&c)

	return c
}

func getConection() *gorm.DB{
	c := getConfiguration()
	// user:password@tcp(server:port)/database_name?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", 
						c.User, c.Password, c.Server, c.Port, c.Database)
	db , err := gorm.Open("mysql",dsn) // dsn = data source name

	if err != nil {
		log.Fatal(err)
	}

	return db
}