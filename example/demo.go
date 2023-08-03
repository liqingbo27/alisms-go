package main

import (
	"gin-html/alisms-go/SmsClient"
	"log"
	"net/http"
)

const (
	accessKeyID     = "yourAccessKeyId"     //
	secretAccessKey = "yourAccessKeySecret" //
)

func main() {
	sc, err := SmsClient.NewSMSClient(accessKeyID, secretAccessKey)
	if err != nil {
		return
	}
	statusCode, _, _ := sc.SendSMS(SmsClient.Params{"18888888888", "李清波", "SMS_186578512", `{"code":"12345"}`})
	if statusCode == http.StatusOK {
		log.Println("发送成功")
	} else {
		log.Println("发送失败")
	}
}
