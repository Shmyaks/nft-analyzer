package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DB, _ = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=" + os.Getenv("DATABASE_HOST") + " user=" + os.Getenv("DATABASE_USER") + " password=" + os.Getenv("DATABASE_PASSWORD") + " dbname=" + os.Getenv("DATABASE_NAME") + " port=" + os.Getenv("DATABASE_PORT") + " sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	DB.AutoMigrate(&User{}, &Role{})

}

type User struct {
	Id           uint64  `gorm:"primary_key"`
	Email        string  `gorm:"column:email;unique_index"`
	PasswordHash string  `gorm:"column:password;not null"`
	Roles        []*Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	Id    uint64  `gorm:"primary_key"`
	Name  string  `gorm:"column:name"`
	Users []*User `gorm:"many2many:user_roles;"`
}
