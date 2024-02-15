package mfa

import "github.com/kataras/iris/v12"

func Route(root iris.Party) {

	root.Post("/{code}/phone_number", bindPhoneNumber)
	root.Post("/{code}/phone_number/{phone_number}", sendPhoneNumberCode)
	root.Post("/{code}/email", bindEmail)
	root.Post("/{code}/email/{email}", sendEmailCode)
	root.Post("/{code}/wechat", bindWechat)
}
