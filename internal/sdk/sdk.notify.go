package sdk

import "github.com/lishimeng/app-starter/factory"

type SendSms func(code, mobile string) error
type SendEmail func(code, email string) error

func GetEmailSender() SendEmail {
	var sender SendEmail
	_ = factory.Get(&sender)
	return sender
}

func GetSmsSender() SendSms {
	var sender SendSms
	_ = factory.Get(&sender)
	return sender
}
