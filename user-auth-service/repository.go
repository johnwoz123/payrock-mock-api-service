package main

import (
	"github.com/jinzhu/gorm"
	proto "github.com/johnwoz123/payrock-mock-api-service/user-auth-service/proto/user"
)

type Repository interface {
	GetAll() ([]*proto.User, error)
	Get(id string) (*proto.User, error)
	Create(user *proto.User) error
	GetByEmailAndPassword(user *proto.User) (*proto.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func (uRepo *UserRepo) GetAll() ([]*proto.User, error) {
	var users []*proto.User
	if err := uRepo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (uRepo *UserRepo) Get(id string) (*proto.User, error) {
	var user *proto.User
	user.Id = id
	if err := uRepo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (uRepo *UserRepo) Create(user *proto.User) error {
	if err := uRepo.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (uRepo *UserRepo) GetByEmailAndPassword(user *proto.User) (*proto.User, error) {
	if err := uRepo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
