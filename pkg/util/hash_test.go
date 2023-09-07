package util

import (
	"fmt"
	"testing"
)

func Test_BcryptHash(t *testing.T) {
	pwd := "123456"
	hash := BcryptHash(pwd)
	fmt.Println(hash) // $2a$10$ga9fLpTQ9F35/93p62WfyewLHF4s0vDw0QnkBeTzXbY6MyQVyxHuu
	fmt.Println("check: ", BcryptCheck(pwd, hash))
}
