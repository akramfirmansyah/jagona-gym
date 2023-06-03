package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/akramfirmansyah/jagona-gym/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ConnectDB() {

	var err error

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect database")
	}
}

func Migrate() {
	_ = DB.AutoMigrate(&models.Trainer{}, &models.TrainerDetail{}, &models.Member{}, &models.Equipment{}, &models.Schedule{}, &models.User{})
	fmt.Println("Success Migrate Database!")
}

func Seeder() {
	var user models.User
	if err := DB.Where("username = ?", "SuperAdmin").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pass := "jagonagym"

			hash, err := hashPassword(pass)
			if err != nil {
				panic(err)
			}

			newUser := models.User{
				Username: "SuperAdmin",
				Password: hash,
			}

			DB.Create(&newUser)
		}
	}
}
