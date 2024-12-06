package users

import (
	"errors"
	"go-echo-boilerplate/common"
	"go-echo-boilerplate/database"
	model "go-echo-boilerplate/src/models/users/model"
	"sync"
)

type usersService struct{}

var singleton UsersService
var once sync.Once

func GetUsersService() UsersService {
	if singleton != nil {
		return singleton
	}
	once.Do(func() {
		singleton = &usersService{}
	})
	return singleton
}

func SetUsersService(service UsersService) UsersService {
	original := singleton
	singleton = service
	return original
}

type UsersService interface {
	FindUserByPhone(phone string) (*model.User, error)
	AddUser(phone string, password string) (*model.User, error)
	ChangePwd(phone string, password string) (*model.User, error)
	FindUserById(id uint) (*model.User, error)
	IsPhoneExist(phone string) bool
}

func (u *usersService) FindUserByPhone(phone string) (*model.User, error) {
	db := database.Instance()
	var user model.User
	res := db.First(&user, "phone = ?", phone)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected != 1 {
		return nil, errors.New("phone number not found")
	}
	return &user, nil
}

func (u *usersService) AddUser(phone string, password string) (*model.User, error) {
	db := database.Instance()
	user := model.User{
		Role:     common.User,
		Phone:    phone,
		Password: password,
		// FirstName: firstName,
		// LastName: lastName,
		// MiddleName: middleName,
		// Email: email,
	}

	res := db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected != 1 {
		return nil, errors.New("doesn't create user")
	}

	return &user, nil
}

func (u *usersService) ChangePwd(phone string, password string) (*model.User, error) {
	db := database.Instance()
	var user model.User

	res := db.First(&user, "phone = ?", phone)
	if res.Error != nil || res.RowsAffected != 1 {
		return nil, errors.New("user not found")
	}
	user.Password = password

	res = db.Save(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected != 1 {
		return nil, errors.New("doesn't update password")
	}

	return &user, nil
}

func (u *usersService) FindUserById(id uint) (*model.User, error) {
	db := database.Instance()
	var user model.User

	res := db.First(&user, "id = ?", id)
	if res.Error != nil || res.RowsAffected != 1 {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (u *usersService) IsPhoneExist(phone string) bool {
	db := database.Instance()
	var user model.User

	res := db.First(&user, "phone = ?", phone)
	if res.Error != nil || res.RowsAffected != 1 {
		return false
	}

	return true
}
