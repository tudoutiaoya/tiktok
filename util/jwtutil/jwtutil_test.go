package jwtutil

import (
	"fmt"
	"testing"
)

func TestJWT(t *testing.T) {
	token, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNjc4NjcwNzg3fQ.KCLC3JZPOmb_KN9o0G2tPY5CNETt2zlm_PXMbkeR3QU")
	if err != nil {
		fmt.Println(err)
		fmt.Println("出现错误：|", err.Error())
	}
	fmt.Println("id为:", token)
}
