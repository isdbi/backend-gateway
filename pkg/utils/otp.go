package utils

import (
	"math/rand"
)


var digits = []byte{'0','1','2','3','4','5','6','7','8','9'}

func GenerateOTP(size int) string{
  otp := make([]byte,size)
  for i:= range size {
      otp[i] = digits[rand.Intn(8)]
  } 

  return string(otp)
}
