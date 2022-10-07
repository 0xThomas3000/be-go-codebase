package seed

import (
	"log"

	"github.com/0xThomas3000/be-go-codebase/api/models"
	"gorm.io/gorm"
)

// We will be seeding the users and posts table when we eventually wire the database.
var users = []models.User{
	{
		Username: "thomas",
		Email:    "tuanit168@gmail.com",
		Password: "12345678",
	},
	{
		Username: "baolinh",
		Email:    "ronnielinh@gmail.com",
		Password: "ronnielinh",
	},
}

var posts = []models.Post{
	{
		Title:   "Title 1",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum",
	},
	{
		Title:   "Title 2",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum",
	},
}

func Load(db *gorm.DB) {
	if err := db.Debug().Migrator().DropTable(&models.Post{}, &models.User{}, &models.Like{}, &models.Comment{}); err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	if err := db.Debug().AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	if err := db.Debug().Migrator().CreateConstraint(&models.User{}, "Post"); err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range users {
		if err := db.Debug().Model(&models.User{}).Create(&users[i]).Error; err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		if err := db.Debug().Model(&models.Post{}).Create(&posts[i]).Error; err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
