package main

import (
	"html/template"

	"github.com/alegrecode/echo/EchoUploadImage/controllers"

	. "github.com/alegrecode/echo/EchoUploadImage/helpers"
	. "github.com/alegrecode/echo/EchoUploadImage/middlewares"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("assets"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("alegrecode"))))

	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseFiles("views/home.html", "views/base.html", "views/alert.partial.html", "views/card_image.partial.html"))

	e.Renderer = &TemplateRegistry{
		Templates: templates,
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return next(cc)
		}
	})

	e.GET("/", controllers.Home)

	e.POST("/upload", controllers.Upload, ValidateImg)

	e.DELETE("/delete/:id", controllers.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
