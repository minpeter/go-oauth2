package database

import (
	"encoding/base64"
	"fmt"
)

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func AddUser(user *User) error {
	_, err := Engine.Insert(user)
	return err
}

func GetUserByName(name string) (*User, error) {
	user := &User{}
	has, err := Engine.Where("name = ?", name).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("not found user by name: %s", name)
	}
	return user, nil
}

func IsExistUserByName(name string) (bool, error) {
	user := &User{}
	has, err := Engine.Where("name = ?", name).Get(user)
	return has, err
}

func (u *User) NewToken() string {

	token := base64.StdEncoding.EncodeToString([]byte(u.Name))

	return token
}

func GetUserByToken(token string) (*User, error) {
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	name := string(decoded)

	user := &User{}

	_, err = Engine.Where("name = ?", name).Get(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
