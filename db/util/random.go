package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabets = "abcdefghijklmnopqrstuvwsyz"

func RandomOwner() string {
	numRange := 3
	var genString strings.Builder

	for i := 0; i < numRange; i++ {
		var randS string = string(alphabets[rand.Intn(len(alphabets))])
		genString.WriteString(randS)
	}
	return genString.String()

}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomInt(min int, max int) int64 {
	return int64(min + rand.Intn(max-min+1))
}

func RandomBalance() int64 {
	min := 0
	max := 100000
	return int64(min + rand.Intn(max-min+1))
}

func RandomString(length int) string {
	var sb strings.Builder

	k := len(alphabets)
	for i := 0; i < length; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomCurrency() string {
	curArr := []string{"USD", "NGN", "AUD"}
	length := len(curArr)
	return curArr[rand.Intn(length)]
}
