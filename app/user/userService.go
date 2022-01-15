package user

import (
	"errors"
	"time"

	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

// 注册
func (us UserService) Register(user model.User) (*model.User, error) {
	Db := utils.GetDB()
	// 判断用户名是否存在
	if err := Db.Table(command.DBUser).Where("username = ?", user.Username).Find(&user).Error; err == nil {
		return nil, errors.New("用户名已存在")
	}

	// 判断邮箱是否注册
	if err := Db.Table(command.DBUser).Where("email = ?", user.Email).Find(&user).Error; err == nil {
		return nil, errors.New("邮箱已注册")
	}
	// 对明文进行加密处理
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// 创建用户
	newUser := model.User{
		Id:         utils.NextId(),
		Username:   user.Username,
		Email:      user.Email,
		Password:   string(newPassword),
		CreateTime: model.Time(time.Now()),
		UpdateTime: model.Time(time.Now()),
	}

	// 数据迁移
	//Db.AutoMigrate(&newUser)
	//创建数据
	// if err = Db.Table("t_user").CreateTable(&newUser).Error; err != nil {
	// 	return newUser, errors.New("数据创建失败")
	// }
	if err := Db.Exec("INSERT INTO "+command.DBUser+"(id,username,password,email,create_time,update_time) VALUES(?,?,?,?,?,?)",
		newUser.Id, newUser.Username, newUser.Password, newUser.Email, newUser.CreateTime, newUser.UpdateTime).Error; err != nil {
		return nil, errors.New("用户注册失败")
	}
	return &newUser, nil
}

// 登录
func (us UserService) Login(user model.User) (*model.User, error) {
	Db := utils.GetDB()
	// 判断用户名是否存在
	if err := Db.Table(command.DBUser).Where("username = ?", user.Username).Find(&user).Error; err != nil {
		return nil, errors.New("账户不存在")
	}
	return &user, nil
}
