package models

import "gin-app/db"

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "user"
}

type User struct {
	ID        int64  `json:"id"`       // 列名为 `id`
	Username  string `json:"username"` // 列名为 `username`
	Sex       string `json:"sex"`      // 列名为 `password`
	City      string `json:"city"`
	Telephone string `json:"telephone"`
}

func FindUser(id int) (user *User, err error) {
	user = new(User)
	if err = db.DB.Where("id=?", id).Find(user).Error; err != nil {
		return nil, err
	}
	return
}
func CreateUser(user *User) (err error) {
	err = db.DB.Create(&user).Error
	return
}

func FindUsers() (users []*User, err error) {
	if err = db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return
}

func DeleteUser(id int) (err error) {
	err = db.DB.Where("id=?", id).Delete(&User{}).Error
	return
}
