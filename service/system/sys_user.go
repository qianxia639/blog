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
	err := global.DB.Debug().Where("username = ?", r.Username).First(&model.User{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名已存在")
	}

	// 对明文进行加密处理
	newPassword, _ := utils.Encrypt(r.Password)
	// 创建用户
	newUser := model.User{
		UUID:     uuid.NewV4().String(),
		Username: r.Username,
		Nickname: r.Username,
		Password: newPassword,
	}

	err = global.DB.Debug().Create(&newUser).Error
	return &newUser, err
}

/**
* 登录
 */
func (*UserService) Login(l request.Login) (*model.User, error) {
	var u model.User

	// 判断用户名是否存在
	err := global.DB.Debug().Where("username = ?", l.Username).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名不存在")
	}

	if err := utils.Decrypt(u.Password, l.Password); err != nil {
		return nil, errors.New("密码不匹配")
	}

	return &u, nil
}

/**
* 获取用户信息
 */
func (*UserService) GetUserInfo(id uint64, uuid string) (*model.User, error) {
	var user model.User
	err := global.DB.Debug().Where("id = ? AND uuid = ?", id, uuid).First(&user).Error
	return &user, err
}

/**
* 修改用户名
 */
func (*UserService) UpdateNickname(nickname string, id uint64, uuid string) error {

	err := global.DB.Debug().Where("nickname = ?", nickname).First(&model.User{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("用户名已存在")
	}

	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 修改user表中的username
		if err := tx.Debug().Model(&model.User{}).Where("id = ? AND uuid = ?", id, uuid).Update("nickname", nickname).Error; err != nil {
			return err
		}

		// 修改blog表中的username
		if err := tx.Debug().Model(&model.Blog{}).Where("user_id = ?", id).Update("nickname", nickname).Error; err != nil {
			return err
		}

		return nil
	})
}

/**
*	修改密码
 */
func (*UserService) UpdatePwd(u request.UpdatePwd, id uint64, uuid string) error {

	var user model.User
	if err := global.DB.Debug().Where("id = ? AND uuid = ?", id, uuid).First(&user).Error; err != nil {
		return errors.New("数据不存在")
	}

	if err := utils.Decrypt(user.Password, u.OldPassword); err != nil {
		return errors.New("旧密码错误")
	}

	// pwd, _ := utils.Encrypt(u.LastPassword)

	// return global.DB.Model(&model.User{}).Debug().Where("signer = ?", u.Signer).Update("password", pwd).Error
	return nil
}

/**
*	找回密码
 */
func (*UserService) ForgetPwd(f request.ForgetPwd) error {
	var user model.User

	if err := utils.Decrypt(user.Password, f.Password); err == nil {
		return errors.New("新密码不能与原密码相同")
	}

	// newPassword, _ := utils.Encrypt(f.Password)

	// global.DB.Model(user).Debug().Where("signer = ?", f.Signer).Update("password", newPassword)
	return nil
}

/**
*	修改头像
 */
func (*UserService) UpdateAvatar(url, uuid string, id uint64) error {
	return global.DB.Model(&model.User{}).Where("id = ? AND uuid = ?", id, uuid).Update("avatar", url).Error
}
