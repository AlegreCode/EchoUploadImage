package middlewares

import (
	"mime/multipart"

	"math"
	"net/http"

	"github.com/inhies/go-bytesize"

	"github.com/gookit/validate"
	"github.com/labstack/echo"
)

func ValidateImg(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := validate.FromRequest(c.Request())
		if err != nil {
			panic(err)
		}

		v := data.Create()
		v.AddValidator("imgSize", func(val interface{}) bool {
			data_img := val.(*multipart.FileHeader)
			MB := math.Floor(float64(data_img.Size)/float64(bytesize.MB)*100) / 100

			if MB < 1 {
				return true
			}
			return false
		})

		v.StopOnError = false
		v.AddRule("name", "required").SetMessage("The name field is required.")
		v.AddRule("name", "minLen", 4).SetMessage("The least length of name field is 4.")
		v.AddRule("file", "required").SetMessage("The file field is required.")
		v.AddRule("file", "image", "jpeg", "jpg", "png", "gif").SetMessage("The file field must be an image jpg, png or gif.")
		v.AddRule("file", "imgSize").SetMessage("Image size exceeded.")
		if v.Validate() {
			return next(c)
		}
		// Set Flash messages
		c.(*CustomContext).SetFlash("error", v.Errors)
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
}
