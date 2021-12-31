package user

import (
	"errors"
	"time"

	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

// 注册
func (us *UserService) Register(user model.User) (model.User, error) {
	var err error
	Db := utils.GetDB()
	// 判断用户名是否存在
	if err = Db.Where("username = ?", user.Username).Find(&user).Error; err == nil {
		//ctx.JSON(500, gin.H{"code": 500, "msg": "用户名已存在"})
		// command.Failed(ctx, http.StatusInternalServerError, 500, "用户名已存在")
		return user, errors.New("用户名已存在")
	}
	// 对明文进行加密处理
	newPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		// ctx.JSON(500, gin.H{"code": 500, "msg": "加密失败"})
		// command.Failed(ctx, http.StatusInternalServerError, 500, "加密失败")
		return user, errors.New("加密失败")
	}
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
	if err = Db.Create(&newUser).Error; err != nil {
		// ctx.JSON(500, gin.H{"code": 500, "msg": "数据创建失败"})
		// command.Failed(ctx, http.StatusInternalServerError, 500, "数据创建失败")
		return newUser, errors.New("数据创建失败")
	}

	return newUser, err
}

// 登录
func (us *UserService) Login(user model.User) (model.User, error) {
	var err error
	Db := utils.GetDB()
	// 判断用户名是否存在
	if err = Db.Where("username = ?", user.Username).Find(&user).Error; err != nil {
		// ctx.JSON(500, gin.H{"code": 500, "msg": "账户不存在"})
		// command.Failed(ctx, http.StatusInternalServerError, 500, "账户不存在")
		return user, errors.New("账户不存在")
	}

	// 校验密码
	// if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user2.Password)); err != nil {
	// 	// ctx.JSON(401, gin.H{"code": 401, "msg": "密码错误"})
	// 	// command.Failed(ctx, http.StatusUnauthorized, 401, "密码错误")
	// 	return user, errors.New("密码错误")
	// }

	return user, err
}
