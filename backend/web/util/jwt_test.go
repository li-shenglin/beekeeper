package util

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	token, _ := GenerateToken(123, "haha")
	fmt.Println("token:", token)
	claim, err := ParseToken(token)
	fmt.Println(claim.Valid())
	if err != nil {
		fmt.Println("ParseToken err：", err)
	} else if time.Now().Unix() > claim.ExpiresAt {
		fmt.Println("timeout")
	} else {
		fmt.Println("userid:", claim.UserID)
	}
}
