package controllers

import (
	"net/http"

	. "github.com/alegrecode/echo/EchoUploadImage/middlewares"
	. "github.com/alegrecode/echo/EchoUploadImage/models"

	uuid "github.com/satori/go.uuid"

	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	var img Image
	images := img.Read()
	flash := c.(*CustomContext).GetFlash()
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"datos": images,
		"flash": flash,
	})
}

func Upload(c echo.Context) error {
	json, _ := c.(*CustomContext).Upload("file", []string{"uploads"}...)
	if json {
		c.(*CustomContext).SetFlash("success", "The file has been uploaded successfully.")
	}
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func Delete(c echo.Context) error {
	var img Image
	img.Id, _ = uuid.FromString(c.Param("id"))
	img.Delete()
	return c.Redirect(http.StatusMovedPermanently, "/")
}
