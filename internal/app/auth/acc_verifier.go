package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lai0xn/isdb/internal/repository"
	"github.com/lai0xn/isdb/pkg/mail"
	"github.com/lai0xn/isdb/pkg/utils"
	"github.com/redis/go-redis/v9"
)

type Verifier interface {
	Send(user repository.User) error
	Verify(user repository.User, token string) error
}

type EmailVerifier struct{
  Rd redis.Client 
}

func (e EmailVerifier) Send(user repository.User) error {
  otp := utils.GenerateOTP(6)
  err := e.Rd.Set(context.Background(),"otp:"+user.Email,otp,time.Hour).Err()
  if err != nil {
    return err
  }
  
  messageStr := fmt.Sprintf("From: nebultech@example.com\r\n"+
		"To: %s\r\n"+
		"Subject: Email Verification\r\n"+
		"\r\n"+
		"Your OTP is %s.\r\n", user.Email, otp)

	messageBytes := []byte(messageStr)

  err = mail.SendEmail([]string{user.Email},messageBytes)

 
  if err != nil {
    return err
  }
  
  return err
}

func (e EmailVerifier) Verify(user repository.User, otp string) error {
  userOTP,err := e.Rd.Get(context.Background(),"otp:"+user.Email).Result()
  if err != nil {
    return err
  }
  if ! (userOTP == otp) {
    return errors.New("otps dont match")
  }

  return nil
}
