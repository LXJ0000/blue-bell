package service

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/encrypt"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"errors"
)

type UserService struct {
}

func (u *UserService) Login(req *models.UserLoginRequest) (token string, err error) {
	userDao := mysql.NewUserDao()
	isExist, user := userDao.CheckUserExist(req.Username)
	if !isExist || !encrypt.CheckPassword(user.Password, req.Password) {
		return "", errors.New("用户名或密码错误")
	}
	return jwt.GenToken(user.UserID, user.UserName)
}

func (u *UserService) Register(req *models.UserRegisterRequest) error {
	userDao := mysql.NewUserDao()
	//1. 判断用户是否存在
	if isExist, _ := userDao.CheckUserExist(req.Username); isExist {
		return errors.New("用户名已存在")
	}
	//2. 生成UID
	userID := snowflake.GenID()
	user := &models.UserModel{
		UserID:   userID,
		UserName: req.Username,
		Password: encrypt.GetPassword(req.Password),
	}
	//3. 入库
	if err := userDao.Create(user); err != nil {
		return err
	}
	return nil
}
