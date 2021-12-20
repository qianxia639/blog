package user

import (
	"time"

	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

func (us *UserService) Register(user model.User) model.User {
	Db := utils.GetDB()
	// 判断用户名是否存在
	Db.Where("username = ?", user.Username).Find(&user)
	// 对明文进行加密处理
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// 创建用户
	newUser := model.User{
		Id:         utils.NextId(),
		Username:   user.Username,
		Password:   string(newPassword),
		CreateTime: time.Now(),
		EditTime:   time.Now(),
	}

	// 数据迁移
	Db.AutoMigrate(&newUser)
	//创建数据
	Db.Create(&newUser)

	return newUser
}
