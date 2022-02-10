package app

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
)

type UserService struct {
}

// 注册
func (us UserService) Register(user model.User) (*model.User, error) {
	// Db := utils.GetDB()
	// 判断用户名是否存在
	// if err := global.RY_DB.Debug().Table(command.DBUser).Where("username = ?", user.Username).Find(&user).Error; err == nil {
	// 	return nil, errors.New("用户名已存在")
	// }
	var u model.User
	global.RY_DB.Debug().Select("username").Where("username = ?", user.Username).Find(&u)

	if u.Username == user.Username {
		return nil, errors.New("用户名已存在")
	}

	// if err := global.RY_DB.Debug().Select("username").Where("username = ?", user.Username).Find(&model.User{}).Error; err == nil {
	// 	return nil, errors.New("用户名已存在")
	// }

	// 判断邮箱是否注册
	// if err := global.RY_DB.Debug().Table(command.DBUser).Where("email = ?", user.Email).Find(&user).Error; err == nil {
	// 	return nil, errors.New("邮箱已注册")
	// }
	global.RY_DB.Debug().Select("email").Where("email = ?", user.Email).Find(&u)
	// if err := global.RY_DB.Debug().Select("email").Where("email = ?", user.Email).Find(&model.User{}).Error; err == nil {
	// 	return nil, errors.New("邮箱已注册")
	// }

	if u.Email == user.Email {
		return nil, errors.New("邮箱已注册")
	}
	// 对明文进行加密处理
	// newPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	newPassword, _ := utils.Enbcrypt(user.Password)
	// 创建用户
	newUser := model.User{
		Id:       utils.NextId(),
		Username: user.Username,
		Email:    user.Email,
		Password: string(newPassword),
	}

	// 数据迁移
	//Db.AutoMigrate(&newUser)
	//创建数据
	// if err = Db.Table("t_user").CreateTable(&newUser).Error; err != nil {
	// 	return newUser, errors.New("数据创建失败")
	// }
	// if err := Db.Exec("INSERT INTO "+command.DBUser+"(id,username,password,email,create_time,update_time) VALUES(?,?,?,?,?,?)",
	// 	newUser.Id, newUser.Username, newUser.Password, newUser.Email, newUser.CreateTime, newUser.UpdateTime).Error; err != nil {
	// 	return nil, errors.New("用户注册失败")
	// }
	if err := global.RY_DB.Debug().Create(&newUser).Error; err != nil {
		return nil, errors.New("用户注册失败")
	}
	return &newUser, nil
}

// 登录
func (us UserService) Login(user model.User) (*model.User, error) {
	// Db := utils.GetDB()
	var u model.User
	// 判断用户名是否存在
	if err := global.RY_DB.Debug().Select("id,username,password,email,avatar").Where("username = ?", user.Username).Find(&u).Error; err != nil {
		return nil, errors.New("账户不存在")
	}

	// 校验密码
	if err := utils.Debcrypt(u.Password, user.Password); err != nil {
		return nil, errors.New("密码错误")
	}
	// if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)); err != nil {
	// 	return nil, errors.New("密码错误")
	// }

	return &u, nil
}
