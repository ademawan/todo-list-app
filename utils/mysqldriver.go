package utils

import (
	"fmt"
	"time"
	"todo-list-app/configs"
	"todo-list-app/entities"
	"todo-list-app/middlewares"

	"github.com/lithammer/shortuuid"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)
	fmt.Println(connectionString)
	// "root:@tcp(127.0.0.1:3306)/be5db?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitMigrate(db)
	return db
}

func InitMigrate(db *gorm.DB) {
	db.Migrator().DropTable(&entities.Task{})
	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Task{})

	userUid := shortuuid.New()
	password, _ := middlewares.HashPassword("xyz")

	db.Create(&entities.User{
		UserUid:  userUid,
		Name:     "Ade Mawan",
		Email:    "anonimus@gmail.com",
		Password: password,
		Address:  "jl.dramaga no.22",
		Gender:   "male",
	})

	// var userUid []string

	// for i := 0; i < 50; i++ {

	// 	userUid := shortuuid.New()
	// 	password, _ := middlewares.HashPassword("xyz")

	// 	db.Create(&entities.User{
	// 		UserUid:  userUid,
	// 		Name:     faker.Name(),
	// 		Email:    faker.Email(),
	// 		Password: password,
	// 		Address:  "jl.dramaga no.22",
	// 		Gender:   "female",
	// 	})
	// 	taskUid := shortuuid.New()

	// 	db.Create(&entities.Task{
	// 		TaskUid:        taskUid,
	// 		Title:          faker.TitleMale(),
	// 		Priority:       "hight",
	// 		UserUid:        userUid,
	// 		Status:         "waithing",
	// 		Note:           "catatan catatan catatan",
	// 		Todo_date_time: time.Now(),
	// 	})

	// }

	for i := 0; i < 20; i++ {

		// userUid := shortuuid.New()
		// password, _ := middlewares.HashPassword("xyz")

		// db.Create(&entities.User{
		// 	UserUid:  userUid,
		// 	Name:     faker.Name(),
		// 	Email:    faker.Email(),
		// 	Password: password,
		// 	Address:  "jl.dramaga no.22",
		// 	Gender:   "male",
		// })
		taskUid := shortuuid.New()

		layoutFormat := "2006-01-02T15:04"
		todoDateTime, _ := time.Parse(layoutFormat, "2022-03-31T12:26")

		db.Create(&entities.Task{
			TaskUid:        taskUid,
			Title:          "testing title blabla",
			Priority:       "hight",
			UserUid:        userUid,
			Status:         "waithing",
			Note:           "catatan catatan catatan",
			Todo_date_time: todoDateTime,
		})

	}
}
