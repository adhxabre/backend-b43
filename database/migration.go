package database

// Import "dumbmerch/models", "dumbmerch/pkg/mysql", "fmt" here ...
import (
	"dumbmerch/models"
	"dumbmerch/pkg/mysql"
	"fmt"
)

// Create RunMigration function here ...
func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Error!")
	}

	fmt.Println("Migration Succeed!")
}
