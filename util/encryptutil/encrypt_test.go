package encryptutil

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	password := "zzqshuai"
	hashPassword, _ := HashPassword(password)
	fmt.Println("密码是:" + hashPassword)
	isEqual := CheckPasswordHash(password, hashPassword)
	fmt.Println(isEqual)
}
