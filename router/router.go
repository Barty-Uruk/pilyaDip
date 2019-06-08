package router

import (

	// cfg "mmt/services/crm/config"

	"fmt"
	"html/template"
	"io"
	"polyadip/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Hook is a function to process middleware.
func Hook() echo.MiddlewareFunc {
	return Authentication
}

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return RequireAuthentication(c, next)
	}
}
func RequireAuthentication(c echo.Context, next echo.HandlerFunc) error {

	_, err := c.Cookie("auth")
	if err != nil {
		c.Redirect(303, "/signin")
		return err
	}
	next(c)

	return nil
}

// Init - binding middleware and setup routers
func Init() *echo.Echo {

	e := echo.New()
	// Recover and CORS middlewar

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	templates, err := template.New("").ParseGlob(("views/*.html"))
	if err != nil {
		fmt.Printf("error parse views, %s", err)
	}
	closeGroup := e.Group("")

	closeGroup.Use(Hook())
	t := &Template{
		templates: templates,
	}

	e.Renderer = t
	closeGroup.GET("/home", handlers.Home)
	closeGroup.GET("/ad/:id", handlers.GetAd)
	e.POST("/auth", handlers.Authorization)
	e.GET("/signin", handlers.Signin)
	closeGroup.GET("/newAd", handlers.NewAd)
	e.GET("/signup", handlers.SignUp)
	closeGroup.GET("/myads", handlers.MyAds)
	e.POST("/reg", handlers.Reg)
	closeGroup.POST("/createAd", handlers.CreateAd)
	closeGroup.POST("/filterads", handlers.FilterAds)
	closeGroup.GET("/logout", handlers.Logout)
	e.Use(middleware.Logger())
	// Middleware for logging
	// logrus.SetLevel(logrus.DebugLevel)
	// e.Use(logrusmiddleware.Hook())
	// newLogger(l)
	// e.Use(Hook())

	// // Recover and CORS middlewar
	// e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	// // Authorisation group
	// jwtSec := config.JWTSecret()
	// jwtGroup := e.Group("/api/" + apiversion + "/auth")
	// jwtGroup.POST("/new", h.RegisterDriverHandler)
	// jwtGroup.POST("/verification", h.PhoneVerificationHandler)
	// jwtGroup.POST("/refresh", h.RefreshTokenHandler)

	// openGroup := e.Group("/api/" + apiversion)
	// openGroup.POST("/checkconnections", h.CheckConnections)
	// openGroup.GET("/checkversion", h.CheckVersion)
	// openGroup.GET("/updatefile", func(c echo.Context) error {
	// 	return c.Attachment("services/driver/updatefile/app-release.apk", "app-release.apk")
	// })
	// // JWT middleware
	// o := e.Group("/api/" + apiversion)
	// jwtSec = config.JWTSecret()
	// o.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(jwtSec),
	// }))

	// // Actions

	// o.POST("/initdata", h.GetInitData)
	// o.POST("/setstate", h.SetDriverState)
	// o.POST("/order", h.OfferHandler)
	// o.GET("/myoffers", h.MyOffers)

	return e
}
