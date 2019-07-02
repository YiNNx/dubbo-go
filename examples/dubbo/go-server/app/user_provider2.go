package main

import (
	"context"
	"fmt"
	"strconv"
)

import (
	"github.com/apache/dubbo-go/config"
	"github.com/dubbogo/hessian2/java_exception"
	perrors "github.com/pkg/errors"
)

func init() {
	config.SetProviderService(new(UserProvider2))
}

type UserProvider2 struct {
}

func (u *UserProvider2) getUser(userId string) (*User, error) {
	if user, ok := userMap[userId]; ok {
		return &user, nil
	}

	return nil, fmt.Errorf("invalid user id:%s", userId)
}

func (u *UserProvider2) GetUser(ctx context.Context, req []interface{}, rsp *User) error {
	var (
		err  error
		user *User
	)

	println("req:%#v", req)
	user, err = u.getUser(req[0].(string))
	if err == nil {
		*rsp = *user
		println("rsp:%#v", rsp)
	}
	return err
}

func (u *UserProvider2) GetUser0(id string, name string) (User, error) {
	var err error

	println("id:%s, name:%s", id, name)
	user, err := u.getUser(id)
	if err != nil {
		return User{}, err
	}
	if user.Name != name {
		return User{}, perrors.New("name is not " + user.Name)
	}
	return *user, err
}

func (u *UserProvider2) GetUser2(ctx context.Context, req []interface{}, rsp *User) error {
	var err error

	println("req:%#v", req)
	rsp.Id = strconv.Itoa(int(req[0].(int32)))
	return err
}

func (u *UserProvider2) GetUser3() error {
	return nil
}

func (u *UserProvider2) GetErr(ctx context.Context, req []interface{}, rsp *User) error {
	return java_exception.NewThrowable("exception")
}

func (u *UserProvider2) GetUsers(req []interface{}) ([]interface{}, error) {
	var err error

	println("req:%s", req)
	t := req[0].([]interface{})
	user, err := u.getUser(t[0].(string))
	if err != nil {
		return nil, err
	}
	println("user:%v", user)

	return []interface{}{user}, err
}

func (s *UserProvider2) MethodMapper() map[string]string {
	return map[string]string{
		"GetUser2": "getUser",
	}
}

func (u *UserProvider2) Reference() string {
	return "UserProvider2"
}
