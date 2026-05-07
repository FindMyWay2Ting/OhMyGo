package user

import (
	"GopherAI/common/code"
	myemail "GopherAI/common/email"
	myredis "GopherAI/common/redis"
	"GopherAI/config"
	"GopherAI/dao/user"
	"GopherAI/model"
	"GopherAI/utils"
	"GopherAI/utils/myjwt"
	"log"
	"strings"
)

func Login(username, password string) (string, code.Code) {
	var userInformation *model.User
	var ok bool
	//1:判断用户是否存在
	if ok, userInformation = user.IsExistUser(username); !ok {

		return "", code.CodeUserNotExist
	}
	//2:判断用户是否密码账号正确
	if userInformation.Password != utils.MD5(password) {
		return "", code.CodeInvalidPassword
	}
	//3:返回一个Token
	token, err := myjwt.GenerateToken(userInformation.ID, userInformation.Username)

	if err != nil {
		return "", code.CodeServerBusy
	}
	return token, code.CodeSuccess
}

func Register(email, password, captcha string) (string, string, code.Code) {

	var ok bool
	var userInformation *model.User

	//1:先判断用户是否已经存在了
	if ok, _ := user.IsExistEmail(email); ok {
		return "", "", code.CodeUserExist
	}

	//2:从redis中验证验证码是否有效
	if ok, _ := myredis.CheckCaptchaForEmail(email, captcha); !ok {
		return "", "", code.CodeInvalidCaptcha
	}

	//3：生成11位的账号
	username := utils.GetRandomNumbers(11)

	//4：注册到数据库中
	if userInformation, ok = user.Register(username, email, password); !ok {
		return "", username, code.CodeServerBusy
	}

	//5：将账号一并发送到对应邮箱上去，后续需要账号登录
	if isEmailConfigured() {
		if err := myemail.SendCaptcha(email, username, user.UserNameMsg); err != nil {
			log.Printf("send username email failed: %v", err)
		}
	}

	// 6:生成Token
	token, err := myjwt.GenerateToken(userInformation.ID, userInformation.Username)

	if err != nil {
		return "", username, code.CodeServerBusy
	}

	return token, username, code.CodeSuccess
}

// 往指定邮箱发送验证码
// 分为以下任务：
// 1：先存放redis
// 2：再进行远程发送
func SendCaptcha(email_ string) (string, code.Code) {
	send_code := utils.GetRandomNumbers(6)
	//1:先存放到redis
	if err := myredis.SetCaptchaForEmail(email_, send_code); err != nil {
		return "", code.CodeServerBusy
	}

	//2:再进行远程发送
	if isEmailConfigured() {
		if err := myemail.SendCaptcha(email_, send_code, myemail.CodeMsg); err != nil {
			return "", code.CodeServerBusy
		}
		return "", code.CodeSuccess
	}

	return send_code, code.CodeSuccess
}

func isEmailConfigured() bool {
	c := config.GetConfig().EmailConfig
	email := strings.TrimSpace(c.Email)
	auth := strings.TrimSpace(c.Authcode)
	if email == "" || auth == "" {
		return false
	}
	lower := strings.ToLower(email + " " + auth)
	if strings.Contains(lower, "your") {
		return false
	}
	return true
}
