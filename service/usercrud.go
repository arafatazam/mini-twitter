package service

import (
	"github.com/arafatazam/mini-twitter/dto"
	"gorm.io/gorm"
)

type UserCRUD struct {
	DB *gorm.DB
}

func (uc *UserCRUD) Create (req *dto.NewUserReq) error {
	//todo
	return nil
}