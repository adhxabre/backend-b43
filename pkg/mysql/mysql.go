package mysql

// Import "fmt", "gorm.io/driver/mysql", "gorm.io/gorm" here ...
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Declare DB varible from *gorm.DB here ...
var DB *gorm.DB

// Create DatabaseInit here ...
func DatabaseInit() {
	var err error
	dsn := "root:@tcp(localhost:3306)/dumbmerch?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database!")
}
