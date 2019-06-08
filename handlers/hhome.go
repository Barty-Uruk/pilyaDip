package handlers

import (
	"fmt"
	"net/http"
	"polyadip/models"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	ads, err := models.GetAllAds()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"Ads": ads,
	})
}
