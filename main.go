package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName, Password string
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:Siyam.95@tcp(127.0.0.1:3306)/blog_site_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db.AutoMigrate(&User{}) database table create and update

	//db.Create(&User{UserName: "Siyam", Password: "123456"} ) //insert data

	// var user User
	//db.First(&user, 1) // find user with id 1
	// db.First(&user, "password = ?", "123456") // find user with user_name Siyam
	// fmt.Println(user.UserName)

	// var users []User
	// db.Find(&users)
	// for _, user := range users {
	// 	fmt.Println(user.UserName)
	// }

	// var user User
	// db.First(&user, 1)
	// db.Model(&user).Update("UserName", "Gandalf")

	// var user User
	// db.First(&user, 2)
	// db.Model(&user).Updates(User{UserName: "pyton", Password: "pip"})

	db.Delete(&User{}, 1)

}
