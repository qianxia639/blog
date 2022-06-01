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
func (us *UserService) Register(r request.Register) error {
	err := global.QX_DB.Debug().Where("email = ?", r.Email).First(&model.User{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("邮箱已注册")
	}

	// 对明文进行加密处理
	newPassword, _ := utils.Encrypt(r.Password)
	// 创建用户
	newUser := model.User{
		UUID:     uuid.NewV4().String(),
		Username: r.Email,
		Email:    r.Email,
		Password: newPassword,
	}

	return global.QX_DB.Debug().Create(&newUser).Error
}

/**
* 登录
 */
func (*UserService) Login(l request.Login) (*model.User, error) {
	var u model.User

	// 判断用户名是否存在
	err := global.QX_DB.Debug().Where("email = ?", l.Email).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
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
func (*UserService) GetUserInfo(id uint64, uuid string) (*model.User, error) {
	var user model.User
	err := global.QX_DB.Debug().Where("id = ? AND uuid = ?", id, uuid).First(&user).Error

	return &user, err
}

func (*UserService) QueryUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := global.QX_DB.Debug().Where("email = ?", email).Find(&user).Error

	return &user, err
}

/**
* 修改用户名
 */
func (*UserService) UpdateUsername(username string, id uint64, uuid string) error {

	err := global.QX_DB.Debug().Where("username = ?", username).First(&model.User{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("用户名已存在")
	}

	return global.QX_DB.Transaction(func(tx *gorm.DB) error {
		// 修改user表中的username
		if err := tx.Debug().Model(&model.User{}).Where("id = ? AND uuid = ?", id, uuid).Update("username", username).Error; err != nil {
			return err
		}

		// 修改blog表中的username
		if err := tx.Debug().Model(&model.Blog{}).Where("user_id = ?", id).Update("username", username).Error; err != nil {
			return err
		}

		return nil
	})
}

/**
*	修改密码
 */
func (*UserService) UpdatePwd(u request.UpdatePwd, id uint64, uuid string) error {

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
*	找回密码
 */
func (*UserService) ForgetPwd(f request.ForgetPwd) error {
	var user model.User
	global.QX_DB.Debug().Where("email = ?", f.Email).Find(&user)

	if err := utils.Decrypt(user.Password, f.Password); err == nil {
		return errors.New("不能与旧密码相同")
	}

	pwd, _ := utils.Encrypt(f.Password)

	global.QX_DB.Model(user).Where("email = ?", f.Email).Update("password", pwd)
	return nil
}

/**
*	修改头像
 */
func (*UserService) UpdateAvatar(u request.UpdateAvatar, id uint64, uuid string) error {
	return global.QX_DB.Model(&model.User{}).Where("id = ? AND uuid = ?", id, uuid).Update("avatar", u.Avatar).Error
}

/**
*	修改邮箱
 */
func (*UserService) UpdateEmail(u request.UpdateEmail, id uint64, uuid string) error {

	// code, _ := utils.GetCache(u.OldEmail)

	// if code != u.Code {
	// 	return errors.New("验证码不相符")
	// }

	return global.QX_DB.Model(&model.User{}).Debug().Where("id = ? AND uuid = ?", id, uuid).Update("email", u.LastEmail).Error
}
