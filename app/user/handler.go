package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/dto"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
	"golang.org/x/crypto/bcrypt"
)

// 注册
func registerHandler(ctx *gin.Context) {
	var user model.User
	// 绑定表单数据
	ctx.ShouldBind(&user)

	// if err := user.Validate(); err != nil {
	// 	command.Failed(ctx, http.StatusInternalServerError, "数据验证失败")
	// 	return
	// }

	var userService UserService
	_, err := userService.Register(user)

	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Db := utils.GetDB()
	// // 判断用户名是否存在
	// if err = Db.Where("username = ?", user.Username).Find(&user).Error; err == nil {
	// 	// ctx.JSON(500, gin.H{"code": 500, "msg": "用户名已存在"})
	// 	command.Failed(ctx, http.StatusInternalServerError, 500, "用户名已存在")
	// 	return
	// }

	// 判断邮箱是否注册
	// if err = Db.Where("email = ?", user.Email).Find(&user).Error; err == nil {
	// 	// ctx.JSON(500, gin.H{"code": 500, "msg": "邮箱已注册"})
	// 	command.Failed(ctx, http.StatusInternalServerError, 500, "邮箱已注册")
	// 	return
	// }
	// 对明文进行加密处理
	// newPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	// ctx.JSON(500, gin.H{"code": 500, "msg": "加密失败"})
	// 	command.Failed(ctx, http.StatusInternalServerError, 500, "加密失败")
	// 	return
	// }

	// 创建用户
	// newUser := model.User{
	// 	Id:         utils.NextId(),
	// 	Username:   user.Username,
	// 	Password:   string(newPassword),
	// 	CreateTime: time.Now(),
	// 	EditTime:   time.Now(),
	// }

	// 数据迁移
	// Db.AutoMigrate(&newUser)
	//创建数据
	// if err = Db.Create(&newUser).Error; err != nil {
	// 	// ctx.JSON(500, gin.H{"code": 500, "msg": "数据创建失败"})
	// 	command.Failed(ctx, http.StatusInternalServerError, 500, "数据创建失败")
	// 	return
	// }

	// token := utils.CreateToken(newUser.Id)
	// ctx.JSON(200, gin.H{"code": 200, "msg": "注册成功", "token": token})
	command.Success(ctx, "注册成功", nil)
}

// 登录
func loginHandler(ctx *gin.Context) {
	// 绑定表单参数
	var form model.User
	ctx.ShouldBind(&form)

	// Db := utils.GetDB()
	// user := new(model.User)

	// if err := Db.Where("username = ?", form.Username).Find(&user).Error; err != nil {
	// 	// ctx.JSON(500, gin.H{"code": 500, "msg": "账户不存在"})
	// 	command.Failed(ctx, http.StatusInternalServerError, 500, "账户不存在")
	// 	return
	// }

	var userService UserService
	user, err := userService.Login(form)
	if err != nil {
		command.Failed(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		// ctx.JSON(401, gin.H{"code": 401, "msg": "密码错误"})
		command.Failed(ctx, http.StatusUnauthorized, "密码错误")
		return
	}

	// 生成token
	token := utils.CreateToken(user.Id)
	// ctx.JSON(200, gin.H{"code": 200, "msg": "登录成功"})
	command.Success(ctx, "登录成功", gin.H{"token": token})
}

// 获取用户信息
func infoHandler(ctx *gin.Context) {
	userInfo, _ := ctx.Get("user")
	// ctx.JSON(http.StatusOK, gin.H{"user": userInfo})
	command.Success(ctx, "登录成功", gin.H{"user": dto.ToUserDto(userInfo.(model.User))})
}
