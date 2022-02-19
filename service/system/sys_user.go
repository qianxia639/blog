package system

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
)

type UserService struct{}

// 注册
func (*UserService) Register(user model.User) (*model.User, error) {
	var u model.User
	if err := global.RY_DB.Debug().Select("username").Where("username = ?", user.Username).Find(&u).Error; err == nil {
		if u.Username == user.Username {
			return nil, errors.New("用户名已存在")
		}
	}

	if err := global.RY_DB.Debug().Select("email").Where("email = ?", user.Email).Find(&u).Error; err == nil {
		if u.Email == user.Email {
			return nil, errors.New("邮箱已注册")
		}
	}

	// 对明文进行加密处理
	newPassword, _ := utils.Enbcrypt(user.Password)
	// 创建用户
	newUser := model.User{
		Id:       utils.NextId(),
		Username: user.Username,
		Email:    user.Email,
		Password: string(newPassword),
	}

	if err := global.RY_DB.Debug().Create(&newUser).Error; err != nil {
		return nil, errors.New("用户注册失败")
	}
	return &newUser, nil
}

// 登录
func (*UserService) Login(user model.User) (*model.User, error) {
	var u model.User
	// 判断用户名是否存在
	if err := global.RY_DB.Debug().Select("id,username,password,email,avatar").Where("username = ?", user.Username).Find(&u).Error; err != nil {
		return nil, errors.New("账户不存在")
	}

	// 校验密码
	if err := utils.Debcrypt(u.Password, user.Password); err != nil {
		return nil, errors.New("密码错误")
	}
	return &u, nil
}
