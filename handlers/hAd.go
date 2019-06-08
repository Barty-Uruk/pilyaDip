package handlers

import (
	"fmt"
	"net/http"
	"polyadip/models"
	"strconv"

	"github.com/labstack/echo"
)

//CreateAd создает новое предложение
func CreateAd(c echo.Context) error {
	var (
		ad models.Ad
	)
	ad.Price = 0
	cookie, _ := c.Cookie("auth")
	err := c.Bind(&ad)
	if err != nil {
		fmt.Println("error binding form date,", err)
	}
	ad.UserID, _ = strconv.Atoi(cookie.Value)
	newAd, err := ad.Create()
	if err != nil {
		msg := fmt.Errorf("error creating ad,%v", err)
		fmt.Println(msg)
		return c.Render(http.StatusBadRequest, "newAd.html", map[string]interface{}{
			"error": "msg",
		})
	}
	path := fmt.Sprintf("/ad/%v", newAd.ID)
	return c.Redirect(http.StatusSeeOther, path)
}

func NewAd(c echo.Context) error {
	return c.Render(http.StatusOK, "newAd.html", map[string]interface{}{
		"error": "",
	})
}
func GetAd(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ad, err := models.GetAdByID(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	user := models.GetUserByID(ad.UserID)
	return c.Render(http.StatusOK, "ad.html", map[string]interface{}{
		"ad":   ad,
		"user": user,
	})
}

func MyAds(c echo.Context) error {
	cookie, _ := c.Cookie("auth")
	userid, err := strconv.Atoi(cookie.Value)
	if err != nil {
		fmt.Println(err)
	}
	ads, err := models.GetAdsByUserID(userid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.Render(http.StatusOK, "myads.html", map[string]interface{}{
		"Ads": ads,
	})
}

func FilterAds(c echo.Context) error {
	var (
		fd models.FilterData
	)
	err := c.Bind(&fd)
	if err != nil {
		fmt.Println("error binding form date,", err)
	}
	ads, err := fd.GetFilterAds()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.Render(http.StatusOK, "filterads.html", map[string]interface{}{
		"FilterData": fd,
		"Ads":        ads,
	})
}
