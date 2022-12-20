package aop

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestExample1(t *testing.T) {

	userAop := NewUser("nice", "qwerty")

	userBeforeAdvice := RegisterBefore[User]
	userAfterAdvice := RegisterAfter[User]

	vna := ValidateNameAdvice{}
	vnaBefore := userBeforeAdvice(vna)
	vnaAfter := userAfterAdvice(vna)

	vpa := ValidatePasswordAdvice{MaxLength: 10, MinLength: 6}
	vpaBefore := userBeforeAdvice(vpa)
	vpaAfter := userAfterAdvice(vpa)

	ctx := context.Background()

	// Be careful with the loading order here.
	err := userAop.Proxy(ctx,
		vpaBefore, vnaBefore,
		vpaAfter, vnaAfter)

	if err != nil {
		panic(err)
	}

}

// NewUser instantiate generic AOP for User.
func NewUser(name, pass string) AOP[User] {
	userProxyAop := AOP[User]{}
	user := User{
		Name: name,
		Pass: pass,
	}
	userProxyAop.SetProxy(user)
	return userProxyAop
}

type User struct {
	Name string
	Pass string
}

func (u User) Auth(ctx context.Context) error {
	fmt.Printf("user:%s, use pass:%s\n", u.Name, u.Pass)
	return nil
}

func (u User) Pointcut(ctx context.Context) error {
	return u.Auth(ctx)
}

type ValidateNameAdvice struct {
}

func (ValidateNameAdvice) Before(ctx context.Context, user *User) error {
	fmt.Println("ValidateNameAdvice before")
	if user.Name == "admin" {
		return errors.New("admin can't be used")
	}
	return nil
}

func (ValidateNameAdvice) After(ctx context.Context, user *User) {
	fmt.Println("ValidateNameAdvice after")
	fmt.Printf("username:%s validate sucess\n", user.Name)
}

type ValidatePasswordAdvice struct {
	MinLength int
	MaxLength int
}

func (advice ValidatePasswordAdvice) Before(ctx context.Context, user *User) error {
	fmt.Println("ValidatePasswordAdvice before")
	if user.Pass == "123456" {
		return errors.New("pass isn't strong")
	}

	if len(user.Pass) > advice.MaxLength {
		return fmt.Errorf("len of pass must less than:%d", advice.MaxLength)
	}

	if len(user.Pass) < advice.MinLength {
		return fmt.Errorf("len of pass must greater than:%d", advice.MinLength)
	}

	return nil
}

func (ValidatePasswordAdvice) After(ctx context.Context, user *User) {
	fmt.Println("ValidatePasswordAdvice after")
	fmt.Printf("password:%s validate sucess\n", user.Pass)
}
