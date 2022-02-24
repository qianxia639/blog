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
		Username: user.Email,
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
	if err := global.RY_DB.Debug().Select("id,username,password,email,avatar").Where("email = ?", user.Email).Find(&u).Error; err == nil {
		if u.Email != user.Email {
			return nil, errors.New("账户不存在")
		}
	}
	// 校验密码
	if err := utils.Debcrypt(u.Password, user.Password); err != nil {
		return nil, errors.New("密码错误")
	}
	return &u, nil
}

// 修改名称
func (*UserService) UpdateUsername(user model.User) error {
	var u model.User
	if err := global.RY_DB.Debug().Select("id,username").Where("username = ?", user.Username).Find(&u).Error; err == nil {
		if u.Username == user.Username {
			return errors.New("不能更改为当前用户名")
		}
	}
	if err := global.RY_DB.Debug().Model(&u).Where("id = ?", user.Id).Update("username", user.Username).Error; err != nil {
		return errors.New("用户名修改失败")
	}
	return nil
}

// 修改密码
func (*UserService) UpdatePassword(m map[string]string, user model.User) error {
	if err := utils.Debcrypt(user.Password, m["oldPwd"]); err != nil {
		return errors.New("密码错误")
	}

	global.RY_DB.Debug().Model(&user).Where("id = ?", user.Id).Update("password", m["latestPwd"])
	return nil
}
