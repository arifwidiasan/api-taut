package database

import (
	"fmt"
	"time"

	"github.com/arifwidiasan/api-taut/config"
	"github.com/arifwidiasan/api-taut/model"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf config.Config) *gorm.DB {

	conectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
		conf.LOC,
	)
	DB, err := gorm.Open(mysql.Open(conectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("error open conection : ", err)
	}

	admin := DB.Migrator().HasTable(&model.Admin{})
	if !admin {
		pass := []byte(conf.ADMIN_PASSWORD)
		hash, _ := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
		DB.Migrator().CreateTable(&model.Admin{})
		DB.Model(&model.Admin{}).Create([]map[string]interface{}{
			{"username": "admin", "password": string(hash), "created_at": time.Now()},
		})
	}

	DB.AutoMigrate(&model.Admin{}, &model.User{}, &model.Sosmed{}, &model.Section{})
	return DB
}
