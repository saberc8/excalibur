package common

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Auth struct{}

// Captcha 获取验证码
func (a Auth) Captcha(ctx *gin.Context) {
	svg, code := utils.GenerateSVG(80, 40)
	// Generate a unique key for the captcha
	captchaKey := uuid.New().String()
	// Store the captcha in Redis with the unique key
	err := global.AquilaRedis.Set(context.Background(), "captcha-"+captchaKey, code, 5*time.Minute).Err()
	if err != nil {
		fmt.Println("Redis set error:", err)
		return
	}
	// Return the captcha key to the client
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"captchaKey": captchaKey,
	// 	"captchaSvg": string(svg),
	// })
	utils.Success(ctx, gin.H{
		"captchaKey": captchaKey,
		"captchaSvg": string(svg),
	})
}

// Login 用户登录
func (a Auth) Login(ctx *gin.Context) {
	var req LoginDto
	err := ctx.ShouldBind(&req)
	if err != nil {
		utils.Fail(ctx, "登录失败")
		return
	}
	captcha, err := global.AquilaRedis.Get(context.Background(), "captcha-"+req.CaptchaKey).Result()
	if captcha != req.Code {
		utils.Fail(ctx, "验证码错误")
		return
	}
	defer global.AquilaRedis.Del(context.Background(), "captcha-"+req.CaptchaKey)
	var user model.UserEntity
	err = global.AquilaDb.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		utils.Fail(ctx, "用户不存在")
		return
	}
	if user.Password != fmt.Sprintf("%x", md5.Sum([]byte(req.Password))) {
		utils.Fail(ctx, "密码错误")
		return
	}
	var LoginInfo LoginVo
	fmt.Println("user:", user.ID)
	LoginInfo.Token = utils.GenerateToken(int(user.ID))
	LoginInfo.UserInfo = UserVo{
		ID:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
		Nickname: user.Nickname,
	}
	utils.Success(ctx, LoginInfo)
}

func (a Auth) Register(ctx *gin.Context) {
	var req RegisterDto
	err := ctx.ShouldBind(&req)
	fmt.Println(err)
	if err != nil {
		fmt.Println("step1", err)
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	captcha, err := global.AquilaRedis.Get(context.Background(), req.CaptchaKey).Result()
	if captcha != req.Code {
		utils.Fail(ctx, "验证码错误")
		return
	}
	defer global.AquilaRedis.Del(context.Background(), req.CaptchaKey)
	var user model.UserEntity
	// q: Find() 和 First() 的区别
	// a: Find() 返回的是一个数组，First() 返回的是一个对象
	err = global.AquilaDb.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		// 创建新用户
		user = model.UserEntity{
			Username: req.Username,
			Password: fmt.Sprintf("%x", md5.Sum([]byte(req.Password))),
		}
		err = global.AquilaDb.Transaction(func(tx *gorm.DB) error {
			err = tx.Create(&user).Error
			if err != nil {
				return err
			}
			utils.Success(ctx, nil)
			return nil
		})
		if err != nil {
			utils.Fail(ctx, "用户创建失败")
			return
		}
		utils.Success(ctx, nil)
		return
	} else {
		utils.Fail(ctx, "用户已存在")
		return
	}
}

func (a Auth) Logout(ctx *gin.Context) {
	utils.Success(ctx, nil)
}
