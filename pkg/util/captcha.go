package util

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore
var driver base64Captcha.Driver = base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)

// Generate 生成验证码
func Generate() (string, string, string) {
	c := base64Captcha.NewCaptcha(driver, store)
	// 获取
	id, b64s, err := c.Generate()
	if err != nil {
		panic(fmt.Sprintf("生成验证码错误：%s", err.Error()))
	}
	content := store.Get(id, false)

	return id, content, b64s
}

// Verify 验证验证码
func Verify(id string, val string) bool {
	if id == "" || val == "" {
		return false
	}
	// 同时清理掉这个图片
	return store.Verify(id, val, true)
}
