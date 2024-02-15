package model

import "github.com/lishimeng/app-starter"

// UserProfile 用户档案,每个用户必须有
type UserProfile struct {
	app.Pk
	UserCode              string     `orm:"column(user_code)"`                // 用户编号
	RealName              string     `orm:"column(real_name)"`                //  真实姓名
	IdCard                string     `orm:"column(id_card)"`                  //  身份证号
	IdCardVerified        VerifyFlag `orm:"column(id_card_verified)"`         // 身份证号验证标记
	PhoneNumber           string     `orm:"column(phone_number)"`             //  手机号
	PhoneNumberVerified   VerifyFlag `orm:"column(phone_number_verified)"`    // 手机号验证标记
	WechatUnionId         string     `orm:"column(wechat_union_id)"`          // 微信UnionId
	WechatUnionIdVerified VerifyFlag `orm:"column(wechat_union_id_verified)"` // 微信UnionId验证标记
	app.TableChangeInfo
}

type VerifyFlag int

const (
	Verified   = 1
	UnVerified = 0
)

type Mfa struct { // 特定设计, 应减少批量检索MFA item
	app.Pk
	MfaCode           string
	UserCode          string
	MfaType           MfaCategory
	SecretPhoneNumber string
	SecretEmail       string
	app.TableChangeInfo
}

type MfaCategory string

const (
	MfaPhoneNumber MfaCategory = "phone_number"            // 手机号
	MfaEmail       MfaCategory = "email"                   //  邮箱
	MfaWechat      MfaCategory = "wechat"                  // 微信消息
	MfaGoogle      MfaCategory = "google_authenticator"    // google验证器
	MfaMicrosoft   MfaCategory = "microsoft_authenticator" // 微软验证器
)

type MfaItem struct {
	app.Pk
	MfaCode  string      `orm:"column(mfa_code)"`     // MFA编号
	Sn       string      `orm:"column(user_code)"`    // 验证编号
	Category MfaCategory `orm:"column(mfa_category)"` // 验证类型
	app.TableChangeInfo
}

// MfaDevice 受信任设备列表
type MfaDevice struct {
	app.Pk
	MfaCode string `orm:"column(mfa_code)"` // MFA编号
	Key     string `orm:"column(key)"`      // MFA设备
	app.TableInfo
}
