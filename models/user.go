package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//User пользователи
type User struct {
	ID       int
	Login    string `form:"login"`
	Phone    string `form:"phone"`
	Name     string `form:"name"`
	Password string `form:"password"`
}

func (user User) NewUser() (User, error) {
	err := user.validate()
	if err != nil {
		return user, fmt.Errorf("Ошибка валидации,%v", err)
	}
	_, err = db.Model(&user).Insert()
	if err != nil {
		return user, fmt.Errorf("Ошибка сохранения пользователя в базу,%v", err)
	}
	return user, nil
}

func (user User) validate() error {
	if user.Login == "" {
		return fmt.Errorf("Пользователь с пустым логином недопустим")
	}
	if user.Password == "" {
		return fmt.Errorf("Пользователь с пустым паролем недопустим")
	}
	return nil
}

func (user User) AuthenticateUser() (User, error) {
	var dbUser User
	err := db.Model(&dbUser).Where("login = ?", user.Login).Select()
	if err != nil {
		return dbUser, fmt.Errorf("Ошибка авторизации, (логин)%v", err)
	}
	encryptionErr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if encryptionErr != nil {
		return dbUser, fmt.Errorf("Ошибка авторизации, (пароль)%v", encryptionErr)
	}
	return dbUser, nil
}
func GetUserByID(id int) User {
	var (
		dbUser User
	)
	_ = db.Model(&dbUser).Where("id = ?", id).Select()
	return dbUser
}
