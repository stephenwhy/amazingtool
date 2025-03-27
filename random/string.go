package random

import (
	"crypto/rand"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
	"math/big"
	mathRand "math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ23456789"

const Length = 8

const (
	numberLength = 16 // 总长度为16位
	dateLength   = 8  // 日期部分长度为8位，格式为YYYYMMDD
)

func Encode(u uuid.UUID) string {
	return base58.Encode(u[:])
}

func Decode(s string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode(s))
}

func GenerateUUID() string {
	return shortuuid.New()
}

func GenerateRandomString() string {
	result := make([]byte, Length)
	for i := 0; i < Length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}
	return string(result)
}

func GenerateRandom16Str() string {
	result := make([]byte, numberLength)
	for i := 0; i < numberLength; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}
	return string(result)
}

func GenerateUniqueCardNumber() string {
	currentTime := time.Now()
	datePart := currentTime.Format("20060102")

	// 生成随机数部分
	mathRand.NewSource(time.Now().UnixNano())
	randomPart := fmt.Sprintf("%0*d", dateLength, mathRand.Intn(100000000))

	// 拼接日期部分和随机数部分
	cardNumber := datePart + randomPart

	return cardNumber
}
