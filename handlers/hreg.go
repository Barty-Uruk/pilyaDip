package handlers

import (
	"fmt"
	"net/http"
	"polyadip/models"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

const (
	cookieName = "auth"
	cookieExp  = 24 * time.Hour
)

func Reg(c echo.Context) error {
	var (
		user models.User
	)

	err := c.Bind(&user)
	if err != nil {
		fmt.Println("error binding form date,", err)
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error generate encrypted pass,", err)
		return err
	}
	user.Password = string(encryptedPassword)
	_, err = user.NewUser()
	// password := c.FormValue("password")
	// if err != nil {
	// 	return c.JSON(http.StatusOK, fmt.Errorf("ошибка получения данных с формы,%s", err))
	// }
	if err != nil {
		return c.Render(http.StatusOK, "signup.html", map[string]interface{}{
			"name":   []string{"Tom", "Bob", "Sam"},
			"errors": fmt.Sprintf("%v", err),
		})
	}
	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Value = strconv.Itoa(user.ID)
	cookie.Expires = time.Now().Add(cookieExp)
	c.SetCookie(cookie)
	return c.Redirect(http.StatusSeeOther, "/home")
}

func Authorization(c echo.Context) error {
	var (
		user models.User
	)
	err := c.Bind(&user)
	if err != nil {
		fmt.Println("error binding form date,", err)
	}
	newUser, err := user.AuthenticateUser()
	if err != nil {
		msg := fmt.Errorf("error Authenticate User,%v", err)
		fmt.Println(msg)
		return c.Render(http.StatusOK, "signin.html", map[string]interface{}{
			"name":  []string{"Tom", "Bob", "Sam"},
			"error": "Ошибка авторизации",
		})
	}
	// password := c.FormValue("password")
	// if err != nil {
	// 	return c.JSON(http.StatusOK, fmt.Errorf("ошибка получения данных с формы,%s", err))
	// }
	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Value = strconv.Itoa(newUser.ID)
	cookie.Expires = time.Now().Add(cookieExp)
	c.SetCookie(cookie)
	return c.Redirect(http.StatusMovedPermanently, "/home")
}

func Signin(c echo.Context) error {
	return c.Render(http.StatusOK, "signin.html", map[string]interface{}{
		"name":  []string{"Tom", "Bob", "Sam"},
		"error": "",
	})
}
func Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = cookieName
	cookie.Expires = time.Now()
	c.SetCookie(cookie)
	return c.Redirect(http.StatusSeeOther, "/home")
}
func SignUp(c echo.Context) error {
	return c.Render(http.StatusOK, "signup.html", map[string]interface{}{
		"error": "",
	})
}
