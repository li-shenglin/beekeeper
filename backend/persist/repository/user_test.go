package repository

import (
	"fmt"
	"testing"
)

func TestFindByNameAndPass(t *testing.T) {
	user := NewUser()
	pass := user.FindByNameAndPass("lsl", "ere")
	fmt.Println(pass)
}
