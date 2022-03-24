package system

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
)

type UserService struct{}

/**
* 注册
 */
func (*UserService) Register(user model.User) (*model.User, error) {
	var u model.User

	global.QX_DB.Debug().Select("email").Where("email = ?", user.Email).Find(&u)
	if u.Email == user.Email {
		return nil, errors.New("邮箱已注册")
	}

	// 对明文进行加密处理
	newPassword, _ := utils.Encrypt(user.Password)
	// 创建用户
	newUser := model.User{
		Username: user.Email,
		Email:    user.Email,
		Password: newPassword,
	}

	if err := global.QX_DB.Debug().Create(&newUser).Error; err != nil {
		return nil, errors.New("用户注册失败")
	}
	return &newUser, nil
}

/**
* 登录
 */
func (*UserService) Login(user model.User) (*model.User, error) {
	var u model.User

	// 判断用户名是否存在
	global.QX_DB.Debug().Select("id,username,avatar,email,password").Where("email = ?", user.Email).Find(&u)
	if u.Email != user.Email {
		return nil, errors.New("邮箱未注册")
	}

	if err := utils.Decrypt(u.Password, user.Password); err != nil {
		return nil, errors.New("密码不匹配")
	}

	return &u, nil
}

/**
* 获取用户信息
 */
func (*UserService) GetUserInfo(id uint64) (*model.User, error) {
	var user model.User
	err := global.QX_DB.Debug().Select("id,username,avatar").Where("id = ?", id).Find(&user).Error

	return &user, err
}

/**
* 修改用户名
 */
func (*UserService) UpdateUsername(user model.User) error {
	var u model.User
	global.QX_DB.Debug().Select("id,username").Where("username = ?", user.Username).Find(&u)
	if u.Username == user.Username {
		return errors.New("用户名已存在")
	}

	global.QX_DB.Debug().Model(&u).Where("id = ?", user.Id).Update("username", user.Username)
	return nil
}
