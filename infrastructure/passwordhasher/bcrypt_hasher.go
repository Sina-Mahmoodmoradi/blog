package passwordhasher

import (
	"fmt"

	"github.com/Sina-Mahmoodmoradi/blog/internal/usecase"
	"golang.org/x/crypto/bcrypt"
)



type BcryptHasher struct{}

func NewBcryptHasher()usecase.PasswordHasher{
	return &BcryptHasher{}
}


	
func (b *BcryptHasher)Hash(password string)(string,error){
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),10)
	if err!=nil{
		return "",fmt.Errorf("failed to hash password: %w",err)
	}
	return string(bytes),nil
}

func (b *BcryptHasher)Compare(hash, password string) bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err==nil
}
