package system

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/request"
	"github.com/qianxia/blog/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct{}

/**
* 注册
 */
func (us *UserService) Register(r request.Register) (*model.User, error) {
	var u model.User

	if !errors.Is(global.QX_DB.Debug().Where("email = ?", r.Email).First(&u).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("邮箱已注册")
	}

	// 对明文进行加密处理
	newPassword, _ := utils.Encrypt(r.Password)
	// 创建用户
	newUser := model.User{
		UUID:     uuid.NewV4(),
		Username: r.Email,
		Email:    r.Email,
		Password: newPassword,
	}

	if err := global.QX_DB.Debug().Create(&newUser).Error; err != nil {
		return nil, errors.New("注册失败")
	}
	return &newUser, nil
}

/**
* 登录
 */
func (*UserService) Login(l request.Login) (*model.User, error) {
	var u model.User

	// 判断用户名是否存在
	if errors.Is(global.QX_DB.Debug().Where("email = ?", l.Email).First(&u).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("邮箱未注册")
	}

	if err := utils.Decrypt(u.Password, l.Password); err != nil {
		return nil, errors.New("密码不匹配")
	}

	return &u, nil
}

/**
* 获取用户信息
 */
func (*UserService) GetUserInfo(id uint64, uuid uuid.UUID) (*model.User, error) {
	var user model.User
	err := global.QX_DB.Debug().Select("id,uuid,username,avatar").Where("id = ? AND uuid = ?", id, uuid).Find(&user).Error

	return &user, err
}

/**
* 修改用户名
 */
func (*UserService) UpdateUsername(u request.UpdateUsername, id uint64, uuid uuid.UUID) error {

	if !errors.Is(global.QX_DB.Debug().Where("username = ?", u.Username).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已存在")
	}

	return global.QX_DB.Transaction(func(tx *gorm.DB) error {
		// 修改user表中的username
		if err := tx.Debug().Model(&model.User{}).Where("id = ? AND uuid = ?", id, uuid).Update("username", u.Username).Error; err != nil {
			return err
		}

		// 修改blog表中的username
		if err := tx.Debug().Model(&model.Blog{}).Where("user_id = ?", id).Update("username", u.Username).Error; err != nil {
			return err
		}

		return nil
	})
}

/**
*	修改密码
 */
func (*UserService) UpdatePwd(u request.UpdatePwd, id uint64, uuid uuid.UUID) error {

	// 密码校验
	var user model.User
	if err := global.QX_DB.Debug().Where("id = ? AND uuid = ?", id, uuid).First(&user).Error; err != nil {
		return errors.New("数据不存在")
	}

	if err := utils.Decrypt(user.Password, u.OldPassword); err != nil {
		return errors.New("旧密码错误")
	}

	pwd, _ := utils.Encrypt(u.LastPassword)

	return global.QX_DB.Model(&model.User{}).Where("id = ? AND uuid = ?", id, uuid).Update("password", pwd).Error
}

/**
*	修改头像
 */
func (*UserService) UpdateAvatar(u request.UpdateAvatar, id uint64, uuid uuid.UUID) error {
	return global.QX_DB.Model(&model.User{}).Where("id = ? AND uuid = ?", id, uuid).Update("avatar", u.Avatar).Error
}
