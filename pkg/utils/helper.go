package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-mail/mail"
)

func Debug(obj any) {
	raw, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(raw))
}

func LocalTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Calcutta")
	return time.Now().In(loc)
}

func ConvertStringTimeToTime(t string) time.Time {
	// layout := "2006-01-02 15:04:05.999 IST"
	layout := "2006-01-02 15:04:05.999 -0700 MST"
	result, err := time.Parse(layout, t)
	if err != nil {
		log.Printf("Error: Parse time failed: %s", err.Error())
	}
	//log.Println("time : ", result)
	//loc, _ := time.LoadLocation("Asia/Calcutta")
	return result
}

func GenerateOtp() string {
	return fmt.Sprintf("%06d", 100000+rand.Intn(900000))
}

func SendOtpToEmail(fromEmail, password, toEmail, otp string) error {
	m := mail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "OTP Verification")
	m.SetBody("text/plain", "Your OTP is : "+otp)

	// fmt.Println("From : ", fromEmail, "\n to : ", toEmail, "\n Pass : ", password)
	//SMTP port (TLS): 587 - SMTP port (SSL): 465
	d := mail.NewDialer("smtp.gmail.com", 587, fromEmail, password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
