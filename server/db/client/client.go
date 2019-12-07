package client

import (
	config "final-project/server/db/config"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var once sync.Once
var db *gorm.DB

func GetConnectionDB() *gorm.DB {
	once.Do(func() {
		conf := &config.DatabaseConfig{}
		err := config.LoadDatabaseConfig()
		checkError(err)
		if err := viper.Unmarshal(conf); err != nil {
			log.Fatal("Load config: ", err)
		}
		conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True",
		conf.Username, conf.Password, conf.Hostname + ":" + strconv.Itoa(conf.Port), conf.DBName)
		gormDb, err := gorm.Open("mysql", conn)
		if err != nil {
			log.Fatal("Connect to db: ", err)
		}
		gormDb.SingularTable(true)
		gormDb.DB().SetConnMaxLifetime(-1)
		gormDb.DB().SetMaxOpenConns(-1)
		db = gormDb
	})
	return db
}

func checkError(err error) {
	fmt.Println(viper.ConfigFileUsed())
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
